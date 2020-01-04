package parser

import (
	"fmt"
	"strings"
)

//Payload :
type Payload map[string]string

func (p Payload) String() string {
	res := ""
	for k, v := range p {
		res += k + v
	}
	return res
}

//ParseSocketPayload :
func ParseSocketPayload(payload string) (string, Payload, error) {
	sarr := strings.Split(payload, ";")
	name := sarr[0]
	sarr = sarr[1:]
	result := make(map[string]string)
	for _, val := range sarr {
		pair := strings.Split(val, ":")
		if len(pair) != 2 {
			continue
		}
		result[pair[0]] = pair[1]
	}
	fmt.Println("Parsing str from", payload, "to", name, result)
	return name, result, nil
}
