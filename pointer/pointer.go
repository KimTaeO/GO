package main

import "fmt"

type Item struct {
	name  *string
	age   int
	price int
}

// 함수에서 포인터를 이용하여 구조체의 값 변경
func Changename(item *Item, name string) {
	item.name = &name
}

func main() {
	// 포인터 변수 선언 방식
	var a *int

	// 할당되지 않은 포인터의 기본 값은 nild
	fmt.Println(a)

	b := 10
	// 변수의 주솟값에 접근하려면 &연산자를 앞에 붙여주면 된다
	a = &b

	fmt.Printf("%p\n", a)

	// 포인터 변수의 값에 접근하려면 * 연산자를 앞에 붙여 참조한다
	fmt.Printf("%d\n", *a)

	// new 내장 함수를 사용하여 구조체의 인스턴스를 생성할 수 있다
	// 이때 변수의 타입은 구조체 포인터이다
	p := new(Item)

	fmt.Println(*p)

	// Golang에도 Garbage Collector가 있어 인스턴스를 참조하는 변수가 사라지면 자동으로 제거된다
	// 함수가 끝나더라도 Escape Analysis 기능을 통해 함수를 벗어나는 변수들을 검사하여 힙 메모리에 저장한다
}
