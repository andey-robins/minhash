#!/bin/bash
for i in {2..1000}
do
    echo $i
    go run main.go -L 10007 -Q $i -seed 11 >> res.txt
done