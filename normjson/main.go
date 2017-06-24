package normjson

// T holds data.
type T struct {
	Label  interface{} `json:"label,omitempty"`
	Active interface{} `json:"active,omitempty"`
	Number interface{} `json:"number,omitempty"`
}
