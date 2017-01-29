// URLs는 [리소스를 찾을 수 있는 일관된 방법(uniform way to locate resources)](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)을 제공합니다.
//  여기에 Go에서 URLs를 파싱할 수 있는 방법이 있습니다.

package main

import "fmt"
import "net"
import "net/url"

func main() {

	// 다음의 스키마, 인증 정보, 호스트, 포트, 경로, 쿼리 파라미터, 그리고 쿼리 프레그먼트를 포함하고 있는 예시 URL을 파싱해보겠습니다.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL을 파싱하고 에러가 없는지를 확인합니다.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 스키마 접근은 간단합니다.
	fmt.Println(u.Scheme)

	// `User`는 모든 인증 정보를 포함하고 있습니다. 각각의 값에 대해선 `Username`과 `Password`를 호출하면 됩니다.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host`는 호스트명과 포트를 포함하고 있습니다 (존재하는 경우). 이 값들을 추출하기 위해선 `SplitHostPort`를 사용합니다.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// `경로(path)`와 `#` 뒤의 프레그먼트를 추출하고 있습니다.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	// `k=v` 포맷의 문자열에서 쿼리 파라미터를 얻기 위해 `RawQuery`를 사용합니다.
	//  쿼리 파라미터를 맵으로 파싱할 수도 있습니다.
	//  파싱된 쿼리 파라미터 맵은 문자열부터 문자열의 슬라이스까지이므로 첫번째 값만을 가져오려면 `[0]`를 사용합니다.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
