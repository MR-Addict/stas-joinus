package models

type Choice struct {
	Boy  int `json:"boy"`
	Girl int `json:"girl"`
}

type StatsDepartment struct {
	Name          string `json:"name"`
	First_Choice  Choice `json:"first_choice"`
	Second_Choice Choice `json:"second_choice"`
}
