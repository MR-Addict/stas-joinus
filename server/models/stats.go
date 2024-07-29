package models

type StatsDepartment struct {
	Name          string `json:"name"`
	Boy           int    `json:"boy"`
	Girl          int    `json:"girl"`
	First_Choice  int    `json:"first_choice"`
	Second_Choice int    `json:"second_choice"`
}
