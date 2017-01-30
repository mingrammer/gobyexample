// 간혹 우리의 Go 프로그램이 다른 Go가 아닌 프로세스를 띄워야 할
//  때가 있습니다. 예를 들어, 이 페이지의 syntax highlighting은
//  [`pygmentize`](http://pygments.org/)를 Go 프로그램에서 띄우는 방식으로
//  [구현](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// 되었습니다. Go에서 다른 프로세스를 띄우는 몇 가지 예제를 살펴봅시다.

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

	// 인자와 입력을 받지 않고 stdout에 아무거나 출력하는
	//  간단한 커맨드에서 시작해봅시다. `exec.Command`는
	//  외부 프로세스를 표현하기 위한 객체를 만들어주는 유용한
	//  도구입니다.
	dateCmd := exec.Command("date")

	// `.Output`은 커맨드 실행과 종료 대기, 그리고 output을
	//  가져오는데 사용되는 유용한 도구입니다. 에러가 없다면
	//  `dateOut`에는 날짜 정보를 담은 바이트를 저장할 겁니다.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 다음으로는 데이터를 `stdin`으로 파이프하여 `stdout`에서
	//  output을 가져오는 조금 더 복잡한 케이스를 살펴보도록 하겠습니다.
	grepCmd := exec.Command("grep", "hello")

	// 여기서는 명시적으로 input과 output을 파이프하여
	//  프로세스를 시작하고, input에 데이터를 쓴 뒤
	//  나오는 output을 가져오고, 프로세스 종료를 기다립니다.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 위의 예제에서는 에러 체크를 생략했지만, 모든 경우에 대해
	//  `if err != nil` 패턴을 사용할 수 있습니다.
	//  또한 우리가 위에서 `StdoutPipe`의 결과만 수집했지만,
	//  정확히 같은 방법으로 `StderrPipe`의 결과도 수집할 수 있습니다.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 프로세스를 생성할 때 명시적으로 인자별로 구분된
	//  인자 배열을 제공하거나, 전체 커맨드가 담긴
	//  문자열 하나를 제공한다는 점을 기억하세요.
	//  만약 전체 커맨드를 담은 문자열 하나로 프로세스를 생성하고 싶다면
	//  `bash`의 `-c` 옵션을 사용하면 됩니다.
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
