
//Created Chapter 8.4
package forms

//Errors type for holding validation errors
type errors map[string][]string

//Add method
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//Get method
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}


