package helper

import (
	"reflect"
	"strings"
)

func ConvertSliceToPostgresArray(slice []string) string {
	arrayString := "{"
	for i, value := range slice {
		// Escape any double quotes in the string value
		value = strings.Replace(value, `"`, `\"`, -1)
		// Append each string value to the arrayString
		arrayString += `"` + value + `"`
		// Add a comma separator except for the last element
		if i < len(slice)-1 {
			arrayString += ","
		}
	}
	arrayString += "}"
	return arrayString
}

func ParsePostgresArray(src string) []string {
	// Remove curly braces from the string
	src = strings.Trim(src, "{}")
	// Split the string by comma to get individual values
	values := strings.Split(src, ",")
	// Trim whitespace from each value
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}

func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				return false
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() != 0 {
				return false
			}
		case reflect.Bool:
			if field.Bool() != false {
				return false
			}
		// Add cases for other types as needed
		default:
			// Handle other types if necessary
		}
	}
	return true
}
