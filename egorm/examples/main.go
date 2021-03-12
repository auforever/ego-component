package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
)

/**
1.新建一个数据库叫test
2.执行以下example，export EGO_DEBUG=true && go run main.go --config=config.toml
*/
type User struct {
	Id       int    `gorm:"not null" json:"id"`
	Nickname string `gorm:"not null" json:"name"`
}

func (User) TableName() string {
	return "user2"
}

func main() {
	err := ego.New().Invoker(
		openDB,
		testDB,
	).Run()
	if err != nil {
		elog.Error("startup", elog.Any("err", err))
	}
}

var gormDB *egorm.Component

func openDB() error {
	gormDB = egorm.Load("mysql.test").Build()
	models := []interface{}{
		&User{},
	}
	gormDB.SingularTable(true)
	gormDB.AutoMigrate(models...)
	gormDB.Create(&User{
		Nickname: "ego",
	})
	return nil
}

func testDB() error {
	var user User
	err := gormDB.Where("id = ?", 100).Find(&user).Error
	elog.Info("user info", elog.String("name", user.Nickname))
	return err
}
