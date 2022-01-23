#!/bin/bash

# powered by mion
# please use bash command  run  this script like this: bash autorun.sh


PID_MEM_MAX=" 95 "


SYS_LOAD_MAX=" 1 "

NAME_LIST=" php-fpm mysql v2ray nginx "

for NAME in $NAME_LIST
do

        PID_MEM_SUM=0


        PID_NUM_SUM=` ps aux | grep $NAME | wc -l `


        PID_MEM_LIST=` ps aux | grep $NAME | awk '{ print $4 }' `


        for PID_MEM in $PID_MEM_LIST
        do
                PID_MEM_SUM=` echo $PID_MEM_SUM + $PID_MEM | bc `
        done


        SYS_LOAD=` uptime | awk '{ print $(NF-2) }' | sed 's/,//' `


        MEM_VULE=`awk 'BEGIN{ print("$PID_MEM_SUM" >= "$PID_MEM_MAX"?"1":"0")}'`
        LOAD_VULE=`awk 'BEGIN{ print("$SYS_LOAD" >= "$SYS_LOAD_MAX"?"1":"0")}'`


        if [[ $MEM_VULE = 1 ]] || [[ $LOAD_VULE = 1 ]]  ;then

                echo $(date +"%y-%m-%d %H:%M:%S") "killall $NAME" "(MEM:$PID_MEM_SUM,LOAD:$SYS_LOAD)" >> /root/autoreboot.log

                # /etc/init.d/$NAME stop
                service $NAME stop
                sleep 3

                pkill $NAME


                # /etc/init.d/$NAME start
                service $NAME start

                echo $(date +"%y-%m-%d %H:%M:%S") "start $NAME" "(MEM:$PID_MEM_SUM,LOAD:$SYS_LOAD)" >> /root/autoreboot.log
        else
                echo "$NAME very health!(MEM:$PID_MEM_SUM,LOAD:$SYS_LOAD)" > /dev/null
        fi
        done
        exit 0
          
