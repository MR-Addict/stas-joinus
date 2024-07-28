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
		return c.Status(400).JSON(models.Response{Success: false, Message: "你已经提交过啦，请勿重复提交"})
	}

	// add new record
	addResult := configs.Db.Create(&applicant)
	if addResult.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: addResult.Error.Error()})
	} else if addResult.RowsAffected > 0 {
		return c.Status(201).JSON(models.Response{Success: true, Message: "表单提交成功", Data: applicant})
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
		row := i + 2
		file.SetCellValue(name, "A"+fmt.Sprintf("%d", row), i+1)
		file.SetCellValue(name, "B"+fmt.Sprintf("%d", row), d.Name)
		file.SetCellValue(name, "C"+fmt.Sprintf("%d", row), d.Gender)
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
