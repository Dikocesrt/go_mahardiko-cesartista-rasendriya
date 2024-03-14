package main

import (
	"fmt"
)

type user struct {
	Email    string
	Password string
}

type userRepo struct {
	DB []user
}

//penamaan object dari userRepo dan user tidak jelas hanya r dan u saja dan juga receiver r userRepo seharusnya menggunakan pointer menjadi r *userRepo
func (r *userRepo) Register(u user) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("register failed")
	}

	r.DB = append(r.DB, u)
}

func (r *userRepo) Login(u user) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("login failed")
	}

	//penamaan variable yang menampung hasil perulangan pada setiap object pada r.DB tidak jelas hanya us saja
	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token"
		}
	}

	return ""
}

//tidak terdapat fungsi main
func main() {
	user1 := user{
		"dikocesrt@gmail.com",
		"dikodikodiko",
	}

	repo := userRepo{}

	repo.Register(user1)

	fmt.Println(repo.Login(user1))
}