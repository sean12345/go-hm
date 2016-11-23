<?php
/**
 * 拍卖服务逻辑：
 * 1、计算离最近00、10、20、30、40、50结尾的分钟时间还有多少秒
 * 2、用这个秒来设置sleep时间
 * 3、取得有当前时间是否有拍卖场次，如果没有则sleep 60 秒再跳到第1步，如果有进入第4步
 * 4、把当前场次<paimai_special>的starttime=now(),state=1
 * 5、取得当前场次的要拍卖的车源列表，seq desc排序
 * 6、如果没有记录跳到8步，如果有记录进入第7步
 * 7、循环这个拍卖列表，从第一条开始，update paimai_special_car set starttime=now(), state=1然后 sleep(120)再执行拍卖结算函数
 * 8、循环结束执行场次结束函数
 * 9、跳入第一步
 *
 */

require_once(API_ROOT_PATH . '/../inc/Config.php');
require_once(API_ROOT_PATH . '/../inc/dblink.php');
require_once(API_ROOT_PATH . '/../inc/Mysql.php');
require_once(API_ROOT_PATH . '/../inc/MysqliDb.php');
require_once(API_ROOT_PATH . '/php_redis.php');


//定义服务费规则
define('PERCENTAGE', '0.01');
define('MIN_FEE', '300');
define('MAX_FEE', '2000');

/**
 * 计算交易服务费
 * @param int $price 初始价格
 * @return int $p 交易服务费
 */
function get_commision($price)
{
    $p = $price * PERCENTAGE;
    if ($p <= MIN_FEE) {
        $p = MIN_FEE;
    } elseif ($p >= MAX_FEE) {
        $p = MAX_FEE;
    } else {
        $p = round($p / 100) * 100;
    }
    return $p;
}

/**
 * 获取毫秒级时间戳
 * @return float
 */
function microtime_time() {
    list($sec, $micro) = explode(" ", microtime());
    return ceil(((float) $sec + (float) $micro) * 1000);
}

function getNextSleepSecond()
{
    $date1 = time();
    $date2 = date('Y-m-d H:i:00', strtotime('+1 min', $date1));
    $date2 = strtotime($date2);
    $s = $date2 - $date1;

    $date1 = date('Y-m-d H:i:s', $date1);
    $date2 = date('Y-m-d H:i:00', $date2);

    echo("currtime:$date1   nexttime:$date2   timediff:$s\n");
    return $s;
}


function parseSpecial($res)
{
    if ($res && is_array($res)) {
        $data = array("scene_id" => $res["scene_id"],
            "emp_id" => $res["emp_id"],
            "name" => $res["name"],
            "starttime" => $res["starttime"],
            "endtime" => $res["endtime"],
            "status" => $res["status"],
            "createtime" => $res["createtime"]
        );

        return $data;
    } else {
        return NULL;
    }
}

/*********处理场次************/
//状态（1 待拍,2 拍卖中,3 完毕）
function getOrderScene($strtime)
{
    global $dbm;
    $sql = "select * from `au_order_scene` where `status`=1 and `starttime`='$strtime' limit 1";
    $res = $dbm->fetchOne($sql);
    return parseSpecial($res);
}

function sceneStart($id)
{
    global $dbm;
    $sql = "update `au_order_scene` set `status`=2 where `scene_id`=$id";
    $dbm->query($sql);
}

function sceneEnd($id)
{
    global $dbm;
    $sql = "update `au_order_scene` set `status`=3, `endtime`=now() where `scene_id`=$id";
    $dbm->query($sql);
}

function getOrderList($scene_id)
{
    global $dbm;
    $sql = "select * from `au_order` where `scene_id`=$scene_id and `status`=3 order by `rank` asc";
    echo $sql . "\n";
    $data = $dbm->fetchAll($sql);
    return $data;
}

/*********处理场次************/


