package sub

import "github.com/open-okapi/data-driven-proxies/enums"

type Pairs struct {
	Key        string
	Value      string
	ValueType  enums.PairType
	Encryption string
	Secret     string
	Doc        string
}
