// _range_는 다양한 데이터 구조의 요소들을 순회합니다.
//  우리가 이미 배운 몇 가지 데이터 구조와 함께 `range`를 사용하는 방법을 살펴봅시다.

package main

import "fmt"

func main() {

	// 다음 예시는 슬라이스에 있는 숫자들을 더하기 위해 `range`를 사용합니다.
	//  배열에서도 똑같이 동작합니다.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range`를 배열과 슬라이스에서 사용하면 각 원소의 인덱스와 값을 반환합니다.
	//  위 예시에선 인덱스가 필요 없었기 때문에 공백 식별자 `_`로 인덱스를 무시했습니다.
	//  다음 예시처럼 인덱스가 필요할 때도 있습니다.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 맵에서 `range`는 key/value 쌍들을 순회합니다.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range`로 map의 키값들만 순회할 수도 있습니다.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 문자열에서 `range`는 유니코드 코드 포인트들을 순회합니다.
	//  첫번째 값은 `rune`의 시작 바이트 위치이며 두번째 값은 `rune`값 자체입니다.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
