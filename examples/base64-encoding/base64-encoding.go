// Go는 [base64 인코딩/디코딩(base64 encoding/decoding)](http://en.wikipedia.org/wiki/Base64)을 내장 기능으로 지원합니다.

package main

// 다음 구문은 `encoding/base64` 패키지를 기본값인 `base64` 대신 `b64`라는 이름으로 임포트 합니다.
//  이는 공간을 조금 절약합니다.
import b64 "encoding/base64"
import "fmt"

func main() {

	// 다음은 우리가 인코딩/디코딩할 `문자열(string)` 입니다.
	data := "abc123!?$*&()'-=@~"

	// Go는 표준과 URL 호환 base64 모두 지원합니다.
	//  다음은 표준 인코더로 인코딩하는 방법입니다.
	//  인코더는 `[]byte`를 받으므로 `string`을 이 타입으로 캐스팅 해야합니다.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 디코딩은 에러를 반환할 수도 있는데, 입력값이 올바론 형태인지 모를 경우 이를 통해 확인할 수 있습니다.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 다음은 URL 호환 base64 포맷으로 인코딩/디코딩하는 예입니다.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
