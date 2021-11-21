package merror

import (
	"fmt"
)

type Merror struct {
	timeStamp string
	identity  string
	msg       string
}

func (m *Merror) Error() string {
	return fmt.Sprintf("%s: %s %s", m.timeStamp, m.identity, m.msg)
}

func E(t string, identity string, m string) *Merror {
	return &Merror{t, identity, m} //heap allocation //todo move to stack!!!!
}
