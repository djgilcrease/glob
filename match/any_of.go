package match

import (
	"fmt"
)

type AnyOf struct {
	Matchers Matchers
}

func (self *AnyOf) Add(m Matcher) error {
	self.Matchers = append(self.Matchers, m)
	return nil
}

func (self AnyOf) Match(s string) bool {
	for _, m := range self.Matchers {
		if m.Match(s) {
			return true
		}
	}

	return false
}

//func (self AnyOf) Index(s string) (index int, segments []int) {
//
//}

func (self AnyOf) Len() (l int) {
	l = -1
	for _, m := range self.Matchers {
		ml := m.Len()
		if ml == -1 {
			return -1
		}

		if l == -1 {
			l = ml
			continue
		}

		if l != ml {
			return -1
		}
	}

	return
}

func (self AnyOf) Kind() Kind {
	return KindAnyOf
}

func (self AnyOf) String() string {
	return fmt.Sprintf("[any_of:%s]", self.Matchers)
}
