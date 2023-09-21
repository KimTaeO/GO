package main

import "fmt"

// 배열의 일부를 잘라내어 슬라이스를 만들어 내는 것을 슬라이싱이라 한다
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 배열[startIdx:endIdx] 형식으로 사용하며
	// startIdx부터 endIdx-1번 인덱스의 인자까지 슬라이스로 만든다
	slice := arr[1:5]

	fmt.Println(slice)

	// 또한 슬라이스를 슬라이싱 하는 것도 가능하다
	slice2 := slice[3:5]

	fmt.Println(slice2)

	// 이와 같이 startIndex endIndex를 생력하여 적을 수 있고
	// 각각 생략되면 처음 인덱스, 마지막 인덱스를 의미한다
	// 모두 생략한다면 배열을 슬라이스로 바꿀 수 있다
	slice3 := arr[3:]
	slice4 := arr[:5]
	slice5 := arr[:]

	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	// 슬라이스는 주소를 가리키기 때문에 이런 것도 가능하다
	slice6 := arr[:5]
	slice7 := slice6[:9]

	fmt.Println(slice7) // [1 2 3 4 5 6 7 8 9 10]
}
