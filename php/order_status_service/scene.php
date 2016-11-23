<?php
/**
 *  场拍状态
 */


//专场拍卖开始
print_r("\n\n\n");
print_r("|$sid|:拍卖专场开始：[" . $cur_time_id . "]" . $special['name'] . "\n");

$dbm = get_mysql_handler();
$scene_mg_pids = array();
sceneStart($special['scene_id']);

/***********专场车源拍卖************/
$carlist = getOrderList($sid);
echo "|$sid|共找到场内拍单数：" . count($carlist) . "\n";
if ((count($carlist) > 0)) {
    orderWaitBidding($carlist);//把场所有拍单改成“待竞标”状态
    //P1.启动定时拍单处理进程
    $pid_timing = pcntl_fork();

    if (!$pid_timing) {
        //启动一个新进程
        require_once(API_ROOT_PATH . '/timing_casual_add.php');
        exit(0);
    } else {
        $scene_mg_pids[] = $pid_timing;
        echo "|$sid|:定时场PID-" . $pid_timing . "\n";
    }

    //P2.继续处理非定时的顺序拍单状态
    $pid_seq = pcntl_fork();
    if (!$pid_seq) {
        //启动一个新进程
        require_once(API_ROOT_PATH . '/sequence.php');
        exit(0);
    } else {
        $scene_mg_pids[] = $pid_seq;
        echo "|$sid|:顺序场PID-" . $pid_seq . "\n";
    }

    $dbm->close();
    /***********专场车源拍卖************/
    wait_pids($scene_mg_pids);
}

//专场拍卖结束
sceneEnd($special['scene_id']);
$dbm->close();
print_r("|$sid|:拍卖专场结束：[" . $cur_time_id . "]" . $special['name'] . "\n");
print_r("\n\n\n");









	