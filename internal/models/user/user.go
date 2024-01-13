package user

type UserInfo struct {
	Id         string     `json:"_id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Github     string     `json:"github"`
	Medium     string     `json:"medium"`
	Twitter    string     `json:"twitter"`
	LinkedIn   string     `json:"linkedin"`
	Tagline    string     `json:"tagline"`
	Objective  string     `json:"objective"`
	Theme      string     `json:"theme"`
	Skills     UserSkills `json:"skills"`
	Projects   []Projects `json:"projects"`
	Experience []Experience `json:"experience"`
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

type Experience struct {
	Title         string   `json:"title"`
	Organization  string   `json:"organization"`
	StartDate     string   `json:"start_date"`
	EndDate       string   `json:"end_date"`
	IsPresent     string   `json:"is_present"`
	Contributions []string `json:"contributions"`
}

type UserRoute struct {
	Username string `json:"username"`
	Theme    string `json:"theme"`
}
