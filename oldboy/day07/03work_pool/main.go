package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
 * 1. 开启1个 goroutine 循环生成随机数，发送到 jobChan
 * 2. 开启24个 goroutine 从 jobChan 中取出随机数字各位的和，将结果发送到 resultChan
 * 3. 主goroutine从resultChan中取出结果并打印
 */

// x <- chan int: 从 int 类型的通道中取数据
// chan int <- x：往 int 类型的通道中写数据

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func genData(d chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		d <- newJob
		time.Sleep(time.Second)
	}
}

func hanleData(d <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		jj := <-d
		sum := int64(0)
		n := jj.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    jj,
			result: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go genData(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go hanleData(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Println(result.job.value, result.result)
	}
	wg.Wait()
}
