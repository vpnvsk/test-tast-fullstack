package models

import "errors"

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type UserUpdate struct {
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Age     *int    `json:"age"`
}

func (i UserUpdate) Validate() error {
	if i.Name == nil && i.Age == nil && i.Surname == nil && *i.Age < 1 {
		return errors.New("invalid request")
	}
	return nil
}
