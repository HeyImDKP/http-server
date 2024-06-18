package lib

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	Username     string
	Password     string
	Hostname     string
	DatabaseName string
	DB           *gorm.DB
}

func (mySQLDB *MySQLDB) ConnectToDB() error {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		mySQLDB.Username,
		mySQLDB.Password,
		mySQLDB.Hostname,
		mySQLDB.DatabaseName,
	)

	var err error
	mySQLDB.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return mySQLDB.DB.AutoMigrate(&User{})
}

func (mySQLDB MySQLDB) DisconnectFromDB() error {
	sqlDB, err := mySQLDB.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (mySQLDB MySQLDB) CreateUser(user *User) error {
	result := mySQLDB.DB.Create(user)
	return result.Error
}

func (mySQLDB MySQLDB) GetUserByID(userID uint, user *User) error {
	result := mySQLDB.DB.First(&user, userID)
	return result.Error
}

func (mySQLDB MySQLDB) UpdateUser(user *User) error {
	result := mySQLDB.DB.Save(user)
	return result.Error
}

func (mySQLDB MySQLDB) DeleteUser(user *User) error {
	result := mySQLDB.DB.Delete(user)
	return result.Error
}

func (mySQLDB MySQLDB) GetUserByUsername(username string, user *User) error {
	result := mySQLDB.DB.Where("username = ?", username).First(&user)
	return result.Error
}
