package contact

type Results map[string]*Result

func (that Results) Get(name string) *Result {
	return that[name]
}

func (that Results) Set(name string, value *Result) {
	that[name] = value
}

func (that Results) Error() string {
	return ""
}