/*********处理拍单************/
//拍单状态(1审核中,2审核驳回,3投标中,301等待竞标,4竟标中,5成效待确认,6拍单流拍,7待签约,8付付首款,9待过户,10过户中,11待付尾款,12拍单失败,13拍单成功)]
function orderStart($id)
{
    global $db, $dbm, $sid;
    //$sql = "SELECT * FROM (SELECT price,dealer_id FROM au_bid_log WHERE order_id='{$id}' GROUP BY dealer_id ORDER BY createtime DESC) AS price ORDER BY price DESC LIMIT 0,1";
    $sql = "SELECT price,dealer_id FROM au_bid_log a INNER JOIN (SELECT MAX(bid_id) bid_id FROM au_bid_log WHERE order_id='{$id}' GROUP BY dealer_id) b ON a.`bid_id` = b.bid_id ORDER BY price DESC,createtime LIMIT 1;";
    $res = $dbm->fetchOne($sql);
    $sql = "update `au_order` set bid_best_price='{$res['price']}',bid_best_dealer_id='{$res['dealer_id']}', `bidding_start_time`=now(), `status`=4 where `order_id`=$id";

    echo "|$sid|:" . '记入投标最高车商:' . $sql . "\n";
    $dbm->query($sql);

    if ($res['dealer_id']) {
        $sql = "call paimai_refund('{$id}' , '{$res['dealer_id']}' );";
        echo "|$sid|:" . '解冻投标阶段的保证金:' . $sql . "\n";
        $dbm->query($sql);
    } else {
        echo "|$sid|:" . '无人投标!' . "\n";
    }
    //更新拍单进度状态
    $now = microtime_time();
    $orderInfo = $db->where('order_id', $id)->getOne('order','car_id');
    $db->insert('order_trace_log_list',[
        'order_id' => $id,
        'car_id'   => $orderInfo['car_id'],
        'emp_name' => '--',
        'action_no' => 1007,
        'action_name' => '开始竞拍',
        'createtime' => $now
    ]);
    $carInfo = $db->where('car_id', $orderInfo['car_id'])->getOne('cars','owner_id');
    $db->insert('car_trace_log_list',[
        'owner_id'   => $carInfo['owner_id'],
        'car_id' => $orderInfo['car_id'],
        'emp_name' => '--',
        'action_no' => 1013,
        'action_name' => '开始竞拍',
        'createtime' => $now
    ]);
}

function orderWaitBidding($ids)
{
    global $sid, $dbm;
    $new_ids = array();
    foreach ($ids as $key => $row) {
        array_push($new_ids, $row['order_id']);
    }
    $new_ids = implode(',', $new_ids);
    $sql = "update `au_order` set `status`=301 where `order_id` in ({$new_ids})";
    echo "|$sid|:" . '场内全部拍单列表:' . $new_ids . "\n";
    $dbm->query($sql);
}

function getOrderTimes($id)
{
    global $dbm;
    $sql = "select * from `au_order` where `order_id`=$id";
    $res = $dbm->fetchOne($sql);
    if ($res) return $res['act_elapsed_time'];
    else return 0;
}

function orderEnd($id)
{
    global $db, $redis, $sid;
    $lock_key = "bidding_".$id."_lock";
    $r = $redis->lock_wait($lock_key,60);
    if($r){
        $sql = "call paimai_complete($id);";
        echo "|$sid|:" . '拍单结束:' . $sql . "\n";
        $db->rawQuery($sql);

        //更新拍单进度状态
        $now = microtime_time();
        $orderInfo = $db->where('order_id', $id)->getOne('order','car_id');
        $db->insert('order_trace_log_list',[
            'order_id' => $id,
            'car_id'   => $orderInfo['car_id'],
            'emp_name' => '--',
            'action_no' => 1008,
            'action_name' => '竞拍结束',
            'createtime' => $now
        ]);
        $carInfo = $db->where('car_id', $orderInfo['car_id'])->getOne('cars','owner_id');
        $db->insert('car_trace_log_list',[
            'owner_id'   => $carInfo['owner_id'],
            'car_id' => $orderInfo['car_id'],
            'emp_name' => '--',
            'action_no' => 1014,
            'action_name' => '竞拍结束',
            'createtime' =>$now
        ]);
        updateCommision($id);
        $redis->unlock($lock_key);
    }
}

/**
 * 更新交易服务费
 * @param $id
 */
