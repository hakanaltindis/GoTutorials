package main

import "fmt"

type User struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

type Payment struct {
	Salary float64
	Bonus  float64
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

// Methods
func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) GetUserName() string {
	return u.UserName
}

func (u User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}

func main() {
	// Generate user v1
	user1 := &User{
		FirstName: "Hakan",
		LastName:  "Altındiş",
		UserName:  "hakana",
		Age:       32,
		Pay: &Payment{
			Salary: 4320,
			Bonus:  725,
		},
	}

	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetPayment())

	// Generate user v2
	user2 := NewUser()
	user2.FirstName = "Hakan"
	user2.LastName = "Altındiş"
	user2.UserName = "hakanaltindis"
	user2.Age = 32
	user2.Pay.Salary = 5250
	user2.Pay.Bonus = 750

	fmt.Println(user2.GetUserName())
	fmt.Println(user2.GetPayment())

}
