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

		fmt.Println(a)
	}
}

func toTeacher(a Abstract) {
	if a != nil {
		a.updateName("it is abstract")

		// 타입 변환이 되었는지 체크하는 방법
		// 타입을 변환할 때 두번째 인자는 타입을 변환했을때 성공 여부를
		if t, ok := a.(*Teacher); ok {
			t.updateName("now it is teacher")
		}
	}
}

// 빈 인터페이스는 모든 타입을 가질 수 있다
func Print(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println(v, " is int")
	case float64:
		fmt.Println(v, " is float64")
	case string:
		fmt.Println(v, " is string")
	default:
		fmt.Printf("Not supported %T:%v\n", v, v)
	}
}

// 1. 메서드는 반드시 메서드명이 있어야 한다.
// 2. 매개 변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다.
// 3. 인터페이스에서는 메스드 구현을 포함하지 않는다.
func main() {

	u := &User{
		name: "gopher",
		age:  12,
	}

	fmt.Println(u)

	updateName(u, "golang")

	t := &Teacher{
		name: "teacher",
		age:  27,
	}

	fmt.Println(t)

	toTeacher(t)

	Print(10)
	Print(3.14)
	Print("Hello, world!")
	Print(t)
}