function updateCommision($id)
{
    global $db;
    $orderInfo = $db->where('order_id', $id)->getOne('order');
    //$best_price = $orderInfo['bidding_best_price'] > $orderInfo['bid_best_price'] ? $orderInfo['bidding_best_price'] : $orderInfo['bid_best_price'];
    if($orderInfo['bidding_best_price'] > $orderInfo['bid_best_price']){
        $best_price = $orderInfo['bidding_best_price'];
        $success_dealer_id = $orderInfo['bidding_best_dealer_id'];
    }else{
        $best_price = $orderInfo['bid_best_price'];
        $success_dealer_id = $orderInfo['bid_best_dealer_id'];
    }
    $commision = get_commision($best_price);
    $db->where('order_id', $id)->update('order', [
        'success_price' => $best_price,
        'success_dealer_id' => $success_dealer_id,
        'commision' => $commision
    ]);
}

//处理违约重拍
function breachRedo($id)
{
    global $db;
    $order = $db->where('order_id', $id)->getOne('order');
    $car = $db->where('car_id', $order['car_id'])->getOne('cars');
    //拍单是否违约重拍
    if ($car['is_dealer_breach'] == 1) {
        echo $order['order_id'] . " - " . $order['order_no'] . "违约重拍处理\n";
        $old_order = $db->rawQueryOne("SELECT * FROM `au_order` WHERE `car_id`='{$order['car_id']}' ORDER BY order_id DESC limit 1,1");
        $best_price = $order['bidding_best_price'] > $order['bid_best_price'] ? $order['bidding_best_price'] : $order['bid_best_price'];
        $success_dealer_id = $order['bidding_best_price'] > $order['bid_best_price'] ? $order['bidding_best_dealer_id'] : $order['bid_best_dealer_id'];
        $now = date("Y-m-d H:i:s");
        //判断车辆来源
        if ($car['car_source'] == 1) { //4S店车源
            //更新拍单状态
            $data = [];
            $data['status'] = 8;
            $data['success_price'] = $best_price;
            $data['success_dealer_id'] = $success_dealer_id;
            $data['return_check_status'] = 5;
            $data['first_money'] = $old_order['first_money'];
            if ($car['pay_status'] == 2) {
                $data['first_pay_status'] = 1;
            }
            if ($car['delivery_mode'] == 1) { //先付款后验车
                //添加财务收款信息,
                //$data['dealer_pay_status'] = 1;
                $db->insert('proceeds_log', [
                    'order_id' => $id,
                    'createtime' => $now
                ]);
            } else { //先验车后付款
                $data['check_car_status'] = 1;
                //添加车商验车信息,待验车order_id,dealer_id,createtime,check_limit_time
                $dealer_id = $order['bidding_best_price'] > $order['bid_best_price'] ? $order['bidding_best_dealer_id'] : $order['bid_best_dealer_id'];
                $limit_time = date("Y-m-d H:i:s", strtotime("+1 day"));
                $db->insert('car_dealer_check', [
                    'order_id' => $id,
                    'dealer_id' => $dealer_id,
                    'createtime' => $now,
                    'check_limit_time' => $limit_time
                ]);
            }
        } else { //个人车源
            $data = [];
            $data['status'] = 8;
            $data['success_price'] = $best_price;
            $data['success_dealer_id'] = $success_dealer_id;
            $data['first_money'] = $old_order['first_money'];
            $old_price = $old_order['success_price'] + $old_order['company_subsidies'];
            if ($car['three_in_one'] == 1) {
                if ($car['pay_status'] > 1) {
                    $data['first_pay_status'] = 1;
                }
                if ($best_price > $old_price) {
                    $data['tail_money'] = $old_order['tail_money'] + ($best_price - $old_price);
                } else {
                    $data['tail_money'] = $old_order['tail_money'];
                }
            } else {
                $data['first_pay_status'] = 1;
                $data['tail_money'] = $best_price > $old_price ? $best_price : $old_price ;
            }

            if ($car['delivery_mode'] == 1) { //先付款后验车
                //$data['dealer_pay_status'] = 1;
                //添加收款记录
                $db->insert('proceeds_log', [
                    'order_id' => $id,
                    'createtime' => $now
                ]);
            } else { //先验车后付款
                //更改拍单验车状态为待验车
                $data['check_car_status'] = 1;
                //添加车商验车信息,待验车 order_id,dealer_id,createtime,check_limit_time
                $dealer_id = $order['bidding_best_price'] > $order['bid_best_price'] ? $order['bidding_best_dealer_id'] : $order['bid_best_dealer_id'];
                $limit_time = date("Y-m-d H:i:s", strtotime("+1 day"));
                $db->insert('car_dealer_check', [
                    'order_id' => $id,
                    'dealer_id' => $dealer_id,
                    'createtime' => $now,
                    'check_limit_time' => $limit_time
                ]);
            }
        }
        $db->where('order_id', $id)->update('order', $data);
        //违约重拍这里发券
        drawRaffle($id);
    }
}

