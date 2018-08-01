./diskScheduler fcfs01.txt>fcfs01.out
diff fcfs01.base fcfs01.out
./diskScheduler sstf01.txt>sstf01.out
diff sstf01.base sstf01.out
./diskScheduler scan01.txt>scan01.out
diff scan01.base scan01.out
./diskScheduler c-scan01.txt>c-scan.out
diff c-scan.out c-scan01.base
./diskScheduler look01.txt>look01.out
diff look01.base look01.out
./diskScheduler c-look01.txt>c-look01.base
