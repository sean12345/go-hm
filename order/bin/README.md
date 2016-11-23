****************************************************************************
## order 服务使用方法
```

处理 开始竞拍的订单
order -s start
order -s start -d 1477561980

处理 结束竞拍的订单
order -s end
order -s end -d 1477561980

或

处理 开始竞拍的订单
order -s=start
order -s=start -d=1477561980

处理 结束竞拍的订单
order -s=end
order -s=end -d=1477561980

```

****************************************************************************
## order 队列格式
```
队列名称 : hammer_start_list_10位时间戳
格式: {"order_id": 1456, "scene_id": 2144, "bidding_start_time": "2016-10-27 17:50:00", "bidding_end_time": "0000-00-00 00:00:00", "est_elapsed_time": 180, "act_elapsed_time": 0, "rank": 1, "is_timing_order":1}
订单ID，会场ID，拍卖开始时间，拍卖结束时间，预计时长，实际时长，排名，是否是定时拍

```
