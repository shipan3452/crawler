package engine

import (
	"awesomeProject/crawler/fetcher"
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

func (engine ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	engine.Scheduler.SetChannel(in)
	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}



	for i = 0; i <= engine.WorkerCount; i++ {
		CreateWorker(in, out)
	}

	for ;; {
		result := <-out
		for _, r = range result.Requests {
			engine.Scheduler.Submit(r)
		}
	}



	for _, item = range result.Items {
		fmt.Printf("%s",item)
	}

}

func CreateWorker(in chan Request, out chan ParseResult) {
	for ; ; {
		r := <-in
		go func() {
			result,err:=Worker(r)
			if(err!=nil){
				continue;
			}
			out <- result
		}()


	}
}
