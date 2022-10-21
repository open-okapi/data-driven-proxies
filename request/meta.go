package request

import "github.com/open-okapi/data-driven-proxies/sub"

type Meta struct {
	UUID   string
	Http   sub.Http
	Entity sub.Entity
	Header []sub.Header
	Pairs  []sub.Pairs
	Query  []sub.Query
}
