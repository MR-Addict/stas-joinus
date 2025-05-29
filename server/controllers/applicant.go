package controllers

import (
	"server/configs"
	"server/models"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

		// check First_Choice, Second_Choice and Introduction is same or not
		if duplicatedUser.First_Choice == applicant.First_Choice && duplicatedUser.Second_Choice == applicant.Second_Choice && duplicatedUser.Introduction == applicant.Introduction {
			return c.Status(400).JSON(models.Response{Success: false, Message: "志愿未发生变化"})
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
	getAll, _ := strconv.ParseBool(c.Query("all", "false"))

	// Add query param
	search := c.Query("query", "")

	var total int64
	query := configs.Db.Model(&models.Applicant{})

	// If the search query is not empty, add a where clause to filter applicants
	// using LIKE for multiple fields
	if search != "" {
		command := ""
		like := "%" + search + "%"
		fields := []string{"name", "gender", "phone", "email", "qq", "student_id", "college", "major", "first_choice", "second_choice", "introduction"}

		for i, field := range fields {
			if i > 0 {
				command += " OR "
			}
			command += field + " LIKE ?"
		}

		query = query.Where(command, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Count total applicants
	query.Count(&total)

	// Get pagination parameters
	page := 1
	pageSize := 20

	if getAll {
		page = 1
		pageSize = int(total)
	} else {
		page, _ = strconv.Atoi(c.Query("page", "1"))
		pageSize, _ = strconv.Atoi(c.Query("page_size", "20"))
	}

	// Query applicants with pagination
	findResult := query.Order("submitted_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applicants)
	if findResult.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: findResult.Error.Error()})
	} else {
		return c.Status(200).JSON(models.Response{Success: true, Message: "报名信息获取成功", Data: applicants, Pagination: models.Pagination{Page: page, Total: total, Page_Size: pageSize, Query: search}})
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
