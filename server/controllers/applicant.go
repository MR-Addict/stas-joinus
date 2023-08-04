package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"server/configs"
	"server/models"
	"server/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()
var applicantCollection = configs.GetCollection("applicant")

func ApplicantCreate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var applicant models.Applicant
	defer cancel()

	if err := c.BodyParser(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	}

	if err := validate.Struct(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "表单内容不合法", Data: err.Error()})
	}

	applicant.ID = primitive.NewObjectID()
	applicant.Create_At = time.Now().UnixMilli()

	var duplicatedApplicant models.Applicant
	filter := bson.M{"name": applicant.Name, "student_id": applicant.Student_ID}
	err := applicantCollection.FindOne(ctx, filter).Decode(&duplicatedApplicant)
	if err == mongo.ErrNoDocuments {
		if _, err := applicantCollection.InsertOne(ctx, applicant); err != nil {
			return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
		}
		return c.Status(201).JSON(models.Response{Success: true, Message: "表单提交成功", Data: applicant})
	} else if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	} else {
		return c.Status(200).JSON(models.Response{Success: false, Message: "你已经提交过啦，请勿重复提交"})
	}
}

func ApplicantGetAll(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pageNumber, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))
	findOptions := options.Find()
	findOptions.SetSkip(int64((pageNumber - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSort(bson.D{{Key: "create_at", Value: -1}})

	results, err := applicantCollection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}
	defer results.Close(ctx)

	var applicants []models.Applicant
	for results.Next(ctx) {
		var singleApplicant models.Applicant
		if err = results.Decode(&singleApplicant); err != nil {
			return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
		}
		applicants = append(applicants, singleApplicant)
	}

	total, err := applicantCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}

	if len(applicants) == 0 {
		return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: []models.Applicant{}, Pagination: models.Pagination{Page: pageNumber, Total: total, Page_Size: pageSize}})
	}

	return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: applicants, Pagination: models.Pagination{Page: pageNumber, Total: total, Page_Size: pageSize}})
}

func ApplicantDownload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := applicantCollection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}
	defer results.Close(ctx)

	// query databse
	var applicants []models.Applicant
	for results.Next(ctx) {
		var singleApplicant models.Applicant
		if err = results.Decode(&singleApplicant); err != nil {
			return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
		}
		applicants = append(applicants, singleApplicant)
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
		file.SetCellValue(name, "J"+fmt.Sprintf("%d", row), utils.FormatDate(d.Create_At))
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
