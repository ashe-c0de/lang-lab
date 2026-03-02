type Once struct {
	done uint32      // 核心标志位：0 表示未执行，1 表示已执行
	m    Mutex       // 互斥锁
}

func (o *Once) Do(f func()) {
	// 【第一步】原子读取：快速路径 (Fast Path)
	// 如果 done 已经是 1，直接返回，无需加锁。
	// 这相当于 DCL 的“第一次检查”，但使用的是原子操作，速度极快。
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	// 【第二步】加锁慢速路径
	// 只有当 done != 1 时，才需要竞争锁。
	o.m.Lock()
	defer o.m.Unlock()

	// 【第三步】再次检查 (Inside Lock)
	// 拿到锁后，必须再次检查！因为可能在排队拿锁的过程中，
	// 另一个 goroutine 已经执行完了 f() 并将 done 设为 1 了。
	if o.done == 0 {
		// 【第四步】执行并标记
		// 使用 defer 确保即使 f() panic，done 也会被设置为 1，防止重复执行。
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
