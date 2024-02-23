package user

import (
	"errors"
	"regexp"
)

type User struct {
	ID             int
	Name           string  `json:"name"`
	Age            int     `json:"age"`
	DocumentNumber string  `json:"document_number"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	Address        Address `json:"address"`
}
type Address struct {
	ZipCode      string `json:"zip_code"`
	Country      string `json:"country"`
	State        string
	City         string
	Neighborhood string
	Street       string
	Number       string `json:"number"`
	Complement   string `json:"complement"`
}

func nameValidation(name string) error {
	isValid, err := regexp.MatchString("[A-Z][a-zA-Z]{2,} [A-Z][a-zA-Z ]+", name)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("the name is not valid")
	}
	return nil
}

func ageValidation(age int) error {
	if age < 18 || age > 100 {
		return errors.New("the age is not valid")
	}
	return nil
}

func documentValidation(documentNumber string) error {
	isValid, err := regexp.MatchString("^\\d{3}\\.\\d{3}\\.\\d{3}-\\d{2}$", documentNumber)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("the CPF is not valid")
	}
	return nil
}

func emailValidation(email string) error {
	isValid, err := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", email)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("the email is not valid")
	}
	return nil
}

func phoneValidation(phone string) error {
	isValid, err := regexp.MatchString("^\\+55[ ]\\d{2}[ ]9\\d{4}[ ]\\d{4}$", phone)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("the phone number is not valid")
	}
	return nil
}
