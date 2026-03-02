package main

import (
	"fmt"
	"math/rand"
	"time"
)

// cookTask: 模拟任务
// 关键点 1: 任务结束后，必须 defer close(ch) 通知接收方“数据流结束了”
func cookTask(name string, duration time.Duration, ch chan<- string) {
	// 确保即使发生 panic，channel 也会被关闭（防止接收方死锁）
	defer close(ch)

	time.Sleep(duration)
	ch <- fmt.Sprintf("%s 做好了!", name)
	// 函数结束，defer 执行，channel 自动关闭
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. 创建 channel
	chNoodle := make(chan string)
	chSteak := make(chan string)
	chBread := make(chan string)

	// 2. 启动任务
	go cookTask("煮面", time.Duration(rand.Intn(3)+1)*time.Second, chNoodle)
	go cookTask("煎牛排", time.Duration(rand.Intn(3)+1)*time.Second, chSteak)
	go cookTask("烤面包", time.Duration(rand.Intn(3)+1)*time.Second, chBread)

	fmt.Println("👨‍🍳 厨师们开始做菜了...")
	fmt.Println("🤵 服务员开始监听...")

	// 关键点 2: 维护一个“活跃通道”计数器
	// 初始有 3 个 channel 需要监听
	activeChannels := 3

	// 关键点 3: 循环监听，直到 activeChannels 变为 0
	for activeChannels > 0 {
		select {
		// 关键点 4: 使用 comma-ok 惯用语法接收数据
		// msg: 接收到的数据
		// ok: 如果 channel 未关闭且有数据，ok=true; 如果 channel 已关闭且无数据，ok=false
		case msg, ok := <-chNoodle:
			if !ok {
				// channel 已关闭，不再监听它
				chNoodle = nil // 将 channel 置为 nil，select 会自动跳过 nil 的 case
				activeChannels--
				fmt.Println("🍜 煮面通道已关闭 (剩余任务:", activeChannels, ")")
				continue // 跳过本次循环剩余逻辑，直接进入下一次 select
			}
			fmt.Println("🍜 收到消息:", msg)

		case msg, ok := <-chSteak:
			if !ok {
				chSteak = nil
				activeChannels--
				fmt.Println("🥩 煎牛排通道已关闭 (剩余任务:", activeChannels, ")")
				continue
			}
			fmt.Println("🥩 收到消息:", msg)

		case msg, ok := <-chBread:
			if !ok {
				chBread = nil
				activeChannels--
				fmt.Println("🍞 烤面包通道已关闭 (剩余任务:", activeChannels, ")")
				continue
			}
			fmt.Println("🍞 收到消息:", msg)
		}
	}

	fmt.Println("✅ 所有菜都上齐了，所有通道已关闭，服务员优雅下班！")
}
