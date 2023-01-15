package main

import "fmt"

type User interface {
	getPermissions() int // 1-5
	getName() string
}

type Admin struct {
	name string
}

func (this Admin) getPermissions() int {
	return 5
}

func (this Admin) getName() string {
	return this.name
}

type Editor struct {
	name string
}

func (this Editor) getPermissions() int {
	return 3
}

func (this Editor) getName() string {
	return this.name
}

func auth(user User) string {
	permissions := user.getPermissions()
	if permissions > 4 {
		return user.getName() + " tiene permisos de administrador"
	} else if permissions == 3 {
		return user.getName() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	admin := Admin{"Alex"}
	fmt.Println(auth(admin))

	editor := Editor{"Uriel"}
	fmt.Println(auth(editor))

	users := []User{Admin{"Uriel 2"}, Editor{"Alex 2"}}
	for _, user := range users {
		fmt.Println(auth(user))
	}
}
