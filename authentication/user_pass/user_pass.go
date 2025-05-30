package user_pass

import (
	"gorm.io/gorm"
)

func ValidateUserPass(username string, password string, connection *gorm.DB) bool {

	type Users struct {
		Id int
	    Username string
		Password string
	}

	user := Users{Username: username,Password: password}
    var rows []map[string]any
    connection.Where("username=? AND password=?",username,password).First(&user).Scan(&rows)
    return len(rows)==1
}

