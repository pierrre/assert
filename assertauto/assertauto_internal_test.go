package assertauto

func Update(u bool) Option {
	return update(u)
}

func Directory(d string) Option {
	return directory(d)
}

func SetTestName(n string) Option {
	return testName(n)
}
