# 이 프로그램을 실행시키면 signal을 받을 때까지 블록될 것입니다.
#  `ctrl-C` (^C로 표시됩니다.) 를 눌러서 `SIGINT`를 보낼 수 있습니다.
#  SIGINT를 받은 프로그램은 `interrupt`를 출력하고 종료될 것입니다.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
