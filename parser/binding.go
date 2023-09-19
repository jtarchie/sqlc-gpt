package parser

// Binding represents the Ruby Binding struct.
type Binding struct {
	Name        string
	Placeholder string
}

// NewBinding is a constructor function for Binding.
func NewBinding(name, placeholder string) *Binding {
	return &Binding{Name: name, Placeholder: placeholder}
}

// ArgName returns the arg name, similar to the arg_name method in Ruby.
func (b *Binding) ArgName() string {
	return b.Name
}

// ArgType returns the argument type based on the arg name.
func (b *Binding) ArgType() string {
	switch b.ArgName() {
	case "name", "pivot", "value":
		return "string"
	case "names":
		return "[]string"
	case "start", "end", "index", "offset":
		return "int64"
	default:
		return "int64" // You can adjust this to another default if needed.
	}
}
