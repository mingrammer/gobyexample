// _라인 필터 (line filter)_는 stdin으로 입력을 읽고, 처리한 후, 나온 결괏값을 stdout으로 출력하는 일반적인 유형의 프로그램입니다.
//  `grep`과 `sed`가 일반적인 라인 필터입니다.

// 다음은 모든 입력 텍스트의 대문자 버전을 작성하는 Go의 라인 필터 예제입니다.
//  여러분만의 Go 라인 필터를 작성하는데 이 패턴을 사용할 수 있습니다.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 버퍼링되지 않은 `os.Stdin`을 버퍼링된 스캐너로 래핑하는건 스캐너를 다음 토큰으로 진행시키는 편리한 `Scan` 메서드를 제공합니다.
	//  기본 스캐너에서 다음 토큰은 다음 라인입니다.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text`는 입력에서 현재 토큰을 반환합니다. 여기선 다음 라인입니다.
		ucl := strings.ToUpper(scanner.Text())

		// 대문자로된 문자열을 라인에 씁니다.
		fmt.Println(ucl)
	}

	// `Scan`중 에러를 체크합니다. 파일의 끝이 예상되며 `Scan`에 의한 에러로 보고되지 않습니다.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
