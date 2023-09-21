package main

import "fmt"

// 배열은 포인터 타입이 아니므로 새로운 인스턴스를 생성해 복사하기 때문에 원본 배열의 값이 바뀌지 않는다
func changeArr(arr [5]int) {
	arr[0] = 100
}

// 슬라이스는 값을 복사하는 것이 아닌 메모리 주소값을 복사하기 때문에 원본 슬라이스 값이 바뀐다
func changeSlice(slice []int) {
	slice[0] = 100
}

func main() {
	// 슬라이스는 Golang에서 제공하는 동적 배열 타입(배열을 가리키는 포인터 타입)이다

	// 배열 선언
	arr := [5]int{1, 2, 3, 4, 5}
	// 슬라이스 선언
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(slice)

	changeArr(arr)
	changeSlice(slice)

	fmt.Println(arr)
	fmt.Println(slice)

	// 현재 슬라이스에서 사용중인 길이
	fmt.Println(len(slice))
	// 현재 슬라이스가 사용중인 길이 + 비어있는 공간의 길이
	fmt.Println(cap(slice))

	// make 함수를 통해 슬라이스를 만들 수도 있다
	// 슬라이스 타입, 사용할수 있는 인자 개수, 슬라이스의 크기가 make함수의 파라미터이다
	slice2 := make([]int, 2, 10)

	fmt.Println(slice2)
}
