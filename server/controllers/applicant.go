package controllers

import (
	"fmt"
	"server/configs"
	"server/models"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

var validate = validator.New()

func ApplicantCreate(c *fiber.Ctx) error {
	var applicant models.Applicant
	applicant.Submitted_At = time.Now()

	// parse time from config
	startTime, err := time.Parse(time.RFC3339, configs.Config.App.StartTime)
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}

	endTime, err := time.Parse(time.RFC3339, configs.Config.App.EndTime)
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}

	// check if the application is still open
	if applicant.Submitted_At.Before(startTime) || applicant.Submitted_At.After(endTime) {
		return c.Status(400).JSON(models.Response{Success: false, Message: "不在报名时间内无法提交"})
	}

	// parse request body
	if err := c.BodyParser(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	}

	if err := validate.Struct(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "表单内容不合法", Data: err.Error()})
	}

	var duplicatedUser models.Applicant
	findOptions := &models.Applicant{Student_ID: applicant.Student_ID}
	findResult := configs.Db.Where(findOptions).Limit(1).Find(&duplicatedUser)

	// check whether applicant has already exists
	if findResult.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: findResult.Error.Error()})
	} else if findResult.RowsAffected > 0 {
		// check if the applicant has been modified
		if duplicatedUser.Modified {
			return c.Status(400).JSON(models.Response{Success: false, Message: "抱歉，志愿仅能修改一次"})
		}

		// only update First_Choice, Second_Choice, Introduction and set Modified to true
		duplicatedUser.First_Choice = applicant.First_Choice
		duplicatedUser.Second_Choice = applicant.Second_Choice
		duplicatedUser.Introduction = applicant.Introduction
		duplicatedUser.Modified = true

		updateResult := configs.Db.Save(&duplicatedUser)
		if updateResult.Error != nil {
			return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: updateResult.Error.Error()})
		} else if updateResult.RowsAffected > 0 {
			return c.Status(200).JSON(models.Response{Success: true, Message: "恭喜，志愿更新成功", Data: duplicatedUser})
		}
	}

	// add new record
	addResult := configs.Db.Create(&applicant)
	if addResult.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: addResult.Error.Error()})
	} else if addResult.RowsAffected > 0 {
		return c.Status(201).JSON(models.Response{Success: true, Message: "恭喜，报名成功", Data: applicant})
	}

	return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误"})
}

func ApplicantQuery(c *fiber.Ctx) error {
	var applicants []models.Applicant
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))

	var total int64
	configs.Db.Model(&models.Applicant{}).Count(&total)

	findResult := configs.Db.Order("submitted_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applicants)
	if findResult.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: findResult.Error.Error()})
	} else {
		return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: applicants, Pagination: models.Pagination{Page: page, Total: total, Page_Size: pageSize}})
	}
}

func ApplicantStats(c *fiber.Ctx) error {
	departments := []string{
		"技术开发部", "科普活动部", "组织策划部", "新闻宣传部", "对外联络部", "双创联合服务部",
	}

	var results []struct {
		Department string
		Gender     string
		Choice     string
		Count      int
	}

	// query database
	query := `
		SELECT first_choice as department, gender, 'first' as choice, count(*) as count
		FROM applicants
		WHERE first_choice IN ?
		GROUP BY first_choice, gender
		UNION ALL
		SELECT second_choice as department, gender, 'second' as choice, count(*) as count
		FROM applicants
		WHERE second_choice IN ?
		GROUP BY second_choice, gender
	`

	configs.Db.Raw(query, departments, departments).Scan(&results)

	// create stats map
	statsMap := make(map[string]*models.StatsDepartment)
	// initialize stats map
	for _, department := range departments {
		statsMap[department] = &models.StatsDepartment{
			Name: department,
			First_Choice: models.Choice{
				Boy:  0,
				Girl: 0,
			},
			Second_Choice: models.Choice{
				Boy:  0,
				Girl: 0,
			},
		}
	}

	// fill stats map
	for _, result := range results {
		deptStats := statsMap[result.Department]
		if result.Choice == "first" {
			if result.Gender == "boy" {
				deptStats.First_Choice.Boy = result.Count
			} else if result.Gender == "girl" {
				deptStats.First_Choice.Girl = result.Count
			}
		} else if result.Choice == "second" {
			if result.Gender == "boy" {
				deptStats.Second_Choice.Boy = result.Count
			} else if result.Gender == "girl" {
				deptStats.Second_Choice.Girl = result.Count
			}
		}
	}

	// convert stats map to stats slice
	var stats []models.StatsDepartment
	for _, stat := range statsMap {
		stats = append(stats, *stat)
	}

	// return stats
	return c.Status(200).JSON(models.Response{
		Success: true,
		Message: "数据读取成功",
		Data:    stats,
	})
}

