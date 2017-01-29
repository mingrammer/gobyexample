// 문자열로부터 숫자를 파싱하는건 많은 프로그램에서 기본적이지만 일반적인 작업입니다.
//  Go에서 이를 어떻게 하는지 살펴봅시다.

package main

// 내장 패키지 `strconv`는 숫자 파싱을 제공합니다.
import "strconv"
import "fmt"

func main() {

	// `ParseFloat`를 사용할 때, `64`는 얼마나 많은 비트의 정밀도를 파싱할지를 알려줍니다.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt`에서, `0`은 문자열로부터 베이스값을 추론함을 의미합니다.
	//  `64`를 사용하려면 결괏값이 64 비트에 적합해야합니다.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt`는 16진법 수를 인식합니다.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 'ParseUint' 또한 사용할 수 있습니다.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi`는 기본적인 10진수 `int` 파싱을 위한 편리한 함수입니다.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 파싱 함수는 잘못된 입력값이 들어오면 에러를 반환합니다.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
