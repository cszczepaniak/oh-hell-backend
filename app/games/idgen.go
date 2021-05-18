package games

import "time"

type IdGenerator interface {
	NextId() int64
}

type TimeStampIdGenerator struct{}

func (ts TimeStampIdGenerator) NextId() int64 {
	return time.Now().UnixNano()
}
