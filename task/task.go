package task

type task struct {
	mruby mruby
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

func New() task {
	t := task{}
	t.mruby = NewMRuby()
	dockerInit(t.mruby)
	return t
}
