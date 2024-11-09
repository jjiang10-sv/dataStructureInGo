package thres_sync

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// data inconsistency in concurrency.
type worker struct {
	Id uint16
}

type ProduceAndConsume struct {
	data         []struct{}
	buffLen      uint16
	in           uint16
	out          uint16
	producers    []*worker
	consumers    []*worker
	wg           sync.WaitGroup
	proMutex     sync.Mutex
	conMutex     sync.Mutex
	consumeChan  chan struct{}
	producerChan chan struct{}
}

func produceAndConsume() {
	bufLen, prodNum, consumNum := 10, 15, 15
	pc := NewProducerAndConsumer(uint16(bufLen), uint16(prodNum), uint16(consumNum))

	pc.wg.Add(prodNum + consumNum)

	for i := 0; i < bufLen; i++ {
		pc.consumeChan <- struct{}{}
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*4))
	defer cancel()
	for id, p := range pc.producers {

		p = NewWorker(uint16(id))
		go pc.produce(p, ctx)
		//time.Sleep(2*time.Second)
	}
	for id, c := range pc.consumers {
		c = NewWorker(uint16(id))
		go pc.consume(c, ctx)
		//time.Sleep(time.Second)
	}

	pc.wg.Wait()

}

func NewProducerAndConsumer(dataLen, producerNum, consumeNum uint16) *ProduceAndConsume {
	return &ProduceAndConsume{
		data:         make([]struct{}, dataLen),
		buffLen:      dataLen,
		producers:    make([]*worker, producerNum),
		consumers:    make([]*worker, consumeNum),
		wg:           sync.WaitGroup{},
		consumeChan:  make(chan struct{}, dataLen),
		producerChan: make(chan struct{}, dataLen),
	}

}

func NewWorker(id uint16) *worker {
	return &worker{Id: id}
}

func (pc *ProduceAndConsume) produce(worker *worker, ctx context.Context) {

	defer func() {
		pc.wg.Done()
	}()
	//defer pc.wg.Done()
	for range pc.consumeChan {
		select {
		case <-ctx.Done():
			return
		default:
			pc.proMutex.Lock()
			in := pc.in
			pc.data[in] = struct{}{}
			fmt.Println("producer ", worker.Id, " produce in ", in)
			pc.in = (in + 1) % pc.buffLen
			fmt.Println("producer ", worker.Id, " updated  in to ", pc.in)
			pc.proMutex.Unlock()
			pc.producerChan <- struct{}{}
		}

		//time.Sleep(2*time.Second)
	}

}

// pc.in == buffLen-1; the buff is full; pc.out ==
func (pc *ProduceAndConsume) consume(worker *worker, ctx context.Context) {
	defer pc.wg.Done()
	for range pc.producerChan {

		select {
		case <-ctx.Done():
			return
		default:
			pc.conMutex.Lock()
			out := pc.out
			fmt.Println("consumer ", worker.Id, " consumed  ", out)
			pc.out = (out + 1) % pc.buffLen
			fmt.Println("consumer ", worker.Id, " update  out to  ", pc.out)
			pc.conMutex.Unlock()
			pc.consumeChan <- struct{}{}

		}
	}

}
