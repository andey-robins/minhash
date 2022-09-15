#!/bin/bash
go build -o main
mv res.txt res.txt.bak
for i in {2..1000}
do
    echo $i
    ./main -L 10007 -Q $i -seed 22 >> res.txt
done
rm main