package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func main(){

	type User struct {
		Name string 	`validate:"min=6,max=10"`
		Age  int    	`validate:"min=1,max=100"`
		Email string 	`validate:"emailOrNull"`
		EAge int 		`validate:"eqfield=Age"`
		Mobile string 	`validate:"mobileOrNull"`
		QQ string 		`validate:"QQOrNull"`
	}

	validate := New()

	u1 := User{
		Name: "lidajun",
		Age: 18,
		Email:"123@gg.co",
		//Email:"",
		EAge:18,
		Mobile:"18844445555",
		QQ:"78444",
	}
	err := validate.Struct(u1)
	fmt.Println(err)

	fmt.Println("==========================")
	u2 := User{Name: "dj", Age: 101}
	err = validate.Struct(u2)
	fmt.Println(err)

}


var (
	qq_regex = `^[1-9][0-9]{4,11}$`
	tel_regex = `^[1][3-9][0-9]{9}$`
)

var ValidateV10 *validator.Validate

func init() {
	ValidateV10 = validator.New()
	err1 := ValidateV10.RegisterValidation("mobile", CheckMobile)
	if err1 != nil{
		panic("validate_v10 CheckMobile err:" + err1.Error())
	}
	err11 := ValidateV10.RegisterValidation("mobileOrNull", CheckMobileV2)
	if err11 != nil{
		panic("validate_v10 CheckMobile err:" + err11.Error())
	}
	err2 := ValidateV10.RegisterValidation("QQ", CheckQQ)
	if err2 != nil{
		panic("validate_v10 CheckQQ err:" + err2.Error())
	}
	err22 := ValidateV10.RegisterValidation("QQOrNull", CheckQQV2)
	if err22 != nil{
		panic("validate_v10 CheckQQ err:" + err22.Error())
	}
	err3 := ValidateV10.RegisterValidation("emailOrNull", CheckEMailV2)
	if err3 != nil{
		panic("validate_v10 CheckQQ err:" + err3.Error())
	}
}

func New() *validator.Validate {
	return ValidateV10
}

// CheckMobile 校验手机号码
func CheckMobile(no validator.FieldLevel) bool {
	value := no.Field().String()
	mobileMatch, err := regexp.MatchString(tel_regex, value)
	if err != nil || !mobileMatch {
		return false
	}
	return mobileMatch
}

// CheckMobile 校验手机号码
func CheckMobileV2(no validator.FieldLevel) bool {
	value := no.Field().String()
	if len(strings.TrimSpace(value)) == 0 {
		return true
	}
	mobileMatch, err := regexp.MatchString(tel_regex, value)
	if err != nil || !mobileMatch {
		return false
	}
	return mobileMatch
}

// CheckQQ 校验QQ号码
func CheckQQ(no validator.FieldLevel) bool {
	value := no.Field().String()
	qqMatch, err := regexp.MatchString(qq_regex, value)
	if err != nil || !qqMatch {
		return false
	}
	return qqMatch
}

// CheckQQ 校验QQ号码
func CheckQQV2(no validator.FieldLevel) bool {
	value := no.Field().String()
	if len(strings.TrimSpace(value)) == 0 {
		return true
	}
	qqMatch, err := regexp.MatchString(qq_regex, value)
	if err != nil || !qqMatch {
		return false
	}
	return qqMatch
}

// CheckQQ 校验QQ号码
func CheckEMailV2(no validator.FieldLevel) bool {
	value := no.Field().String()
	if len(strings.TrimSpace(value)) == 0 {
		return true
	}
	err := ValidateV10.Var(value, "email")
	if err != nil{
		return false
	}
	return true
}