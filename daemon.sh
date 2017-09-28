#!/usr/bin/env bash
# daemon.sh to /etc
# chmod 777 daemon.sh
# do ./etc/daemon.sh
# vi /etc/rc.loal
# add ./etc/daemon.sh &
function  CheckProcess
{
  if [ "$1" = "" ];
  then
    return 1
  fi
  PROCESS_NUM=`ps -ef | grep "$1" | grep -v "grep" | wc -l`
  if [ $PROCESS_NUM -eq 1 ];
  then
    return 0
  else
    return 1
  fi
}

while [ 1 ] ; do
 CheckProcess "appname"
  CheckQQ_RET=$?
  if [ $CheckQQ_RET -eq 1 ];
  then
    /usr/local/app/appname &
 fi

# other proc
 CheckProcess "appname1"
 CheckQQ_RET=$?
 if [ $CheckQQ_RET -eq 1 ];
then
   /usr/local/app/appname1 &
 fi

 sleep 2
done
