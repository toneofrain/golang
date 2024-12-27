package main

import "fmt"

type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.Id, user.Age)
	fmt.Printf("유저: %s ID: %s 나이: %d\n", vip.UserInfo.Name, user.Id, user.Age)
}
