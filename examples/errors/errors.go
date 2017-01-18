// Go에선 명시적으로 별도의 반환값을 통해 에러를 전달하는게 일반적이고 관용적입니다.
//  익셉션을 사용하는 Java와 Ruby 그리고 때때로 결괏값과 에러값을 단일값에 오버로딩하는 C와는 대조적입니다.
//  Go의 이런 접근법을 통해 어떤 함수가 에러를 반환했는지를 보고 다른 에러 없는 작업에 사용되는 구조체를 사용하여 쉽게 처리할 수 있습니다.

package main

import "errors"
import "fmt"

// 관용적으로, 에러는 마지막 반환값이며 내장 인터페이스인 `error`를 타입으로 갖습니다.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New`는 전달된 에러 메시지로 기본적인 `error`값을 생성합니다.
		return -1, errors.New("can't work with 42")

	}

	// 에러 위치의 nil값은 에러가 없다는걸 나타냅니다.
	return arg + 3, nil
}

// `Error()` 메서드를 구현함으로써 커스텀 `error` 타입을 사용하는것도 가능합니다.
//  다음은 인자 에러를 명시적으로 나타내기 위해 커스텀 타입을 사용한 위의 변형된 예제입니다.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// 이 경우 우리는 `&argError` 구문을 사용하여 두 필드 `arg`와 `prob` 값을 가지고 새로운 구조체를 만듭니다.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// 아래의 두 루프는 에러를 반환하는 각 함수를 테스트합니다.
	// 참고로 `if` 라인에서 한 줄로 에러를 체크하는 방법은 Go의 공통적인 관용적 표현입니다.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 커스텀 에러에서 프로그래밍적으로 데이터를 사용하려면, 타입 단언(type assertion)을 통해 커스텀 에러의 인스턴스로 에러값을 받아와야 합니다.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
