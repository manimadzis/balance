package main

import "encoding/json"

type Status int

const (
	one Status = iota
	two
	three
)

var (
	S2S = map[Status]string{
		one:   "1",
		two:   "2",
		three: "3",
	}
)

func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(S2S[s]), nil
}

type Some struct {
	I int    `json:"id"`
	S Status `json:"status"`
}

func main() {
	// var s Status = 1
	s := Some{1, 2}

	bytes, err := json.Marshal(s)
	if err != nil {
		println(err)
	}

	println(string(bytes))
}
