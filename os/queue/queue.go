package queue

import (
	"time"
	"sync"
)

const DEFAULT_SIZE = 3
const DEFAULT_TIMEOUT = 1 * time.Second

type Queue struct {
	timeout time.Duration
	// Queue    string // 组
	// PlayLoad string // 序列数据
	// Attempts int    // 次数
	// Delay time.Duration
	Name string
	// ReservedAt  time.Time // 处理时间
	// AvailableAt time.Time // 有效时间
	queue chan Job
	wg    *sync.WaitGroup
}

type Job struct {
	Name string
	Func func(name string)
}

func New(size uint, timeout time.Duration) *Queue {
	queue := &Queue{
		timeout: timeout,
		queue:   make(chan Job, size),
		wg:      &sync.WaitGroup{},
	}

	go queue.Background()

	return queue
}

func Default() *Queue {
	return New(DEFAULT_SIZE, DEFAULT_TIMEOUT)
}

func (self *Queue) Add(name string, v func(name string)) {
	self.queue <- Job{Name: name, Func: v}
}

func (self *Queue) Done() {
	self.wg.Done()
}

func (self *Queue) Background() {
	for {
		select {
		case job := <-self.queue:
			go func() {
				finish := make(chan string)

				self.wg.Add(1)
				go func(finish chan string) {
					job.Func(job.Name)
					finish <- job.Name
				}(finish)

				for {
					select {
					case jobName := <-finish:
						println("finish: " + jobName)
						return
					case <-time.After(self.timeout):
						println("timeout")
						self.wg.Done()
						return
					}
				}
			}()
		}
	}
}

func (self *Queue) Wait() {
	self.wg.Wait()
}
