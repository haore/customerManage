package main

import (
	"fmt"
)

type customerViews struct {
	key  string
	loop bool
}

func (acc *customerViews) MainMenu() {
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
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
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
		fmt.Println(acc.loop)
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
	customerViews.MainMenu()
}
