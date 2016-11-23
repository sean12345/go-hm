<?php
/**
 * 场内定时拍单状态
 */


global $stype;
print_r("\n");
echo "|$sid|:场内顺序拍单开始\n";
$dbm = get_mysql_handler();
$carlist = orderNotTimingOrder($special['scene_id']); //获得场内所有非定时拍单
$dbm->close();
if (isset($carlist) && is_array($carlist)) {
    echo "|$sid|:顺序拍单总数:" . count($carlist) . "\n";
    foreach ($carlist as $item) {
        if ($item != null) {
            $pid1 = pcntl_fork();
            if (!$pid1) {
                //启动一个新进程，开始拍单
                $stype = '顺序';
                require_once(API_ROOT_PATH . '/order.php');
                exit(0);
            } else {
                //等待子进程结束
                echo "|$sid|:顺序场拍单PID-" . $pid1 . "|order_no=" . $item['order_no'] . "\n";
                pcntl_wait($status_sequence);
                sleep(1);
            }
        }
    }
}


echo "|$sid|:场内顺序拍单结束\n";

