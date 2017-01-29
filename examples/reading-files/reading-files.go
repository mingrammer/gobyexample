// 파일 읽기와 쓰기는 많은 Go 프로그램에서 필요로하는 기본적인 작업입니다.
//  우선 파일 읽기의 몇 가지 예시를 살펴보겠습니다.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 파일을 읽으려면 대부분의 에러 호출을 확인해야합니다. 이 헬퍼는 아래에서의 에러 체크를 간소화합니다.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 가장 기본적인 파일 읽기 작업은 아마 파일의 전체 내용을 메모리로 올리는 작업일겁니다.
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 종종 파일 내용에 대해 더 많은 제어를 하고 싶을때가 있습니다.
	//  이를 위해선 파일을 `Open`하여 `os.File` 값을 얻습니다.
	f, err := os.Open("/tmp/dat")
	check(err)

	// 파일의 첫 부분에서 일부 바이트를 읽습니다.
	//  5 바이트까지 읽기를 허용하고 실제로 얼마나 많이 읽었는지를 기록합니다.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 또한 파일의 특정 위치를 `Seek`하고 그 지점으로부터 `Read`를 할 수도 있습니다.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// `io` 패키지는 파일을 읽을때 유용할 만한 몇 가지 함수를 제공합니다.
	//  예를 들어, 위와 같은 읽기는 `ReadAtLeast`로 좀 더 견고하게 구현할 수 있습니다.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 내장 기능중 되감기 기능은 없지만 `Seek(0, 0)`으로 대신 할 수 있습니다.
	_, err = f.Seek(0, 0)
	check(err)

	// `bufio` 패키지는 많은 작은 읽기의 효율성과 패키지가 제공하는 추가적인 읽기 메서드 덕분에 유용할 수 있는 버퍼링된 리더를 구현하고 있습니다.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 작업이 끝나면 파일을 닫습니다. (보통 `Open` 직후 `defer`로 스케쥴링합니다.)
	f.Close()

}
