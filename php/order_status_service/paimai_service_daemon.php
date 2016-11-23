<?php
/**
 * 场状态进程调度程序
 * ALTER TABLE `aums_sam`.`au_order_scene`     ADD COLUMN `endtime` DATETIME NULL COMMENT '场拍卖实际结束时间' AFTER `createtime`;
 */
//error_reporting ( 0 );
date_default_timezone_set("Asia/Shanghai");
set_time_limit(0);
if (isset ($_SERVER ['REQUEST_URI'])) {
    die ('error');
}
if (!defined('API_ROOT_PATH')) {
    define('API_ROOT_PATH', dirname(__FILE__));
}

global $dbh, $sid, $dbm, $stype;
require_once(API_ROOT_PATH . '/func.php');

//$special_is_null_sleep = 60;
$special_is_null_sleep = 0;
$special_car_sleep = 120;
$special_car_sleep1 = 5;
$cur_time_id = null;

$i = 0;
$sce_pids = array();
while (true) {
    $dbm = get_mysql_handler();
    print_r("\n");
    @ob_flush();

    $sleeptimes = getNextSleepSecond();
    print_r("sleep:" . $sleeptimes . "\n");
    sleep($sleeptimes);

    $cur_time_id = date("Y-m-d H:i:00");
    echo '查找' . $cur_time_id . '场' . "\n";

    $scenes = getScenes($cur_time_id);
    $dbm->close();

    foreach ($scenes as $sce_key => $scene) {
        $special = $scene;
        $sid = $scene['scene_id'];
        if ($special != null) {
            $i++;
            $pid = pcntl_fork();
            if (!$pid) {
                //启动一个新进程，开始场拍卖
                require_once(API_ROOT_PATH . '/scene.php');
                exit(0);
            } else {
                echo "|$sid|:场PID-" . $pid . "\n";
            }
        } else {
            print_r("$sid:" . date("Y-m-d H:i:s") . "special is null\n");
            sleep($special_is_null_sleep);
        }
    }
}



