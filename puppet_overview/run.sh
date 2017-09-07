#!/bin/bash

/usr/bin/supervisord -c /etc/supervisord.conf

cd /mnt/build
find puppet_overview.slide graphs | justrun -c "make graphs && supervisorctl restart present" -w -delay 2s -stdin &

while true
do
    sleep 2
done
