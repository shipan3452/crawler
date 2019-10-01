package scheduler

import "awesomeProject/crawler/engine"

type SimpleScheduler struct {
	c chan engine.Request
}

func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		scheduler.c <- request
	}()

}

func (scheduler *SimpleScheduler) SetChannel(c chan engine.Request) {
	scheduler.c = c
}

func (scheduler *SimpleScheduler) WorkerIsReady(c chan engine.Request) {
	panic("implement me")
}