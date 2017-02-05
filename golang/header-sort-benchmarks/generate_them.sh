#!/bin/bash
benchoutput=benchmarks.txt
echo "Benchmark start date: $(date)" > $benchoutput

for ((i=0; i<10; i++))
do
  go test -bench=. >> $benchoutput
done

echo "Benchmark end date: $(date)" >> $benchoutput
