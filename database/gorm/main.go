package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {


	//自定义数据库日志
	loger := logger.New(
			log.New(os.Stdout,"\r\n",log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Info,
			},
		)

	//open := mysql.Open("root:root@tcp(localhost:3306)/future?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:root@tcp(localhost:3306)/future?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                    "mysql",
		ServerVersion:                 "5.7",
		DSN:                           dsn,
		DefaultStringSize:             256,
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: loger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})

	if err != nil {
		panic(err)
	}

	//tx := db.Session(&gorm.Session{})
	//fmt.Println(db)
	//tx.First(&struct {
	//	A int
	//	B int
	//}{
	//
	//})

	// Create
	db.Create(&Product{Code: "D1", Price: 100})

	// Read
	var product Product
	db.First(&product, 8) // 根据整型主键查找
	db.First(&product, "code = ?", "D1") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段

	//如果想修改非空的，是为了避免将其他的也更新了默认值，所以需要在设置0值或者空值的时候
	db.Model(&product).Updates(Product{Price: 200, Code: "F1"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F1"})

	fmt.Println(product)
	// Delete - 删除 product
	db.Delete(&product, 1)

}