package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		id,
		name,
		gender,
		age,
		phone,
		email,
	}
}

func (acc *Customer) GetCustomerInfo() {
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", acc.Id, acc.Name, acc.Gender,
		acc.Age, acc.Phone, acc.Email)
}
