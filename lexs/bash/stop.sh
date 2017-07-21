#!/bin/bash

# ----------------------
# shutdow all server
# ----------------------

ps aux | grep -E '(conf.x|account.x|image.x|post.x|feeds.x|recsys.x|search.x)' | grep -v 'grep' > _PIDS

PNAME=(`awk -F ' ' '{print $2   $11}' _PIDS`)

PIDS=(`awk -F ' ' '{print $2}' _PIDS`)

echo "============== server List  ========================="
echo "PID      SERVER"
for v in ${PNAME[*]}
do
	echo $v
done
echo "===================================================="

for pid in ${PIDS[*]}
do
	kill -9 $pid > /dev/null 2>&1
done

echo "All Server Killed."
echo

rm -f _PIDS


