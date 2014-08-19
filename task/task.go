package task

func Run(script string, method string) {
	mruby := NewMRuby()
	dockerInit(mruby)
	mruby.LoadString(script)
	mruby.LoadString(method)
}
