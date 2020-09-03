package job

import (
	it "jobs/internal/platform/datastructure/iterator"
	"jobs/internal/platform/datastructure/node"
	"jobs/internal/platform/datastructure/queue"
)

const JOBSTATUSQUEUED = "QUEUED"
const JOBSTATUSINPROGRESS = "IN_PROGRESS"
const JOBSTATUSCONCLUDED = "CONCLUDED"

const JOBTYPETIMECRITICAL = "TIME_CRITICAL"
const JOBTYPENOTTIMECRITICAL = "NOT_TIME_CRITICAL"

type Config struct {
	*queue.Queue
}

type Job struct {
	ID     int
	Type   string
	Status string
}

func Init() *Config {
	q := queue.New()
	return &Config{q}
}

func (c Config) FindByID(id int) *node.Node {
	var iter it.IIterator = c.Iterator()
	for iter.HasNext() {
		job := iter.Next()
		if id == job.Item() {
			return job
		}
	}
	return nil
}

func (c Config) GetNextFreeJob() *node.Node {
	var iter it.IIterator = c.Iterator()
	for iter.HasNext() {
		job := iter.Next()
		if job.Status() == JOBSTATUSQUEUED {
			job.SetStatus(JOBSTATUSINPROGRESS)
			return job
		}
	}
	return nil
}

func (c Config) ConcludeJobBiID(id int) *node.Node {
	job := c.FindByID(id)
	if job != nil && job.Status() == JOBSTATUSINPROGRESS {
		job.SetStatus(JOBSTATUSCONCLUDED)
		return job
	}
	return nil
}
