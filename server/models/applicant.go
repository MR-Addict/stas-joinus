package models

import "time"

type Applicant struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Modified      bool      `json:"modified" gorm:"default:false"`
	Submitted_At  time.Time `json:"submitted_at" gorm:"autoCreateTime"`
	Name          string    `json:"name" validate:"required,min=2,max=20"`
	Gender        string    `json:"gender" validate:"required,oneof=boy girl"`
	Phone         string    `json:"phone" validate:"required,len=11,number"`
	Email         string    `json:"email" validate:"required,email,max=320"`
	QQ            string    `json:"qq" validate:"required,min=5,max=11,number"`
	Student_ID    string    `json:"student_id" validate:"required,len=12,number"`
	College       string    `json:"college" validate:"required,max=50"`
	Major         string    `json:"major" validate:"required,max=50"`
	First_Choice  string    `json:"first_choice" validate:"required,min=5,max=7"`
	Second_Choice string    `json:"second_choice" validate:"required,min=5,max=7,nefield=First_Choice"`
	Introduction  string    `json:"introduction" validate:"required,min=4,max=500"`
}
