#!/bin/bash

# add plan to crontab 
# powered by mion

# please use bash command  run  this script like this bash crontab.sh 
crontab -l > crontab.bak
echo "* * * * * /root/shell/freemem.sh" >> crontab.bak
crontab crontab.bak

/etc/init.d/cron restart
