// _스위치문_은 여러 분기에 걸친 조건문들을 표현합니다.

package main

import "fmt"
import "time"

func main() {

	// 여기에 기본적인 `switch`문이 있습니다.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 동일한 `case`문에서 여러개의 표현식을 구분하기 위해 콤마(,)를 사용할 수 있습니다.
	//  이 예시에서 우리는 또한 `default` 케이스도 사용했습니다. `default`는 선택사항입니다.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 표현식이 없는 `switch`는 if/else 로직을 표현하기 위한 또 다른 방법입니다.
	//  여기서 우리는 또한 상수가 아닌 `case`문을 사용하는 방법을 볼 수 있습니다.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 타입 `switch`는 값 대신 타입을 비교합니다.
	//  여러분은 인터페이스 값의 타입을 알아내기 위해 이를 사용할 수 있습니다.
	//  이 예시에서, 변수 `t`는 해당 절에 해당하는 타입을 가질 것입니다.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
