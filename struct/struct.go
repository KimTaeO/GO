package main

import "fmt"

// 여러 타입의 필드를 묶어서 제공 가능
type User struct {
	name string
	age  int
}

// 구조체가 다른 구조체 포함 가능 Nested Struct라고 부름
type Class struct {
	student  User
	classNum int
}

// 필드명을 사용하지 않고 다른 구조체 포함 가능 Embedded field라고 부름
type School struct {
	Class
	name string
}

func main() {
	var u User
	u.name = "Golang"
	u.age = 10

	fmt.Println(u)
	fmt.Printf("name: %s, age: %d\n", u.name, u.age)

	// 구조체 선언 시 필드 순서로 값을 대입하여 선언 가능
	var u2 User = User{"GoGoGo", 10}
	// 필드명 지정하여 초기화도 가능
	var u3 User = User{name: "Gopher", age: 10}

	fmt.Println(u2)
	fmt.Println(u3)

	var class Class = Class{
		student:  u,
		classNum: 3,
	}

	// nested struct에 접근할 때에는 구조체명.중복된 구조체.중복된구조체의 필드로 접근해야 한다
	fmt.Printf("user.name: %s, user.age: %d, classNum: %d\n", class.student.name, class.student.age, class.classNum)

	var school School = School{
		Class: class,
		name:  "this is school",
	}
	// embedded field에 접근할 때에는 구조체명.중복된 구조체의 필드로 접근할 수 있다
	// 만약에 필드명이 중복된다면 nested struct와 같은 방법으로 참조하면 된다
	fmt.Printf("student.name: %s student.age: %d, classNum: %d, name: %s\n", school.student.name, school.student.age, school.classNum, school.name)

}
