package generate

import (
	"encoding/json"
	"math/rand"
	"time"
)

// Result is service result.
type Result []byte

func newResult(id int64, max int64, rnd *rand.Rand) (Result, error) {
	type result struct {
		Worker    int64 `json:"worker,string"`
		Timestamp int64 `json:"timestamp,string"`
		Number    int64 `json:"number,string"`
	}
	r := result{
		Worker:    id,
		Timestamp: time.Now().UTC().UnixNano(),
		Number:    rnd.Int63n(max),
	}
	raw, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return Result(raw), nil
}
