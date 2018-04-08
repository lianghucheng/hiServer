package scheduler

import "hiServer/engine"

type Scheduler struct {
	workerChan chan engine.Request
}

func (s *Scheduler) ConfigureMasterWorkerChan(c chan engine.Request){
	s.workerChan=c
}

func (s *Scheduler)Submit(r engine.Request){
	go func(){
		s.workerChan<-r
	}()
}
