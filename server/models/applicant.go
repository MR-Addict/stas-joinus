package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Applicant struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name" validate:"required,min=2,max=20"`
	Gender        string             `json:"gender" validate:"required,oneof=男 女"`
	Phone         string             `json:"phone" validate:"required,len=11,number"`
	Email         string             `json:"email" validate:"required,email,max=320"`
	QQ            string             `json:"qq" validate:"required,min=5,max=11,number"`
	Student_ID    string             `json:"student_id" validate:"required,len=12,number"`
	College       string             `json:"college" validate:"required,max=50"`
	Major         string             `json:"major" validate:"required,max=50"`
	Create_At     int64              `json:"create_at"`
	First_Choice  string             `json:"first_choice" validate:"required,min=5,max=7"`
	Second_Choice string             `json:"second_choice" validate:"required,min=5,max=7,nefield=First_Choice"`
	Introduction  string             `json:"introduction" validate:"required,min=4,max=500"`
}
