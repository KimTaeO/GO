# Variable(변수)

* ## Go의 변수는 변수에 담긴 값에 대한 메모리 위치의 이름이나 주소를 가지고 있다 Go는 컴파일 언어이고, 숫자는 기계 표현에 표시되고 이 작업은 매우 Go를 빠르게 해준다

> ## 숫자가 무언인지에 대한 추상화가 포합되어 있지 않다

## 변수 선언
```Go 
    // 일반적인 변수 선언
    var a int = 1

    // 타입 추론 사용 시
    var b = 2

    // 여러 개의 변수를 한번에 선언할 때
    var (
        c = 3
        d = 4.5
    )

    // func 안에서만 사용할 수 있는 축약형 표현
    e := 6
```
> ### Go는 한번 타입을 결정하면 해당 타입으로만 사용해야 한다

## 변수의 초기화

* ### 모든 변수는 선언 시점에서 초기화된다

* ### 초기화되지 않는 변수는 없으며, 변수가 초기화되지 않아 생길 수 있는 모든 문제를 막기 위해 이렇게 설계되었다

* ### 모든 변수는 0으로 초기화되며, 문자열은 빈 문자열로 초기화됩니다