package task

/*
#include <mruby.h>

void proc_call(mrb_state *mrb, void *p) {
  mrb_value proc = *(mrb_value *)p;
  mrb_funcall(mrb, proc, "call", 0);
}

*/
import "C"

import (
	"time"
	"unsafe"
)

type task struct {
	mruby mruby
	ts    []taskSchedule
}

func (this *task) Load(fileName string) {
	this.mruby.FromFile(fileName)
}

func (this *task) RunString(s string) {
	this.mruby.FromString(s)
}

func (this *task) Close() {
	this.mruby.Close()
}

func (this *task) ProcCall(proc proc) {
	C.proc_call(this.mruby.mrb, unsafe.Pointer(proc))
}

func waitTick(ch chan proc, tick time.Duration, proc proc) {
	time.Sleep(tick)
	ch <- proc
}

func (this *task) NextTick() <-chan proc {
	var min time.Duration
	var minproc proc
	min = 24 * time.Hour * 365 * 10
	now := time.Now()
	for _, ts := range this.ts {
		tick := ts.nextTick(now)
		if min > tick {
			min = tick
			minproc = ts.proc
		}
	}

	ch := make(chan proc)
	go waitTick(ch, min, minproc)
	return ch
}

func New() *task {
	t := new(task)
	t.ts = make([]taskSchedule, 0)
	t.mruby = NewMRuby()
	dockerInit(t.mruby)
	cronInit(t.mruby, t)
	return t
}
