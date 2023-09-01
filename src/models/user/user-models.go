package userModels

import "github.com/google/uuid"

type UserInfo struct {
	Id        uuid.UUID  `json:"_id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Github    string     `json:"github"`
	Medium    string     `json:"medium"`
	Twitter   string     `json:"twitter"`
	LinkedIn  string     `json:"linkedin"`
	Tagline   string     `json:"tagline"`
	Objective string     `json:"objective"`
	Theme     string     `json:"theme"`
	Skills    UserSkills     `json:"skills"`
	Projects  []Projects `json:"projects"`
}

type UserSkills struct {
	Hard     []string `json:"hard"`
	Soft     []string `json:"soft"`
	Interest []string `json:"interest"`
}

type Projects struct {
	Name        string `json:"name"`
	Preview     string `json:"preview"`
	Description string `json:"description"`
}