func ApplicantDownload(c *fiber.Ctx) error {
	// query databse
	var applicants []models.Applicant
	findResul := configs.Db.Order("submitted_at desc").Find(&applicants)
	if findResul.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: findResul.Error.Error()})
	}

	name := "科协报名表单"
	file := excelize.NewFile()
	sheet, _ := file.NewSheet(name)
	defer file.Close()

	// setup sheet header
	file.SetCellValue(name, "A1", "序号")
	file.SetCellValue(name, "B1", "姓名")
	file.SetCellValue(name, "C1", "性别")
	file.SetCellValue(name, "D1", "手机")
	file.SetCellValue(name, "E1", "邮箱")
	file.SetCellValue(name, "F1", "QQ")
	file.SetCellValue(name, "G1", "学号")
	file.SetCellValue(name, "H1", "学院")
	file.SetCellValue(name, "I1", "专业")
	file.SetCellValue(name, "J1", "提交时间")
	file.SetCellValue(name, "K1", "第一志愿")
	file.SetCellValue(name, "L1", "第二志愿")
	file.SetCellValue(name, "M1", "自我介绍")

	// write data to sheet
	for i, d := range applicants {
		// Map gender values
		gender := d.Gender
		if gender == "girl" {
			gender = "女"
		} else if gender == "boy" {
			gender = "男"
		}

		row := i + 2
		file.SetCellValue(name, "A"+fmt.Sprintf("%d", row), i+1)
		file.SetCellValue(name, "B"+fmt.Sprintf("%d", row), d.Name)
		file.SetCellValue(name, "C"+fmt.Sprintf("%d", row), gender)
		file.SetCellValue(name, "D"+fmt.Sprintf("%d", row), d.Phone)
		file.SetCellValue(name, "E"+fmt.Sprintf("%d", row), d.Email)
		file.SetCellValue(name, "F"+fmt.Sprintf("%d", row), d.QQ)
		file.SetCellValue(name, "G"+fmt.Sprintf("%d", row), d.Student_ID)
		file.SetCellValue(name, "H"+fmt.Sprintf("%d", row), d.College)
		file.SetCellValue(name, "I"+fmt.Sprintf("%d", row), d.Major)
		file.SetCellValue(name, "J"+fmt.Sprintf("%d", row), d.Submitted_At.Format("2006-01-02 15:04:05"))
		file.SetCellValue(name, "K"+fmt.Sprintf("%d", row), d.First_Choice)
		file.SetCellValue(name, "L"+fmt.Sprintf("%d", row), d.Second_Choice)
		file.SetCellValue(name, "M"+fmt.Sprintf("%d", row), d.Introduction)
	}

	// setup collum with
	file.SetColWidth(name, "D", "L", 15)
	file.SetColWidth(name, "M", "M", 100)

	// setup sheet style
	wrapStyle, _ := file.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{WrapText: true}})
	alignStyle, _ := file.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{Vertical: "top", Horizontal: "left"}})
	file.SetColStyle(name, "M", wrapStyle)
	file.SetColStyle(name, "A:L", alignStyle)

	// using default sheet
	file.SetActiveSheet(sheet)

	// Set the appropriate headers for the response
	c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=科协报名表单.xlsx")

	// Write the Excel data to the client
	if err := file.Write(c.Response().BodyWriter()); err != nil {
		return err
	}

	return nil
}
