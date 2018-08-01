./diskScheduler fcfs20.txt>fcfs20.out
diff fcfs20.base fcfs20.out
./diskScheduler sstf20.txt>sstf20.out
diff sstf20.base sstf20.out
./diskScheduler scan20.txt>scan20.out
diff scan20.base scan20.out
./diskScheduler c-scan20.txt>c-scan20.out
diff c-scan20.out c-scan20.base
./diskScheduler look20.txt>look20.out
diff look20.base look20.out
./diskScheduler c-look20.txt>c-look20.out
diff c-look20.out c-look20.base