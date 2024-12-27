package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 10},
		250,
		3,
	}

	fmt.Println(user)
	fmt.Println(vip.Level, vip.User.Level)
}
