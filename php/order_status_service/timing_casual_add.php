<?php
/**
 *   场内定时拍单状态,场拍结束前随便添加定时订单,便需要增加一个定时单时间长一点
 */


print_r("\n");
echo "|$sid|:场内定时拍单开始\n";
global $stype;
$dbm = get_mysql_handler();
$rows = orderTimingOrder($special['scene_id']);

echo "|$sid|:定时拍单总数:" . count($rows) . "\n";
$order_list = array();
foreach ($rows as $key => $row) {
    $order_list[$row['bidding_start_time']][] = 1;
}
$order_list_num = count($order_list);
$timing_order_pids = array();
$dbm->close();

while ($order_list_num > 0) {
    $cur_time_id = date("Y-m-d H:i:00");
    $dbm = get_mysql_handler();
    $orders = orderTimingOrderByTimer($sid, $cur_time_id);
    $dbm->close();

    if ($orders) {
        echo "|$sid|:此时定时拍单数:" . $cur_time_id . "|" . count($orders) . "\n";
        foreach ($orders as $key => $order) {
            $item = $order;
            $timing_order_pid = pcntl_fork();
            if (!$timing_order_pid) {
                //启动一个新进程，开始定时拍单
                $stype = '定时';
                require_once(API_ROOT_PATH . '/order.php');
                exit(0);
            } else {
                echo "|$sid|:定时场拍单PID-" . $timing_order_pid . "|order_no=" . $item['order_no'] . "\n";
                $timing_order_pids[] = $timing_order_pid;
            }
        }
        unset($order_list[$cur_time_id]);
        $order_list_num--;
    }
    if ($order_list_num > 0) {
        $sleeptimes = getNextSleepSecond();
        sleep($sleeptimes);
    }
}


wait_pids($timing_order_pids);
echo "|$sid|:场内定时拍单结束\n";
