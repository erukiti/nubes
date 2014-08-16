package task

func Run(script string, method string) {
	mruby := NewMRuby()
	mruby.LoadString(script)
	mruby.LoadString(method)
}
