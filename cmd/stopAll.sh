ps -ef|grep gateway|grep -v grep|awk '{print $2}' | xargs kill -9
ps -ef|grep game|grep -v grep|awk '{print $2}' | xargs kill -9
ps -ef|grep world|grep -v grep|awk '{print $2}' | xargs kill -9