function drawRaffle($id)
{
    global $db;
    $now_date = date('Y-m-d H:i:s');
    //获取所在分公司
    $order = $db->where('order_id', $id)->getOne('order');
    $car = $db->where('car_id', $order['car_id'])->getOne('cars');
    $branch_id = $db->where('city_code', $car['location_area'])->getValue('branch_city', 'branch_id');
    //发放抽奖卡券
    $aid = is_have_activity($now_date, $branch_id, 1);
    sendCard($id,$aid);
    //发放抽代金券卡券
    $aid = is_have_activity($now_date, $branch_id, 2);
    sendCard($id,$aid);
}

function sendCard($id,$aid)
{
    global $redis;
    if($aid<1) return false;
    $arr_num = get_denomination_type($aid);
    $price = get_rand_price($arr_num);
    //加锁 堵塞
    $key = "order_binding_card";
    $redis->lock_wait($key);
    $cid = get_one_card($price, $aid);
    if ($cid) {
        $ok = binding($cid, $id);
        $redis->unlock($key);
    } else {
        //该面值卡券已全部分发出去，递归需要重新去池子里面找
        $redis->unlock($key);
        $new_arr = array();
        foreach ($arr_num as $k => $v) {
            if ($v['price'] != $price) {
                $new_arr[] = $v;
            }
        }
        if (count($new_arr) < 1) {
            return 0;
        }
        $ok = again_search($new_arr, $aid, $id);
        $redis->unlock($key);
    }
    return $ok;
}

/**
 * 该面值卡券已全部分发出去，需要重新去池子里面找
 * @return
 */
function again_search($arr, $aid, $order_id)
{
    global $redis;
    $key = "order_binding_card";
    $price = $this->get_rand_price($arr);
    //根据获取随机面额找到一条未分发卡券
    $redis->lock_wait($key);
    $cid = $this->get_one_card($price, $aid);
    if ($cid) {
        //把这张卡券绑定订单
        $ok = binding($cid, $order_id);
        $redis->unlock($key);
        return $ok;
    } else {
        $redis->unlock($key);
        $new_arr = array();
        foreach ($arr as $k1 => $v1) {
            if ($v1['price'] != $price) {
                $newarr[] = $v1;
            }
        }
        if (count($new_arr) < 1) {
            return 0;
        }
        return again_search($new_arr, $aid, $order_id);
    }
}


/**
 *  给卡券绑定订单
 * @return  int
 */
function binding($cid, $order_id)
{
    global $db;
    $ok = $db->where('cid', $cid['cid'])->update('scratch_card', ['order_id' => $order_id]);
    return $ok;
}

/**
 * 当前时间有没有对应活动类型的抽奖活动
 * @return  int 活动id
 */
function is_have_activity($nowdate, $branch_id, $type)
{
    global $db;
    $aid = $db->where('starttime', $nowdate, '<=')
        ->where('endtime', $nowdate, ">=")
        ->where('branch_id', $branch_id)
        ->where('type', $type)
        ->orderBy('createtime', 'DESC')
        ->getOne('activity');
    return $aid['aid'];
}

/**
 * 获取指定活动下所有类型卡卷的数量
 * @param int $aid
 * @return  array
 */
function get_denomination_type($aid)
{
    global $db;
    return $db->where('a_id', $aid)->groupBy('price')->orderBy('price', "DESC")->get('scratch_card', null, 'count(1) as num,price');
}

/**
 * 随机算法 有点恶心
 * @return  String 随机出来的面额
 **/
