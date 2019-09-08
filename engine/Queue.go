package engine

import "fmt"

type QueueEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (engine *QueueEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	
	for i := 0; i <= engine.WorkerCount; i++ {
		CreateAWorker( out,engine.Scheduler)
	}

	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}

	for ; ; {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("%s\n", item)
		}
		for _, r := range result.Requests {
			go func() {
				engine.Scheduler.Submit(r)
			}()
		}

	}
}

func CreateAWorker( out chan ParseResult,scheduler Scheduler) {
	c:=make(chan Request);
	go func() {
		for ; ; {
			scheduler.WorkerIsReady(c)
			r := <-c
			result, err := Worker(r)
			if err != nil {
				continue;
			}
			out <- result
		}
	}()
}
