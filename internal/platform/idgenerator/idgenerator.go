package idgenerator

var firstId = 0

type Counter struct {
	count int
}

func Init() *Counter {
	return &Counter{firstId}
}

func (c *Counter) NextId() int {
	c.count++
	return c.count
}