function get_rand_price($arr_num)
{
    //各种面额算出总数
    $allnum = 0;
    foreach ($arr_num as $value) {
        $allnum += $value['num'];
    }
    $aaa = array();
    $all_num = '';
    $i = 0;
    foreach ($arr_num as $v) {
        if ($i == 0) {
            //把面值高的概率调小一点
            //现在调整，代码不注释
            $aaa[$i]['a_num'] = ceil(($v['num'] / $allnum) * 1000000);
            $all_num += $aaa[$i]['a_num'];
        } else {
            $aaa[$i]['a_num'] = ceil(($v['num'] / $allnum) * 1000000);
            $all_num += $aaa[$i]['a_num'];
        }
        $aaa[$i]['price'] = $v['price'];
        $i++;
    }
    $bbb = array();
    //各种面额闭合区间
    foreach ($aaa as $k1 => $v1) {
        $bbb[$k1]['start_num'] = 1 + $bbb[$k1 - 1]['end_num'];
        $bbb[$k1]['end_num'] = $v1['a_num'] + $bbb[$k1 - 1]['end_num'];
        $bbb[$k1]['price'] = $v1['price'];
    }
    //随机1-$all_num区间的数
    $id = mt_rand(1, $all_num);
    foreach ($bbb as $k2 => $v2) {
        if ($id > $v2['start_num'] && $id < $v2['end_num']) {
            return $v2['price'];
        }
    }
}

/**
 * 根据获取随机面试找到一条卡券
 * @return  array 卡券id
 */
function get_one_card($price, $aid)
{
    global $db;
    $one_card = $db->where('price', $price)->where('order_id', 0)->where('a_id', $aid)->getOne('scratch_card', 'cid');
    return $one_card;
}

//获得场内定时拍单
function orderTimingOrder($sid)
{
    global $dbm;
    $sql = "select * from `au_order` where `scene_id`=$sid and bidding_start_time!='0000-00-00 00:00:00' and status=301";
    $data = $dbm->fetchAll($sql);
    return $data;
}

//获得场内当前时间定时拍单
function orderTimingOrderByTimer($sid, $timer = false)
{
    global $dbm;
    if (!$timer)
        $cur_time_id = date("Y-m-d H:i:00");
    else
        $cur_time_id = $timer;
    $sql = "select * from `au_order` where `scene_id`=$sid and bidding_start_time='$cur_time_id' and status=301";
    $data = $dbm->fetchAll($sql);
    return $data;
}

//获得场内所有非定时拍单
function orderNotTimingOrder($sid)
{
    global $dbm;
    $sql = "select * from `au_order` where `scene_id`=$sid and bidding_start_time='0000-00-00 00:00:00' and status=301";
    $data = $dbm->fetchAll($sql);
    return $data;
}

//获得指定时间点的所有场
function getScenes($timer)
{
    global $dbm;
    $sql = "select * from `au_order_scene` where `status`=1 and `starttime`='$timer'";
    $data = $dbm->fetchAll($sql);
    return $data;
}

//等待进程id列表结束$pids=[pid1,pid2,pid...];
function wait_pids($pids)
{
    while (count($pids) > 0) {
        //echo 'aaa',"\n";
        foreach ($pids as $key => $pid) {
            $res = pcntl_waitpid($pid, $status, WNOHANG);

            if ($res == -1 || $res > 0)
                unset($pids[$key]);
        }
        sleep(1);
    }
}

//获得数据库操作对象
function get_mysql_handler()
{
    $dbm = new Mysql(MYSQLDB, MYSQLDB_USER, MYSQLDB_PWD, MYSQLDB_DATABASE, 'utf8', false);
    return $dbm;
}

//获得更好的数据库操作对象
function get_mysqli_db()
{
    $db = new MysqliDb([
        'host' => MYSQLDB,
        'username' => MYSQLDB_USER,
        'password' => MYSQLDB_PWD,
        'db' => MYSQLDB_DATABASE,
        'prefix' => MYSQLDB_PREFIX,
        'charset' => 'utf8'
    ]);
    return $db;
}

function get_redis()
{
    $redis = new PhpRedis();
    return $redis;
}
