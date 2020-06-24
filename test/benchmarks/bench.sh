#!/bin/bash"
TARGET="countgo"



CMD_RESOURCES=`printf 'https://golang.org/doc/effective_go.html\nhttps://ru.wikipedia.org/wiki/Go\n%.0s' {1..200}`
rm -f test/benchmarks/plot.dat 

for worker_num in 1 2 4 8 16 32 64
do
	echo "Running echo -e <big_resourse_number> | ./$TARGET -k=$worker_num"
	TIME_SEC=`date "+%s"`
	echo -e "$CMD_RESOURCES" | ./countgo -k=$worker_num > /dev/null
	EXCODE="$?"
	printf "\t$((`date "+%s"`-$TIME_SEC)) sec\n"
	echo "$worker_num $((`date "+%s"`-$TIME_SEC))" >> ./test/benchmarks/plot.dat
	sleep 1 # Let previous process free all resources
	if [[ $EXCODE != 0 ]]
	then
		echo "./build/$1 failed with error $?, aborting"
		exit
	fi
done


gnuplot <<< "set term png size 1920,1080; \
             set output 'test/benchmarks/bench.png'; \
			 set ylabel 'Time in seconds' font 'Times Italic, 20'; \
			 set xlabel 'worker number' font 'Times Italic, 20'; \
             set tmargin 2; \
             plot 'test/benchmarks/plot.dat' u 1:2 w l;"