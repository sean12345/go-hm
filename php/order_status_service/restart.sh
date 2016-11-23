ps -ef |egrep "php .*paimai_service_daemon.php" | awk '{print $2}' | xargs kill 
nohup /usr/bin/php `pwd`/paimai_service_daemon.php >> /tmp/paimai_status_service.log &

