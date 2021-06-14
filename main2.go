package main

import "fmt"

func superAdd(numbers ...int) int {
	// 첫번째는 인덱스, 두번째에 Item을 할당.
	total := 0
	for i, number := range numbers {
		fmt.Println(i, number)
		total += number
	}

	return total
}

func canIDring(age int) bool {
	// if문 안에서 변수를 생성해 사용할 수 있다.
	// 이렇게 되면, if문 안에서만 사용하는 변수라는 의도를 명확히 전달할 수 있다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7))

	fmt.Println(canIDring(16))

}
