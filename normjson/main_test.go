package normjson

import (
	"encoding/json"
	"testing"
)

var testCases = []struct {
	in   T
	want []byte
}{
	{
		in:   T{Label: "", Active: false, Number: 0.000},
		want: []byte(`{}`),
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {

		normalize(&tc.in)
		got, err := json.Marshal(&tc.in)
		if err != nil {
			t.Error(err)
		}
		if string(got) != string(tc.want) {
			t.Errorf("got %s; want %s", got, tc.want)
		}
	}
}
