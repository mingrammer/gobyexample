// Go는 구조체 타입에 정의할 수 있는 _methods_를 지원합니다.

package main

import "fmt"

type rect struct {
	width, height int
}

// `area` 메서드는 `*rect`를 _리시버 타입(receiver type)_으로 갖습니다.
func (r *rect) area() int {
	return r.width * r.height
}

// 메서드는 포인터 또는 값을 리시버 타입으로하여 정의될 수 있습니다.
//  여기에 값을 리시버로 하는 예시가 있습니다.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 구조체에 정의된 두 개의 메서드를 호출했습니다.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go는 메서드 호출에 대해 값과 포인터간의 변환을 자동으로 처리합니다.
	//  메서드 호출시 값 복사를 피하기 위해 포인터 리시버 타입을 사용할 수도 있고 전달되는 구조체를 변경할 수 있도록 할 수도 있습니다.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
