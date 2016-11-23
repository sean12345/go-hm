ps -ef |egrep "*order -s start" | awk '{print $2}' | xargs kill
nohup ./order -s start >> /tmp/order_start.log &
