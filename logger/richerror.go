package logger

import "fmt"

type RichError struct {
	Parent     error
	MethodName string
	Message    string
	MetaData   map[string]string
}

func (re RichError) Error() string {
	return fmt.Sprintf("methodName: %s, parent error: %v, message: %s, meta data: %v", re.MethodName, re.Parent, re.Message, re.MetaData)
}
