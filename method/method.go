package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 구조체는 임베디드 타입을 필드로 가질 수 있고 임베디드 필드는 메서드를 가지고 있을 수 잇다
type Group struct {
	Person
	grade int
}

// 타입에 종속된 메서드를 통해 객체 지향 프로그래밍 구현
// 단순히 값을 불러오거나 객체의 동일성을 보장하지 않아도 될 대에는 값 타입 메서드를 사용해도 된다
func (p Person) greeting() string {
	return "Hello, " + p.name
}

// 나이가 증가해도 같은 사람이여야 하기 때문에 포인터 타입 메서드를 사용해야 한다
func (p *Person) addAge() {
	p.age += 1
}

func main() {
	// 함수는 타입과 독립적이지만 메서드는 타입에 종속됨

	p := Person{
		name: "스프링컨트리뷰터김희망",
		age:  18,
	}

	fmt.Println(p.greeting())

	fmt.Println(p)

	g := Group{
		Person: p,
		grade:  1,
	}

	g.addAge()

	fmt.Println(g)
}
