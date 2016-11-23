<?php
/**
 * 拍单状态
 */

// 单个车源拍卖开始
global $stype;
print_r("|$sid|:拍卖车源开始[" . $cur_time_id . "]：" . $item['order_id'] . "\n");
$dbm = get_mysql_handler();
$db = get_mysqli_db();
orderStart($item['order_id']);
$dbm->close();
echo "|$sid|：{$stype}拍单" . $item['order_no'] . "竞拍中。。。。。。。。\n";

$special_car_sleep = $item["est_elapsed_time"];
// 单个车源拍卖2分钟结束
sleep($special_car_sleep);

/********处理拍卖延时**********/
$overtime = true;
$times = $special_car_sleep; //预计时长，（如：60）
while ($overtime) {
    $dbm = get_mysql_handler();
    $_times = getOrderTimes($item['order_id']); // 真实要结束的时长
    $dbm->close();
    if ($_times > $times) {
        sleep($_times - $times);
        $times = $_times;
    } else {
        $overtime = false;
    }
}
/********处理拍卖延时**********/


// 单个车拍卖结束
$db = get_mysqli_db();
$redis = get_redis();
orderEnd($item['order_id']);
breachRedo($item['order_id']);
print_r("|$sid|:拍卖车源结束[" . date("Y-m-d H:i:s") . "]：" . $item['order_id'] . "\n");
exit(0);
