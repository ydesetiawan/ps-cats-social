package helper

import "strings"

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
