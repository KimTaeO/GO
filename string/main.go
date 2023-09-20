package main

import "fmt"

func main() {
	s := "élite"

	// 논리적으로 배열의 인자가 5개지만 len 함수를 돌리면 UTF-8 인코딩 스타일 때문에 유니코드가 2바이트를 차지하여 물리적인 길이는 6이 나온다
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))

	b := []byte(s)

	// UTF-8 이 유니코드를 인코딩하여 띄워주기 때문에 배열의 인자가 5개가 아닌 6개이다
	fmt.Printf("%8T %[1]v %d\n", b, len(b))

	s := "the quick brown fox"

	a := len(s)
	b := s[:3]                  // 기존 문자열과 같은 위치의 문자열 사용
	c := s[4:9]                 // 기존 문자열과 같은 위치의 문자열 사용
	d := s[:4] + "slow" + s[9:] // 새로운 문자열을 메모리에 할당

	s[5] = 'a' // Syntax err
	s += "es"  // 기존 문자열을 불러와 es를 이어붙임

}
