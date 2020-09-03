package idgenerator

const firstID = 0

type Counter struct {
	count int
}

func Init() *Counter {
	return &Counter{firstID}
}

func (c *Counter) NextID() int {
	c.count++
	return c.count
}
