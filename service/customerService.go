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

//func (acc *CustomerService) Add(account model.Customer) {
//	acc.customerNum ++
//	acc.customers = append(acc.customers, account)
//}
//func (acc *CustomerService) Delete(id int) {
//	customers := acc.customers
//	customers = append(customers[0:id], customers[id+1:]
//
//}
