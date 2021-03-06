package messages

import (
	"strings"
)

//Message :
type Message struct {
	Name string
	Body Payload
}

//NewMessage :
func NewMessage(msg string) (*Message, error) {

	n, r, er := parseSocketMessage(msg)
	res := &Message{
		Name: n,
		Body: r,
	}
	return res, er
}

// NewEmptyMessage :
func NewEmptyMessage() *Message {
	return &Message{}
}

//Compile :
func (m *Message) Compile() string {
	pstr := m.Body.String()
	return m.Name + ";" + pstr
}

//Payload :
type Payload map[string]string

//NewPayload :
func NewPayload() Payload {
	return make(Payload)
}

func (p Payload) String() string {
	res := ""
	for k, v := range p {
		res += k + v
	}
	return res
}

//Compile :
func (p Payload) Compile() string {
	res := ""
	for k, v := range p {
		res += k + ":" + v + ";"
	}
	return res
}

//ParseSocketPayload :
func parseSocketMessage(payload string) (string, Payload, error) {
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
	return name, result, nil
}
