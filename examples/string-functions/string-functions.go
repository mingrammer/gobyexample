// 표준 라이브러리의 `strings` 패키지는 많은 유용한 문자열 관련 함수들을 제공합니다.
//  이 패키지에 대한 감을 얻을 수 있는 몇 가지 예시가 있습니다.

package main

import s "strings"
import "fmt"

// `fmt.Println`를 짧은 이름으로 aliasing 합니다. 이는 아래에서 많이 사용합니다.
var p = fmt.Println

func main() {

	// `strings`에서 사용할 수 있는 몇 가지 함수의 예시입니다.
	//  이들은 문자열 객체 자체의 메서드가 아니라 패키지의 함수이기 때문에 문자열을 함수의 첫번째 인자에 전달해줘야 합니다.
	//  [`strings`](http://golang.org/pkg/strings/) 패키지 문서에서 더 많은 함수들을 찾아볼 수 있습니다.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// `strings`의 일부는 아니지만 언급할만한게 있는데, 바이트에서의 문자열 길이와 인덱스를 통한 바이트를 구하는 메커니즘입니다.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// 참고로 위의 `len`과 인덱싱은 바이트 수준에서 동작합니다.
//  Go는 UTF-8로 인코딩된 문자열을 사용하므로 종종 유용합니다.
//  잠재적으로 다중 바이트 바이트 문자로 작업하는 경우 인코딩 인식 작업을 하고자 할겁니다.
//  좀 더 자세한 내용은 [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)를 보십시오.
