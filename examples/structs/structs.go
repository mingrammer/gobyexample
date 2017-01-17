// Go의 _structs_는 필드들로 이루어진 타입을 갖는 컬렉션입니다.
//  레코드를 구성하기 위해 데이터들을 그룹핑 하는데 유용합니다.

package main

import "fmt"

// 이 `person` 구조체 타입은 `name`과 `age` 필드를 갖습니다.
type person struct {
	name string
	age  int
}

func main() {

	// 새로운 구조체를 생성합니다.
	fmt.Println(person{"Bob", 20})

	// 필드명을 사용해 구조체를 초기화 할 수 있습니다.
	fmt.Println(person{name: "Alice", age: 30})

	// 생략된 필드는 제로값을 갖게됩니다.
	fmt.Println(person{name: "Fred"})

	// 앞에 `&`를 붙이면 구조체 포인터를 생성할 수 있습니다.
	fmt.Println(&person{name: "Ann", age: 40})

	// 닷(.)을 사용해 구조체 필드에 접근합니다.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 구조체 포인터에서도 닷(.)을 사용할 수 있습니다.
	//  이 때 포인터는 자동으로 역참조됩니다.
	sp := &s
	fmt.Println(sp.age)

	// 구조체는 수정이 가능(mutable)합니다.
	sp.age = 51
	fmt.Println(sp.age)
}
