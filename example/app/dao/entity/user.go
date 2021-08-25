package entity

type User struct{
	ID int `gorm:"primary_key"`
	Username string
	CreatedAt int64
	Article Article `gorm:"foreignkey:UserID"` //关联查询的案例，mysql不需要建立外键，这里指定就行
}

type Article struct{
	ID int `gorm:"primary_key"`
	UserID int
	Title string
	CreatedAt int
}
