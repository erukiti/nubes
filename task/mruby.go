package task

/*
#cgo CFLAGS: -I../mruby/include
#cgo LDFLAGS: -L../mruby/build/host/lib -lmruby -lm

#include <mruby.h>
#include <mruby/compile.h>

*/
import "C"

type mruby struct {
	mrb *C.mrb_state
}

func NewMRuby() mruby {
	mruby := mruby{}

	mruby.mrb = C.mrb_open()
	return mruby
}

func (this *mruby) LoadString(s string) interface{} {
	return C.mrb_load_string(this.mrb, C.CString(s))
}
