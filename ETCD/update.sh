#!/bin/bash
cd ../Raft || exit
git stash && git stash clear && git pull
cd ../RabiaGo || exit
git stash && git stash clear && git pull
cd ../PineappleGo || exit
git stash && git stash clear && git pull
cd ../ETCD || exit
git stash && git stash clear && git pull

HOST=$(hostname | awk -F "." '{print $1}')
echo "Hostname: $HOST"

if [ $HOST = "node-1" ]; then
    IP="10.10.1.1"
elif [ $HOST = "node-2" ]; then
    IP="10.10.1.2"
elif [ $HOST = "node-3" ]; then
    IP="10.10.1.3"
elif [ $HOST = "node-4" ]; then
    IP="10.10.1.4"
elif [ $HOST = "node-5" ]; then
    IP="10.10.1.5"
fi
echo "Local IP: $IP"

if [ "$1" = "pineapple" ]; then
  export PINEAPPLE="true"
else
  export PINEAPPLE="false"
fi

if [ "$1" = "memory" ]; then
  export PINEAPPLE_MEMORY="true"
else
  export PINEAPPLE_MEMORY="false"
fi

if [ "$1" = "rs-rabia" ]; then
  export RS_RABIA="true"
else
  export RS_RABIA="false"
fi

if [ "$1" = "" ]; then
  export SETUP="--initial-cluster node-1=http://10.10.1.1:12380,node-2=http://10.10.1.2:12380,node-3=http://10.10.1.3:12380"
fi

sudo rm -rf "$HOST.etcd"
make build
sudo ./bin/etcd --log-level panic \
--name "$HOST" \
--initial-cluster-token etcd-cluster-1 \
--listen-client-urls http://"$IP":2379,http://127.0.0.1:2379 \
--advertise-client-urls http://"$IP":2379 \
--initial-advertise-peer-urls http://"$IP":12380 \
--listen-peer-urls http://"$IP":12380 \
$SETUP \
--initial-cluster-state new



