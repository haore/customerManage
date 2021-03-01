package service

import "customerManage/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 18,
		"18658688569", "453738632@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (acc *CustomerService) List() []model.Customer {
	return acc.customers
}

func (acc *CustomerService) Add(account model.Customer) bool {
	acc.customerNum ++
	account.Id = acc.customerNum
	acc.customers = append(acc.customers, account)
	return true
}

func (acc *CustomerService) Delete(id int) bool {
	index := acc.FindById(id)
	if index == -1 {
		return false
	} else {
		acc.customers = append(acc.customers[:index], acc.customers[index+1:]...)
		return true
	}
}

func (acc *CustomerService) Update(index int, name string, gender string,
	age int, phone string, email string) bool {
	//index := acc.FindById(id)
	//if index == -1 {
	//	return false
	//}
	if name != "" {
		acc.customers[index].Name = name
	}
	if gender != "" {
		acc.customers[index].Gender = gender
	}
	if age > 0 {
		acc.customers[index].Age = age
	}
	if phone != "" {
		acc.customers[index].Phone = phone
	}
	if email != "" {
		acc.customers[index].Email = email
	}
	return true
}

func (acc *CustomerService) FindById(id int) int {
	customers := acc.customers
	for index, value := range customers {
		if value.Id == id {
			return index
		}
	}
	return -1

}
//func (acc *CustomerService) Delete(id int) {
//	customers := acc.customers
//	customers = append(customers[0:id], customers[id+1:]
//
//}
