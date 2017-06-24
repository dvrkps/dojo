package normjson

// T holds data.
type T struct {
	Label  interface{} `json:"label,omitempty"`
	Active interface{} `json:"active,omitempty"`
	Number interface{} `json:"number,omitempty"`
}

func normalize(t *T) error {
	fields := []interface{}{
		t.Label,
		t.State,
		t.Number,
	}

	for _, v := range fields {
		switch v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			if v == 0 {
				*v = nil
			}
		case float32, float64:
			if *v == 0 {
				*v = nil
			}
		case string:
			if *v == "" {
				*v = nil
			}
		}

		//if *v == 0 || *v == "0" || *v == "false" || *v == false || *v == "" {
		//	*v = nil
		//}
	}

	return nil
}
