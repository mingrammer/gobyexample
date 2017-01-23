# 프로그램은 5개의 잡이 여러개의 워커에 의해 실행되고 있음을 보여줍니다.
#  총 작업 시간은 5초지만 3개의 워커가 동시에 실행되고 있기 때문에, 전체 작업은 약 2초만에 끝납니다.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
