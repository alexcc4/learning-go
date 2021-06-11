package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"week2/dao"
)

func GetUserByName(n string) interface{} {
	user, err := dao.GetUser(n)

	// handler all error in this layer
	// return nil or user struct
	if err != nil {
		// handler user not exist
		if errors.Cause(err) == dao.RecordNotFoundError{
			fmt.Println(err)
			return nil
		}

		// other error
		log.Fatalln(err)
		return nil
	}
	return user
}

func main() {
	u := GetUserByName("aaa")
	fmt.Printf("%+v", u)
}
