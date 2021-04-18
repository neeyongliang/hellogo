package main

import (
	"fmt"
	"os"
)

/*
 * 函数学生管理系统，写一个系统能够查看、新增删除学生
 */
type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student
)

func showAllStudents() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d，姓名：%s\n", k, v.name)
	}
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Printf("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Printf("请输入学生的姓名：")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	allStudent[id] = s
}

func deleteStudent() {
	var (
		stuId int64
	)
	fmt.Printf("请输出要删除的用户ID：")
	fmt.Scanln(&stuId)
	delete(allStudent, stuId)
}

func main() {
	allStudent = make(map[int64]*student, 50)
	// 1. 打印菜单
	fmt.Println("欢迎使用学生管理系统:")
	fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 推出
	`)
	for {
		fmt.Printf("请输入你的操作：")
		// 2. 等待用户选择
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择的是：%d", choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("你说什么？？？？")
		}
	}
}
