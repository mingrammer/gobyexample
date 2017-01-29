# 프로그램을 실행하면 해시를 계산하고 이를 사람이 읽을 수 있는 16진수 포맷으로 출력합니다. 
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 위에서 봤던 패턴과 유사한 패턴을 사용하여 다른 해시를 계산할 수 있습니다.
#  예를 들면, MD5 해시를 계산하기 위해선 `crypto/md5`를 임포트하고 `md5.New()`를 사용합니다.

# 참고로 암호학적으로 안전한 해시가 필요하다면 [해시 강도(hash strength)](http://en.wikipedia.org/wiki/Cryptographic_hash_function)를 주의깊게 조사해야합니다.