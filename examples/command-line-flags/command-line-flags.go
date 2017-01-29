// [_커맨드라인 플래그 (Command-line flags)_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)는 커맨드라인 프로그램에 옵션을 지정하는 일반적인 방법입니다.
//  예를 들어, `wc -l`에서 `-l`은 커맨드라인 플래그입니다.

package main

// Go는 기본적인 커맨드라인 플래그 파싱을 지원하는 `flag` 패키지를 제공합니다.
//  이 패키지를 사용하여 예제 커맨드라인 프로그램을 구현해보겠습니다.
import "flag"
import "fmt"

func main() {

	// 기본적인 플래그 선언은 문자열, 정수, 그리고 불리언 옵션에서 사용 가능합니다. 다음은 `"foo"`를 기본값으로 하고 간단한 설명을 가진 `word`라는 문자열 플래그의 선언입니다.
	//  `flag.String` 함수는 문자열 포인터를 반환합니다 (문자열 값이 아님)
	//  아래에서 이 포인터를 어떻게 사용하는지 살펴보겠습니다.
	wordPtr := flag.String("word", "foo", "a string")

	// 다음은 `word` 플래그 선언과 유사한 방법을 사용해 선언한 `numb`와 `fork` 플래그입니다.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 프로그램의 기존에 다른 곳에서 선언된 변수를 사용해 옵션을 선언하는 것도 가능합니다.
	//  참고로 플래그 선언 함수에 변수의 포인터를 전달해야합니다.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 모든 플래그가 선언되면, 커맨드라인 파싱을 실행하기 위해 `flag.Parse()`을 호출합니다.
	flag.Parse()

	// 다음은 파싱된 옵션들과 후미에 위치하는 인자들을 덤프합니다.
	//  참고로 실제 옵션값을 얻기 위해선 `*wordPtr`와 같이 포인터를 역참조 해야합니다.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
