// _Slices_는 배열보다 좀 더 강력한 시퀀스 인터페이스를 제공하는 Go의 핵심적인 데이터 타입입니다.

package main

import "fmt"

func main() {

	// 배열과 다르게, 슬라이스는 원소의 갯수가 아닌 포함된 원소들로만 작성됩니다.
	//  길이가 0인 빈 배열을 만들기 위해선 내장 함수 `make`를 사용하면 됩니다.
	//  제로값으로 초기화된 길이가 3인 문자열 배열을 만듭니다.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 배열처럼 값을 저장하고 사용할 수 있습니다.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len`은 예상대로 슬라이스의 길이를 반환합니다.
	fmt.Println("len:", len(s))

	// 이러한 기본 연산에 이어 슬라이스는 배열보다 더 풍부한 기능들을 지원합니다.
	//  그 중 하나는 새로운 값이 추가된 슬라이스를 반환하는 `append` 함수입니다.
	//  주의할건 새로운 슬라이스를 사용하기위해선 append로부터 반환되는 값을 사용해야합니다.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 슬라이스는 또한 복사도 가능합니다.
	//  `s`와 같은 길이를 갖는 빈 슬라이스 `c`를 생성하고 `s`를 `c`로 복사합니다.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 슬라이스는 `slice[low:high]`로 쓸 수 있는 "slice(자르기)" 연산을 지원합니다.
	//  예를 들어, 다음 연산을 통해 `s[2]`, `s[3]`와 `s[4]`로 이루어진 슬라이스를 얻을 수 있습니다.`
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 다음처럼 쓰면 `s[5]`까지 자릅니다. (단, 마지막 원소는 포함하지 않습니다.)
	l = s[:5]
	fmt.Println("sl2:", l)

	// 또한 다음처럼 쓰면 `s[2]`부터 자르게 됩니다. (이 때는 첫원소를 포함합니다.)
	l = s[2:]
	fmt.Println("sl3:", l)

	// 또한 한 줄로 슬라이스 선언 및 초기화를 할 수도 있습니다.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 슬라이스는 다차원 데이터를 구성할 수도 있습니다
	//  다차원 배열과 다르게 다중 슬라이스의 내부 슬라이스들은 가변적 길이를 가질 수 있습니다.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
