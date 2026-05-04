package main

import (
	"fmt"
	"reflect"
)

type SysUser struct {
	ID        int    `json:"id"` //тег - метадата в данном случае это значчит что переводе в джейсон будут использоваться именно эти названия
	FirstName string `json:"firstname"`
	Email     string `json:"email" db:"email"`
}

func main() {

	u := SysUser{1, "sdfvc", "sdfvcwsdc"}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Поле %s, json %s. db : %s", field.Name, field.Tag.Get("json"), field.Tag.Get("db")) //тег паолучить
	}
	fmt.Println()
}
