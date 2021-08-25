package entity

type User struct{
	Id int `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	CreatedAt int64 `json:"created_at"`
}

type Article struct{
	Id int `json:"id"`
	Title string `json:"title"`
	CreatedAt int64 `json:"created_at"`
}
