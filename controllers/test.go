package controllers

import (
	"hiServer/engine"
	"hiServer/scheduler"
)

//@router /engine/test [*]
func (user *UserController)EngineTest(){
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.Scheduler{},
		WorkerCount:1,
	}
	e.Run(engine.RequestTest{
		UrlTest:"",
		ParserFuncTest:nil,
	})
}