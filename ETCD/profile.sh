#!/usr/bin/env bash
cd ../go-ycsb || exit
./bin/go-ycsb load etcd -p etcd.endpoints="10.10.1.1:2379,10.10.1.2:2379,10.10.1.2.2379,10.10.1.3:2379,10.10.1.4:2379,10.10.1.5:2379" -P workloads/workloada_gryff
./bin/go-ycsb run etcd -p etcd.endpoints="10.10.1.1:2379,10.10.1.2:2379,10.10.1.2.2379,10.10.1.3:2379,10.10.1.4:2379,10.10.1.5:2379" -P workloads/workloada_gryff