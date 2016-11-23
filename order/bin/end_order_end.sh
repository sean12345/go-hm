ps -ef |egrep "*order -s end" | awk '{print $2}' | xargs kill
