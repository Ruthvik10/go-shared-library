package validator

import (
	"net/url"
	"regexp"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

var EmailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (v *Validator) Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func (v *Validator) AddErrors(key, errMsg string) {
	_, exists := v.Errors[key]
	if !exists {
		v.Errors[key] = errMsg
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(key string, ok bool, errMsg string) {
	if !ok {
		v.AddErrors(key, errMsg)
	}
}

func (v *Validator) ValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
