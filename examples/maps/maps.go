// _Maps_는 Go의 내장 [연관 데이터 타입(associative data type)](http://en.wikipedia.org/wiki/Associative_array) 입니다.
//  (다른 언어에선 _hashes_ 또는 _dicts_라고 부르기도 합니다.)
package main

import "fmt"

func main() {

	// 빈 맵을 생성하기 위해 내장 함수인 `make`를 사용합니다 : `make(map[key-type]val-type)`
	m := make(map[string]int)

	// `name[key] = val`로 key/value 쌍을 저장합니다.
	m["k1"] = 7
	m["k2"] = 13

	// 맵 출력. `Println`는 맵의 모든 key/value 쌍들을 출력합니다.
	fmt.Println("map:", m)

	// `name[key]`로 키에 해당하는 값을 가져옵니다.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// `len`을 맵에 사용하면 key/value 쌍의 갯수를 반환합니다.
	fmt.Println("len:", len(m))

	// `delete` 내장 함수는 맵의 key/value 쌍을 삭제합니다.
	delete(m, "k2")
	fmt.Println("map:", m)

	// 맵에서 값을 꺼내오면서 반환되는 선택적인 두번째 반환값은 해당 키가 맵에 존재하는지에 대한 여부를 나타냅니다.
	//  이는 키값이 존재하지 않는건지 또는 해당 값이 `0`이나 `""`과 같은 제로값 인지를 명확하게 구분하는데 사용할 수 있습니다.
	//  값 자체가 필요없는 경우엔, _공백 식별자(blank identifier)_ `_`를 사용해 무시할 수 있습니다.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 또한 한 줄로 맵 선언 및 초기화를 할 수도 있습니다.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
