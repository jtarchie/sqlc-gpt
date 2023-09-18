package parser

import "fmt"

// ReturnedValue represents the Ruby ReturnedValue struct.
type ReturnedValue struct {
	Statement string
}

// NewReturnedValue is a constructor function for ReturnedValue.
func NewReturnedValue(statement string) *ReturnedValue {
	return &ReturnedValue{Statement: statement}
}

// ArgName always returns the string "name" for ReturnedValue.
func (rv *ReturnedValue) ArgName() string {
	return fmt.Sprintf("%sRet", rv.Statement)
}

// ArgType returns the argument type based on the arg name.
func (rv *ReturnedValue) ArgType() string {
	switch rv.Statement {
	case "name", "value":
		return "string"
	case "names":
		return "[]string"
	default:
		return "int64"
	}
}

func (rv *ReturnedValue) DefaultValue() string {
	switch rv.Statement {
	case "name", "value":
		return `""`
	case "names":
		return "nil"
	default:
		return "0"
	}
}
