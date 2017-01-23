// _지연(Defer)_은 함수 호출이 나중에 프로그램의 실행에서 수행되도록 보장 하기위해 사용됩니다.
//  일반적으로는 코드 정리를 위한 목적으로 사용합니다.
//  `defer`는 보통 타 언어에서 `ensure`와 `finally`가 사용되는 곳에서 사용됩니다.

package main

import "fmt"
import "os"

// 파일을 만들고, 값을 쓴 다음 다 끝나면 파일을 종료시키고 싶다고 해보자.
//  다음은 `defer`를 사용해 이를 구현한 예시입니다.
func main() {

	// `createFile`로 파일 객체를 얻은 직후, `closeFile`로 파일을 종료시키는 작업을 지연합니다.
	//  이는 `writeFile`이 끝나고, 함수(`main`)가 끝나는 지점에서 실행됩니다.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
