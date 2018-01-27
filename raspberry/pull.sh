#!/bin/bash
sleep 5
while true
do
if ping -q -c 1 -W 1 google.com >/dev/null; then
  echo "The network is up"
 git pull
sleep 60

else
  echo "The network is down"
fi
done

