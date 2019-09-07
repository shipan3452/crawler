package scheduler

import "awesomeProject/crawler/engine"

type SimpleScheduler struct {
	c chan engine.Request
}


func (scheduler SimpleScheduler) Submit(request engine.Request) {
	scheduler.c <- request
}

func (scheduler SimpleScheduler) SetChannel(c chan engine.Request) {
	scheduler.c = c
}
