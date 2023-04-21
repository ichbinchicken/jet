package helper

import "fmt"

func AnyToString(any interface{}) (string, error) {
	anyString, ok := any.(string)
	if !ok {
		return "", fmt.Errorf("failure on casting to string from any type")
	}
	return anyString, nil
}
