package app

import "github.com/asaskevich/govalidator"

type user struct {
	Username string `valid:"required"`
	Email    string `valid:"email,required"`
	Password string `valid:"minstringlength(6),required"`
}

type userLogin struct {
	Email    string `valid:"email,required"`
	Password string `valid:"minstringlength(6),required"`
}

func ValidateUserInput(username string, email string, password string) bool {
	user := user{Username: username, Email: email, Password: password}
	result, _ := govalidator.ValidateStruct(user)
	return result
}

func ValidateUserLogin(email string, password string) bool {
	user := userLogin{Email: email, Password: password}
	result, _ := govalidator.ValidateStruct(user)
	return result
}
