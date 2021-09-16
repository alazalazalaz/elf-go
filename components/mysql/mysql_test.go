package mysql

import (
	"elf-go/components/logs"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

type User struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Phone     int64
	Username  string
	Password  string
	ClassId   int
	ClassName string
	IdCard    string
	Level     int
	Status    int
	SchoolNum int
	CreatedAt int
}

func (User) TableName() string {
	return "t_user"
}

func TestConnect(t *testing.T) {
	//连接
	//db := Connect()
	//_insert(db)
	//_batchInsert(db)
}

func _batchInsert(db *gorm.DB) {
	totalNum := 0
	batchNum := 1000
	for j := 0; j <= totalNum/batchNum; j++ {
		nums := batchNum
		users := make([]User, nums)
		for i := 0; i < nums; i++ {
			userEntity := User{
				Phone:     _getPhoneRandNum(),
				Username:  _randomString(int(_getRandNum(8, 16)), []rune("abcd")),
				Password:  _randomString(int(_getRandNum(8, 16))),
				ClassId:   int(_getRandNum(1, 10)),
				ClassName: "三年二班",
				IdCard:    fmt.Sprintf("%d%s", _getRandNum(100000000000000000, 999999999999999999), _randomString(1)),
				Level:     int(_getRandNum(1, 100)),
				Status:    int(_getRandNum(1, 4)),
				SchoolNum: int(_getRandNum(1, 9999999)),
				CreatedAt: int(time.Now().Unix()),
			}
			users[i] = userEntity
		}
		db.CreateInBatches(users, nums)
	}
}

func _getPhoneRandNum() int64 {
	var min int64 = 1000000000
	var max int64 = 9999999999
	return 10000000000 + _getRandNum(min, max)
}

func _getRandNum(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// RandomString returns a random string with a fixed length
func _randomString(n int, allowedChars ...[]rune) string {
	defaultLetters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func _insert(db *gorm.DB) {
	userEntity := User{
		Phone:     123234,
		Username:  "xiong",
		Password:  "xxx",
		ClassId:   1,
		ClassName: "三年二班",
		IdCard:    "555024131231234s",
		Level:     1,
		Status:    1,
		CreatedAt: int(time.Now().Unix()),
	}
	if err := db.Create(&userEntity).Error; err != nil {
		logs.Error("保存数据失败啦", logs.Content{"err:": err})
	}
	logs.Info("执行完成", nil)
}
