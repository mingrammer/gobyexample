// Go에서 `if`, `else` 분기는 직관적입니다.

package main

import "fmt"

func main() {

	// 여기에 기본적인 예시가 있습니다.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// else 없이 `if`문을 사용할 수 있습니다.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 조건절 전에 명령문이 위치할 수 있습니다. 이 명령문에서 선언된 모든 변수들은 모든 분기문에서 사용 가능합니다.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Go에서 주의할 건 조건절엔 괄호가 필요없지만, 중괄호는 필수입니다.
