package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	workerBits  uint8 = 10                      //工作节点 ID 的位数（10 位），最多支持 1024 个工作节点。
	seqBits     uint8 = 12                      //序列号的位数（12 位），每个节点每毫秒最多生成 4096 个 ID。
	workerMax   int64 = -1 ^ (-1 << workerBits) //计算出的最大工作节点 ID（1023）。
	seqMax      int64 = -1 ^ (-1 << seqBits)    //计算出的最大序列号（4095）。
	timeShift   uint8 = workerBits + seqBits    //时间戳的位移，用于生成 ID 的高位部分。
	workerShift uint8 = seqBits                 //工作节点 ID 的位移，用于生成 ID 的中间部分。
	epoch       int64 = 1567906170596           //自定义的起始时间戳（毫秒级）。
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	seq       int64
}

func NewWorker(worker int64) (*Worker, error) {
	if worker < 0 || worker > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	return &Worker{
		timestamp: 0,
		seq:       0,
		workerId:  worker,
	}, nil
}

func (w *Worker) Next() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.seq = (w.seq + 1) & seqMax
		fmt.Println(w.seq)
		if w.seq == 0 {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.seq = 0
	}
	w.timestamp = now

	fmt.Println(now - epoch)
	fmt.Println(timeShift)
	fmt.Println(int64((now - epoch) << int64(timeShift)))
	ID := int64((now-epoch)<<int64(timeShift) | (w.workerId << int64(workerShift)) | (w.seq))
	return ID
}

func main() {
	fmt.Println(2 << 10)
	worker, err := NewWorker(1)
	if err != nil {
		log.Fatalf("new worker err:%v", err)
		return
	}

	log.Println("---worker 1---")
	id := worker.Next()
	log.Printf("id:%d", id)
	id = worker.Next()
	log.Printf("id:%d", id)

}
