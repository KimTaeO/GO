package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 배열을 슬라이싱한 슬라이스의 cap은 배열의 maxIndex값이다
	// 슬라이싱 할 떄 위와 같은 형식으로 cap사이즈를 조절 가능하다
	slice := arr[1:5:8]

	// slice2 := slice[:10] Err!

	slice2 := slice[:7]
	fmt.Println(slice)

	fmt.Println(slice2)

	// append 함수를 통해 슬라이스에 인자 추가 가능
	// 기존 슬라이스 공간이 충분하지 않으면 새로운 메모리에 슬라이스 복사 후 추가함
	slice3 := append(slice, 6)

	fmt.Println(slice3)

	// Go에서 copy 함수는 다음과 같이 사용 가능하다
	// copy(삽입할 슬라이스, 복사할 슬라이스)

	// 슬라이스 정렬
	// Go에서 기본으로 제공하는 패키지 존재
	slice4 := []int{2, 5, 6, 4}
	sort.Ints(slice4)
	fmt.Println(slice4)

	// 추가적으로 구조체 슬라이스도 정렬 가능
}
