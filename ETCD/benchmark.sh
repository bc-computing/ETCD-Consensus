#!/bin/bash
HOST_1="http://192.168.1.1:2379"
HOST_2="http://192.168.1.2:2379"
HOST_3="http://192.168.1.3:2379"
HOST_4="http://192.168.1.4:2379"
HOST_5="http://192.168.1.5:2379"
go run ./tools/benchmark \
--endpoints=${HOST_1},${HOST_2},${HOST_3},${HOST_4},${HOST_5} --conns=100 --clients=1000 \
put --key-size=8 --sequential-keys --total=100000 --val-size=10