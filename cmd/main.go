package main // Go의 기본 패키지

import (
	"fmt" // fmt는 Go의 Formatting을 관리하는 모듈
	"hello"
	"os"
)

func main() { // 만약에 모듈을 Export하고 싶다면 함수의 첫 글자를 대문자로 해야한다.
	fmt.Println(hello.Say((os.Args[1:])))
}
