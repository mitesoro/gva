package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

const (
	randomLen       = 16
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
	letterBytes     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	// 先判断是否是key锁的持有者，是则续约过期时间，不是则尝试获取锁，当key过期则可以获取锁
	lockScript = redis.NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`)
	delScript = redis.NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`)
)

// A RedisLock is a redis lock.
type RedisLock struct {
	rdb     redis.UniversalClient
	seconds uint32 // 锁续约时间
	key     string // 锁名称
	id      string // 持有锁的值
}

// NewRedisLock returns a RedisLock.
func NewRedisLock(rdb redis.UniversalClient, key string) *RedisLock {
	return &RedisLock{
		rdb: rdb,
		key: key,
		id:  randN(randomLen),
	}
}

// Acquire acquires the lock with the given ctx.
func (rl *RedisLock) Acquire(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	resp, err := lockScript.Run(ctx, rl.rdb, []string{rl.key}, []string{
		rl.id, strconv.Itoa(int(seconds)*millisPerSecond + tolerance),
	}).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		log.Println("Error on acquiring lock for ", rl.key, err.Error())
		return false, err
	} else if resp == nil {
		return false, nil
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	log.Println("Unknown reply when acquiring lock for ", rl.key, resp)
	return false, nil
}

// Release releases the lock with the given ctx.
func (rl *RedisLock) Release(ctx context.Context) (bool, error) {
	resp, err := delScript.Run(ctx, rl.rdb, []string{rl.key}, []string{rl.id}).Result()
	if err != nil {
		return false, err
	}

	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

// SetExpire sets the expiration.
func (rl *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}

func randN(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
