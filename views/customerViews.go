package main

import (
	"customerManage/model"
	"customerManage/service"
	"fmt"
)

type customerViews struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (acc *customerViews) delete() {
	fmt.Println("----------------删除客户---------------")
	fmt.Println("请选择要被删除用户的id(-1)退出")
	var accId int = -1
	fmt.Scanln(&accId)
	if accId == -1 {
		return
	}
	choice := ""
	for {
		fmt.Println("确认是否删除（Y/N）：")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			if acc.customerService.Delete(accId) {
				fmt.Println("-----------删除成功-----------")
			} else {
				fmt.Println("-----------删除失败,id号不存在-----------")
			}
			break
		} else if choice == "n" || choice =="N" {
			fmt.Println("-----------取消删除-----------")
			break
		} else {
			fmt.Println("请输入（y/n）")
		}
	}
}

func (acc *customerViews) add() {
	fmt.Println("----------------添加客户---------------")
	fmt.Println("姓名：")
	var name string
	fmt.Scanln(&name)
	fmt.Println("性别：")
	var gender string
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	var age int
	fmt.Scanln(&age)
	fmt.Println("电话：")
	var phone string
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	var email string
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if acc.customerService.Add(customer) {
		fmt.Println("----------------添加客户成功---------------")
	} else {
		fmt.Println("----------------添加客户失败---------------")
	}
}

func (acc *customerViews) list() {
	fmt.Println("----------------客户列表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	customers := acc.customerService.List()
	for _, value := range customers {
		value.GetCustomerInfo()
	}
	fmt.Println("---------------客户列表完成--------------\n")
}

func (acc *customerViews) mainMenu() {
	for {
		fmt.Println("----------------客户信息管理软件---------------")
		fmt.Println("                1.添加客户")
		fmt.Println("                2.修改客户")
		fmt.Println("                3.删除客户")
		fmt.Println("                4.客户列表")
		fmt.Println("                5.退出")
		fmt.Println("请输入（1-5）: ")
		fmt.Scanln(&acc.key)
		switch acc.key {
		case "1":
			acc.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			acc.delete()
		case "4":
			acc.list()
		case "5":
			var getLoop string
			for {
				fmt.Println("确定是否退出(y/n)")
				fmt.Scanln(&getLoop)
				if getLoop == "y" || getLoop == "y" || getLoop == "n" || getLoop == "N" {
					acc.loop = true
					break
				} else {
					fmt.Println("请输入y或者n继续")
				}
			}
		default:
			fmt.Println("请输入正确的序号")
		}
		if acc.loop {
			fmt.Println("退出客户信息管理软件")
			break
		}
	}
}

func main() {
	customerViews := customerViews{
		key:  "",
		loop: false,
	}
	customerViews.customerService = service.NewCustomerService()
	customerViews.mainMenu()
}
