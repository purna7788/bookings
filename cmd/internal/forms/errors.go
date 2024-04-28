package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates custom Form struct
type Form struct {
	url.Values
	Err errors
}

// New initializes form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {

	formData := r.Form.Get(field)
	if formData == "" {
		f.Err.Add(field, "this field cannot be empty")
		return false
	}
	return true
}

// Required function checks if field has any validation errors
func (f *Form) Required(field ...string) {
	for _, v := range field {
		formData := f.Get(v)
		fmt.Printf("field is ---> %s and value is %s\n", v, formData)
		if strings.TrimSpace((formData)) == "" {
			f.Err.Add(v, "this field cannot be empty")
		}
	}

	fmt.Println(f.Err)
}

// Valid method validates errors on submitted form.
func (f *Form) Valid() bool {
	return len(f.Err) == 0
}

// MinLength function validates lenght of the given field
func (f *Form) MinLength(field string, minlength int) bool {

	if len(f.Get(field)) < minlength {
		f.Err.Add(field, " field musts contain minimum lenght of "+strconv.Itoa(minlength))
		return false
	}
	return true
}

// IsEmailValid function validates email address of Email field
func (f *Form) IsEmailValid(field string) bool {

	if !govalidator.IsEmail(f.Get(field)) {
		f.Err.Add(field, "Email address is invalid")
		return false
	}
	return true

}
