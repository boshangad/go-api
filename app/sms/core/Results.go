package core

import "github.com/boshangad/v1/app/sms/interfaces"

type Results struct {
	results   map[string]interfaces.Result
	okKey     string
	hasFailed bool
}

func (that Results) HasFailed() bool {
	return that.hasFailed
}

func (that *Results) Set(key string, result interfaces.Result) {
	if that.results == nil {
		that.results = map[string]interfaces.Result{}
	}
	if result == nil {
		return
	}
	that.results[key] = result
	if result.IsSuccess() {
		that.okKey = key
	} else {
		that.hasFailed = true
	}
}

func (that Results) Get(key string) (result interfaces.Result) {
	return that.results[key]
}

func (that Results) Each(fn func(k string, v interfaces.Result)) {
	for k, v := range that.results {
		fn(k, v)
	}
}
