//Created Chapter 8.4

package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

//Email validation regex - From W3C
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}" +
	"[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//Create custom for struct
type Form struct {
	url.Values
	Errors errors
}

//To initialise a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Check specific fields.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Cannot be blank")
		}
	}
}

//Max Length method
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("Field too long, %d is maximum", d))
	}
}

//Permitted values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

//Min Length
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("This field is too short, %d minimum", d))
	}
}

//Matches Pattern
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

//Valid method
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
