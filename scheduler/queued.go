package scheduler

import "awesomeProject/crawler/engine"

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (scheduler *QueuedScheduler) SetChannel(c chan engine.Request) {
	panic("implement me")
}

func (scheduler *QueuedScheduler) Submit(r engine.Request) {
	go func() {
		scheduler.RequestChan <- r
	}()
}

func (scheduler *QueuedScheduler) WorkerIsReady(c chan engine.Request) {
	scheduler.WorkerChan <- c
}

func (scheduler QueuedScheduler) Run() {
	var requestQueue []engine.Request
	var workerQueue []chan engine.Request

	var activeRequest engine.Request
	var activeWorker chan engine.Request
	for {
		if len(requestQueue) > 0 && len(workerQueue) > 0 {
			activeRequest = requestQueue[0]
			activeWorker = workerQueue[0]
		}
		select {
		case r := <-scheduler.RequestChan:
			requestQueue = append(requestQueue, r)
		case w := <-scheduler.WorkerChan:
			workerQueue = append(workerQueue, w)
		case activeWorker <- activeRequest:
			requestQueue = requestQueue[1:]
			workerQueue = workerQueue[1:]
		}
	}
}
