package task

/*
#cgo CFLAGS: -I../mruby/include
#cgo LDFLAGS: -L../mruby/build/host/lib -lmruby -lm

#include <mruby.h>
#include <mruby/compile.h>
#include <mruby/error.h>
#include <mruby/string.h>

void err(mrb_state *mrb, mrbc_context *cxt) {
  if (!mrb->exc) {
    return;
  }
  mrb_value exc = mrb_obj_value(mrb->exc);
//  mrb_value backtrace = mrb_get_backtrace(mrb, exc);
//  puts(mrb_str_to_cstr(mrb, mrb_inspect(mrb, backtrace)));

  mrb_value inspect = mrb_inspect(mrb, exc);
  puts(mrb_str_to_cstr(mrb, inspect));
}

static int64_t _mrb_fixnum(mrb_value a) {return mrb_fixnum(a);}
*/
import "C"

type mruby struct {
	mrb *C.mrb_state
	cxt *C.mrbc_context
}

func NewMRuby() mruby {
	mruby := mruby{}

	mruby.mrb = C.mrb_open()
	mruby.cxt = C.mrbc_context_new(mruby.mrb)

	return mruby
}

func (this *mruby) FromFile(fileName string) {
	file := C.fopen(C.CString(fileName), C.CString("rb"))
	C.mrb_load_file_cxt(this.mrb, file, this.cxt)
	// C.err(this.mrb, this.cxt)
}

func (this *mruby) FromString(script string) {
	C.mrb_load_string_cxt(this.mrb, C.CString(script), this.cxt)
	C.err(this.mrb, this.cxt)
}

func (this *mruby) Close() {
	C.mrbc_context_free(this.mrb, this.cxt)
	C.mrb_close(this.mrb)
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
