package main

import "fmt"

type MyInt int    // 自定义
type YouInt = int // 别名

type Person struct {
	name   string
	gender string
}

// 构造函数：约定都以 new 开头
// 返回的是结构体还是指针，取决于当前结构体的大小，如果成员比较多就返回指针
// 建议统一返回指针
func newPerson(name string, gender string) *Person {
	return &Person{
		name:   name,
		gender: gender,
	}
}

type Dog struct {
	name string
}

// 方法是指作用域特定类型的函数
// 接受者表示的是调用的具体类型，多用类型名小写字母表示
// 类似于 java 的this，或者 python 的 self
func (d Dog) Wang() {
	fmt.Println("汪汪汪～")
}

func main() {
	var p Person
	p.name = "AAA"
	p.gender = "male"

	var p2 = new(Person)
	(*p2).name = "BBB"
	(*p2).gender = "female"

	var p3 = &Person{
		"CCC",
		"male",
	}
	fmt.Println(p3)

	var p4 = newPerson("DDD", "female")
	fmt.Println(p4)

	var dog = new(Dog)
	dog.Wang()
}
