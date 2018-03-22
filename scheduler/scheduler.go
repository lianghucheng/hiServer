package scheduler

import "hiServer/engine"

type Scheduler struct {
	workerChan chan engine.RequestTest
}

func (s *Scheduler) ConfigureMasterWorkerChan(c chan engine.RequestTest){
	s.workerChan=c
}

func (s *Scheduler)Submit(r engine.RequestTest){
	go func(){
		s.workerChan<-r
	}()
}
