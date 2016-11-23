<?php
/**
 * redis操作类
 *@author爱民
 */
class PhpRedis
{
    private $redis = null;

    public function __construct($dbindex=false){
        $redis = new \Redis();
        $redis->connect(REDIS_HOST, REDIS_PORT, 1 );
        $this->redis = $redis;
        if($dbindex){
            $this->select($dbindex);
        }
    }

    public function ping()
    {
        return $this->redis->ping();
    }

    public function select($num){
        return $this->redis->select($num);
    }

    public function set($key,$vale,$expire=false){

        $this->redis->set($key,$vale);

        if($expire){
            $this->redis->expire($key,$expire);
        }

    }

    public function get($key){
        return $this->redis->get($key);
    }

    public function del($key){
        return $this->redis->delete($key);
    }

    public function getRedis(){
        return $this->redis;
    }

    public function lock($key){
        return $this->redis->setNx($key,1)?true:false;
    }

    /**
     * @param string $key lockkey
     * @param int $max_delay_times
     * @return bool
     */
    public function lock_wait($key, $max_delay_times = 5)
    {
        $r = $this->lock($key);
        if (!$r) {
            for ($i = 0; $i < $max_delay_times; $i++) {
                usleep(rand(500, 1000));
                $r = $this->lock($key);
                if ($r)
                    break;
                $i++;
            }
        }
        return $r;
    }

    public function unlock($key){
        $this->redis->delete($key);
    }

    public function incr($key){
        $this->redis->incr($key);
    }

    public function rPush($key,$value)
    {
        return $this->redis->rPush($key,$value);
    }
}