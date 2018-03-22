package engine

import (
	"lhc/crawler/fetcher"
	"log"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	Submit(RequestTest)
	ConfigureMasterWorkerChan(chan RequestTest)
}


func (e *ConcurrentEngine)Run(seeds ...RequestTest){
	var requests []RequestTest
	for _,r:=range seeds{
		requests=append(requests,r)
	}

	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]

		body,err:=fetcher.Fetch(r.UrlTest)
		if err !=nil{
			log.Printf("Fetcher: error "+"fetching url %s: %v",r.UrlTest,err)
		}

		parseResult:=r.ParserFuncTest(body)
		requests=append(requests,parseResult.RequestTest...)
		for _,item:=range parseResult.ItemsTest {
			log.Printf("Got item %v", item)
		}
	}
}