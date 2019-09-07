package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (engine *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	engine.Scheduler.SetChannel(in)


	for i := 0; i <= engine.WorkerCount; i++ {
		CreateWorker(in, out)
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

func CreateWorker(in chan Request, out chan ParseResult) {
	go func() {
		for ; ; {
			r := <-in
			result, err := Worker(r)
			if err != nil {
				continue;
			}
			out <- result
		}
	}()

}
