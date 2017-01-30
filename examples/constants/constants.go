// Go는 문자, 문자열, 부울 그리고 숫자값 _상수_를 지원합니다.

package main

import "fmt"
import "math"

// `const`로 상수값을 선언합니다.
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const`문은 `var`문과 동일하게 사용할 수 있습니다.
	const n = 500000000

	// 상수 표현식은 임의의 정밀도로 산술 연산을 수행합니다.
	const d = 3e20 / n
	fmt.Println(d)

	// 숫자 상수는 명시적 캐스팅등으로 타입이 주어지기 전까진 타입을 가지지 않습니다.
	fmt.Println(int64(d))

	// 숫자는 변수 할당이나 함수 호출과 같은 컨텍스트에서 사용하여 타입을 부여할 수 있습니다.
	//  예를 들면, `math.Sin`은 `float64`를 기대합니다.
	fmt.Println(math.Sin(n))
}
