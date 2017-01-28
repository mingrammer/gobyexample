// Go는 패턴 기반의 레이아웃을 통한 시간 포맷팅과 파싱을 지원합니다.

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// 다음은 RFC3339에 해당하는 레이아웃 상수를 사용하여 시간을 포맷팅하는 기본적인 예시입니다.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 시간 파싱은 `Format`과 동일한 레이아웃 값을 사용합니다.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format`과 `Parse`는 예시 기반 레이아웃을 사용합니다. 보통 이 레이아웃들은 `time`에서의 상수를 사용하지만, 커스텀 레이아웃을 사용할 수도 있습니다.
	//  레이아웃은 참조 시간 `Mon Jan 2 15:04:05 MST 2006`을 사용하여 주어진 시간/문자열을 포맷팅/파싱할 패턴을 표시해야합니다.
	//  예시 시간은 다음에 표시된것과 정확히 일치해야합니다: 연도는 2006, 시간은 15, 요일은 월요일 등등.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 순전히 숫자로 표현하기 위해 시간값에서 추출한 컴포넌트와 표준 문자열 포맷팅을 사용할 수도 있습니다.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse`는 잘못된 형식의 입력값이 들어오면 에러를 반환하여 파싱 문제를 설명합니다.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
