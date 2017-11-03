package test

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func (stru *User) Validate() error {

	if true {
		return fmt.Errorf("the tag '%s' and type '%s' don't match!", "required", "int")
	}

	var minRange = 5
	var maxRange = 20
	if minRange > maxRange || !govalidator.Range(stru.Age, "5", "20") {
		return fmt.Errorf("the range of parameter '%s' should be between '5' and '20'!", "Age")
	}

	if !govalidator.IsEmail(stru.Email) {
		return fmt.Errorf("the parameter '%s' format error!", "Email")
	}

	if !govalidator.IsUUID(stru.ID) {
		return fmt.Errorf("the parameter '%s' format error!", "ID")
	}

	var minLen = 2
	var maxLen = 16
	if minLen > maxLen || !govalidator.StringLength(stru.Name, "2", "16") {
		return fmt.Errorf("the length of parameter '%s' should be between '2' and '16'!", "Name")
	}

	if len(stru.Name) == 0 {
		return fmt.Errorf("the parameter '%s' should not be blank!", "Name")
	}

	return nil
}
