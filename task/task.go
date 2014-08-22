package task

func Run(taskName string, taskFile string) {
	mruby := NewMRuby()
	defer mruby.Close()
	dockerInit(mruby)
	mruby.FromFile(taskFile)
	mruby.FromString(taskName)
}
