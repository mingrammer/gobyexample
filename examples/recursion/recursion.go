// Go는 <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>를 지원합니다.
//  여기에 고전적인 팩토리얼 예시가 있습니다.

package main

import "fmt"

// `fact` 함수는 베이스 케이스인 `fact(0)`에 도달할 때까지 자기자신을 호출합니다.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
