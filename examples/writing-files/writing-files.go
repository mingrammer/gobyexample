// Go에서 파일 쓰기는 이전에 본 읽기와 유사한 패턴을 갖습니다.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 시작은, 문자열 (또는 바이트)을 파일로 덤프하는 방법입니다.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 보다 세분화된 쓰기를 위해, 파일을 엽니다.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// 파일을 연 직후 `Close`를 지연시키는건 Go에서의 관행입니다.
	defer f.Close()

	// 예상한대로 바이트 슬라이스를 `Write` 할 수 있습니다.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString`도 사용 가능합니다.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// `Sync`를 실행하여 안정적인 스토리지에 쓰기를 플러시합니다.
	f.Sync()

	// `bufio`는 이전에 봤던 버퍼링된 리더에 추가로 버퍼링된 라이터(writers)를 제공합니다.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// `Flush`를 사용하여 모든 버퍼링된 작업이 라이터에 적용되었는지 확인합니다.
	w.Flush()

}
