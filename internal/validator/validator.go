package validator

import (
	"fmt"
	"log"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
)

func Validate(data interface{}) (err error) {
	// returns operator for value
	v := reflect.ValueOf(&data).Elem().Elem()
	// get type of payload
	t := reflect.TypeOf(data)
	if t == nil {
		err = fmt.Errorf("User data is required")
		return
	}
	// check if it's a struct
	if t.Kind() != reflect.Struct {
		err = fmt.Errorf("error processing user data")
		return
	}
	//  check each field and validate
	for i := 0; i < t.NumField(); i++ {
		// getting the fields
		field := t.Field(i)
		log.Printf("Field Name :: %v", field.Name)

		// getting the tags
		if tag, ok := field.Tag.Lookup("validate"); ok {
			log.Printf("Tags :: %v", tag)
			tags := strings.Split(tag, ";")
			log.Printf("Tags :: %v", tags)
			n := map[string]string{}

			for _, value := range tags {

				t := strings.Split(value, "=")
				n[t[0]] = t[1]
			}
			log.Printf("Validators :: %+v", n)
			for key, value := range n {
				if key == "required" && value == "true" {
					switch v.Field(i).Interface().(type) {
					case string:
						if v.Field(i).Interface().(string) == "" {
							err = fmt.Errorf("%s is required!", field.Name)
							return
						}
					}
				}
				if key == "max" {
					switch v.Field(i).Interface().(type) {
					case string:
						n, convErr := strconv.Atoi(value)
						if convErr != nil {
							err = fmt.Errorf("Invalid max value for validation tag at field %s", field.Name)
							return
						}
						if len(v.Field(i).Interface().(string)) > n {
							err = fmt.Errorf("%s cannot be more than max value of %d", field.Name, n)
							return
						}
					case int:
						n, convErr := strconv.Atoi(value)
						if convErr != nil {
							err = fmt.Errorf("Invalid max value for validation tag at field %s", field.Name)
							return
						}
						if v.Field(i).Interface().(int) > n {
							err = fmt.Errorf("%s cannot be more than max value of %d", field.Name, n)
							return
						}
					}
				}
				if key == "type" {
					if value == "email" {
						switch v.Field(i).Interface().(type) {
						case string:
							_, err = mail.ParseAddress(v.Field(i).Interface().(string))
							if err != nil {
								err = fmt.Errorf("%s must be a valid email address", field.Name)
								return
							}
						}
					}
				}
			}
		}
	}
	return
}

