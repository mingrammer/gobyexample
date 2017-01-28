// Go는 내장 및 커스텀 데이터 타입을 포함하여 JSON 인코딩 및 디코딩을 위한 내장 기능을 지원합니다.

package main

import "encoding/json"
import "fmt"
import "os"

// 아래에서 커스텀 타입에 대한 인코딩 및 디코딩을 시연하기 위해 다음의 두 구조체를 사용합니다.
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// 우선 기본 데이터 타입을 JSON 문자열로 인코딩 해봅시다.
	//  다음은 기본값들에 대한 몇 가지 예시입니다.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 그리고 다음은 예상대로 JSON 배열과 객체로 인코딩하는 슬라이스와 맵에 대한 예시입니다.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 패키지는 커스텀 데이터 타입을 자동으로 인코딩 할 수 있습니다.
	//  이는 노출된 필드만 인코딩 출력값에 포함시키며 JSON 키의 기본값으로는 필드명을 사용합니다.
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// JSON 키명을 커스터마이징 하기위해 구조체 필드 선언부에 태그를 사용할 수 있습니다.
	//  위의 `Response2`의 정의를 보면 태그의 예를 볼 수 있습니다.
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 이제 JSON 데이터를 Go의 값으로 디코딩 해봅시다.
	//  다음은 제네릭 데이터 구조의 예시입니다.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// JSON 패키지가 디코딩 데이터를 저장할 수 있는 변수를 선언해야합니다.
	//  `map[string]interface{}`는 임의의 데이터 타입의 문자열 맵을 갖습니다.
	var dat map[string]interface{}

	// 다음은 실제 디코딩이며, 관련 에러를 검사합니다.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 디코딩된 맵의 값을 사용하기 위해선, 적절한 타입으로 캐스팅 해야합니다.
	//  예를 들어 다음은 `num` 값을 `float64` 타입으로 캐스팅합니다.
	num := dat["num"].(float64)
	fmt.Println(num)

	// 중첩된 데이터에 접근하기 위해선 연속적으로 캐스팅을 해야합니다.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// JSON을 커스텀 데이터 타입으로 디코딩 할 수도 있습니다.
	//  이는 프로그램에 추가적인 타입 안전성을 추가할 수 있다는 장점을 가지며 데이터에 접근하고 디코딩할 때 타입 단언(assertion)을 할 필요성을 없애줍니다.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 위의 예시에서 데이터와 표준 출력상의 JSON 표현간의 매개체로 항상 바이트와 문자열을 사용했습니다.
	//  `os.Stdout`와 같은 `os.Writer`나 심지어는 HTTP 응답 바디에 직접 JSON 인코딩을 스트리밍 할 수도 있습니다.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
