package task

/*
#cgo CFLAGS: -I../mruby/include
#cgo LDFLAGS: -L../mruby/build/host/lib -lmruby -lm

#include <mruby.h>
#include <mruby/compile.h>
#include <mruby/error.h>
#include <mruby/string.h>

static int64_t _mrb_fixnum(mrb_value a) {return mrb_fixnum(a);}
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

func mrbToString(mrb *C.mrb_state, Rvalue C.mrb_value) string {
	if Rvalue.tt != C.MRB_TT_STRING {
		panic("Invalid argument. must String.")
	}
	return C.GoString(C.mrb_string_value_ptr(mrb, Rvalue))
}

func mrbSymToString(mrb *C.mrb_state, Rvalue C.mrb_value) string {
	if Rvalue.tt != C.MRB_TT_SYMBOL {
		panic("Invalid argument. must SYMBOL")
	}
	return C.GoString(C.mrb_sym2name(mrb, C.mrb_obj_to_sym(mrb, Rvalue)))
}

func mrbToInt64(mrb *C.mrb_state, Rvalue C.mrb_value) int64 {
	if Rvalue.tt != C.MRB_TT_FIXNUM {
		panic("Invalid argument. must FIXNUM")
	}
	return int64(C._mrb_fixnum(Rvalue))
}

func mrbToUint64(mrb *C.mrb_state, Rvalue C.mrb_value) uint64 {
	if Rvalue.tt != C.MRB_TT_FIXNUM {
		panic("Invalid argument. must FIXNUM")
	}
	return uint64(C._mrb_fixnum(Rvalue))
}

func mrbToBool(mrb *C.mrb_state, Rvalue C.mrb_value) bool {
	if Rvalue.tt == C.MRB_TT_TRUE {
		return true
	}
	if Rvalue.tt == C.MRB_TT_FALSE {
		return false
	}
	panic("Invalid argument. must bool")
	return false
}
