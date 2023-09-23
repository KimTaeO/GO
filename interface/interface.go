package main

import "fmt"

type User struct {
	name        string
	description string
	age         int
}

type Teacher struct {
	name string
	age  int
}

// 위와 같이 인터페이스를 선언 가능하다
// 메서드명(파라미터) 반환타입 형식으로 선언 가능하다
// 인터페이스의 기본값은 nil이다
type Abstract interface {
	updateName(name string)
}

func (t *Teacher) updateName(name string) {
	t.name = name
}

func (u *User) updateName(name string) {
	u.name = name
}

func updateName(a Abstract, name string) {
	if a != nil {
		a.updateName(name)
	}
}

// 1. 메서드는 반드시 메서드명이 있어야 한다.
// 2. 매개 변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다.
// 3. 인터페이스에서는 메스드 구현을 포함하지 않는다.
func main() {

	u := User{
		name: "gopher",
		age:  12,
	}

	fmt.Println(u)

	updateName(&u, "golang")

	fmt.Println(u)
}
