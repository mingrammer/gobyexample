# 프로그램을 실행하면 우리가 원하던대로 문자열 길이를 기준으로 정렬된 리스트가 보여집니다.
$ go run sorting-by-functions.go 
[kiwi peach banana]

# 커스텀 타입을 생성하고, 해당 타입에 세 개의 `Interface` 메서드를 구현한 다음 커스텀 타입의 컬렉션에 sort.Sort를 호출하는 패턴을 따르면 어떤 함수로도 Go 슬라이스를 정렬할 수 있습니다.