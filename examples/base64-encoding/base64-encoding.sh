# 이 문자열은 표준과 URL base64 인코더를 사용하여 약간 다른 값으로 인코딩하지만, (후미가 `+` vs `-`) 둘 다 우리가 원하는 원래 문자열로 디코딩합니다.  
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
