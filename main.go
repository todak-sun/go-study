package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	// a int, b int를 위와 같이 줄여 작성할 수 있다.
	return a * b
}

// 리턴을 2개하는 함수
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// naked return. 리턴을 꼭 명시하지 않아도 되는경우
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 함수가 종료된 후 defer를 호출한다.
func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("before defer")
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// constant
	const name string = "todak"
	const flag bool = true

	repeatMe("wow", "is", "time")

	name2 := "wow"
	// 위는 이와 같이 바꿔 쓸 수 있다
	// 함수 안에서만 위와 같이 작성할 수 있다.
	// => var name2 string = "wow"
	fmt.Println(name2)

	fmt.Println(multiply(4, 2))

	totalLength, upperName := lenAndUpper("WOW")
	fmt.Println(totalLength, upperName)

	fmt.Println(lenAndUpper2("todak"))
	fmt.Println(lenAndUpper3("todaktodak"))

}
