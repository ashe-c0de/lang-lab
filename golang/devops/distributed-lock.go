package main

import (
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 1. 初始化 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // 请确保你的 Redis 已启动
	})

	// 2. 初始化 Redsync
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	r := gin.Default()

	r.GET("/grab", func(c *gin.Context) {
		mutexName := "my-global-lock"
		// 有效期 6 秒
		mutex := rs.NewMutex(mutexName, redsync.WithExpiry(6*time.Second))

		// 1. 尝试获取锁
		if err := mutex.Lock(); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "系统繁忙"})
			return
		}

		// 2. 准备看门狗控制上下文
		watchdogCtx, cancel := context.WithCancel(context.Background())

		// 确保无论如何最后都会：1.停止看门狗 2.释放锁
		defer func() {
			cancel() // 停止看门狗协程
			if ok, err := mutex.Unlock(); !ok || err != nil {
				log.Printf("最终释放锁结果: %v, err: %v", ok, err)
			}
		}()

		// 3. 异步启动看门狗
		go StartWatchdog(watchdogCtx, mutex, 6*time.Second)

		// 4. 执行耗时业务逻辑 (10s > 6s)
		log.Println("获取锁成功，开始 10s 长耗时业务...")
		time.Sleep(10 * time.Second)

		c.JSON(http.StatusOK, gin.H{
			"message": "抢购成功！",
		})
	})

	r.Run(":8080")
}

// StartWatchdog 启动一个看门狗协程
// ctx: 用于通知看门狗停止（业务执行完后 cancel）
// mutex: redsync 的互斥锁实例
func StartWatchdog(ctx context.Context, mutex *redsync.Mutex, d time.Duration) {
	// 建议续约周期为有效期的一半，例如 6s 的有效期，每 3s 续约一次
	ticker := time.NewTicker(d / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 尝试续约
			ok, err := mutex.Extend()
			if err != nil || !ok {
				log.Printf("[Watchdog] 续约失败 (可能锁已失效或网络问题): %v", err)
				return
			}
			log.Println("[Watchdog] 续约成功")
		case <-ctx.Done():
			// 业务逻辑执行完毕，正常退出协程
			log.Println("[Watchdog] 业务结束，停止续约")
			return
		}
	}
}
