#!/bin/bash
while true
do
if ping -q -c 1 -W 1 google.com >/dev/null; then
  echo "The network is up"
 gut pull
else
  echo "The network is down"
fi
sleep 10
done

