package controllers

import (
	"context"
	"strconv"
	"time"

	"server/configs"
	"server/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()
var applicantCollection = configs.GetCollection("applicant")

func CreateApplicant(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var applicant models.Applicant
	defer cancel()

	if err := c.BodyParser(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "表单内容不合法", Data: err.Error()})
	}

	if validationError := validate.Struct(&applicant); validationError != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "表单内容不合法", Data: validationError.Error()})
	}

	applicant.ID = primitive.NewObjectID()
	applicant.Create_At = time.Now().UnixMilli()

	results, err := applicantCollection.InsertOne(ctx, applicant)
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}

	return c.Status(201).JSON(models.Response{Success: true, Message: "表单提交成功", Data: results.InsertedID})
}

func GetAllApplicants(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var applicants []models.Applicant
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
	for results.Next(ctx) {
		var singleApplicant models.Applicant
		if err = results.Decode(&singleApplicant); err != nil {
			return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
		}
		applicants = append(applicants, singleApplicant)
	}

	if len(applicants) == 0 {
		return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: []models.Applicant{}, Pagination: models.Pagination{Page: pageNumber, Page_Size: pageSize, Total_Pages: 1}})
	}

	totalApplicantsCount, err := applicantCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}
	totalPages := totalApplicantsCount / int64(pageSize)
	if totalApplicantsCount%int64(pageSize) != 0 {
		totalPages++
	}

	return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: applicants, Pagination: models.Pagination{Page: pageNumber, Page_Size: pageSize, Total_Pages: totalPages}})
}
