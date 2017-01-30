// 직전의 예제에서 우리는 [외부 프로세스를 생성하는 법](/gobyexample/spawning-processes)
// 을 봤습니다. 우리는 실행되고 있는 Go 프로세스에서 외부 프로세스에
//  접근해야 할 때 프로세스를 생성합니다.
//  가끔, 우리는 그냥 프로세스 전체를 다른 프로세스로
//  대체하고싶을 때가 있습니다.
//  이를 위해서, 우리는 Go의 <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// 라는 고전적인 함수의 구현체를 이용할 것입니다.

package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// 우리의 예제에서는 `ls`를 exec할 것입니다.
	//  Go는 우리가 실행시킬 binary의 절대경로를 요구합니다.
	//  따라서, 우리는 `exec.LookPath`라는 함수를 이용하여
	//  절대경로를 찾아볼 것입니다. (아마도 `/bin/ls` 일 것입니다.)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec`는 커맨드만 담긴 커다란 문자열 하나가 아닌
	//  slice 형태로 된 인자를 요구합니다.
	//  `ls`에 흔히 넘기는 인자 몇개를 넘겨볼 것입니다.
	//  첫 번째 인자가 프로그램의 이름이어야 한다는 것을 명심하세요.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` 는 [환경변수](/gobyexample/environment-variables)도 필요로 합니다.
	//  여기서는 우리의 작업환경을 넘겨봅시다.
	env := os.Environ()

	// 여기서 `syscall.Exec`가 실제로 호출됩니다.
	//  만약 호출에 성공했다면 우리의 프로세스 실행은 여기서 끝나고
	//  `/bin/ls -a -l -h`로 대체될 것입니다.
	//  만약 에러가 있다면 여기서 반환된 값을 얻을 것입니다.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
