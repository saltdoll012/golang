// 메인 패키지
package main

import "calc"
import "fmt"

// 메인 함수
func main() {
	fmt.Println(calc.Sum(1, 2))	// calc 패키지의 Sum 함수 사용
	//fmt.Println(calc.multiply(1, 2))	// 사용할 수 없어서 오류가 납니다.
}

