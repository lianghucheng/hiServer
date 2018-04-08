package engine

import (
	"log"
	"hiServer/fetcher"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}


func (e *ConcurrentEngine)Run(seeds ...Request){
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]


			body, err := fetcher.Fetch(r.Url)
			if err != nil {
				log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
			}

			log.Printf("%s\n", body)
			parseResult := r.ParserFunc(body)
			requests = append(requests, parseResult.Requests...)
			for _, item := range parseResult.Items {
				log.Printf("Got item %v", item)
			}
	}
	log.Println("run end")
}