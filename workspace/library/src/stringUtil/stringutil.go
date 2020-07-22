package stringUtil

import "strings"

func FullName(f,l string) string{
	full_name:= strings.ToUpper(f+" "+l)
	return full_name
}
