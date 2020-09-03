package queue

import (
	"testing"
)

// Queue test
func TestQueue(t *testing.T) {
	var queue *Queue = New()
	queue.Enqueue(32, "QUEUED", "TIME_CRITICAL")
	queue.Enqueue(40, "QUEUED", "TIME_CRITICAL")
	queue.Enqueue(100, "QUEUED", "TIME_CRITICAL")

	peek := queue.Peek()
	if peek.Item() != 32 {
		t.Errorf("got %v expected %v \n", peek.Item(), 32)
	}

	if peek.Status() != "QUEUED" {
		t.Errorf("got %v expected %v \n", peek.Status(), "QUEUED")
	}
}

func TestDequeue(t *testing.T) {
	var queue *Queue = New()
	queue.Enqueue(32, "QUEUED", "TIME_CRITICAL")
	queue.Enqueue(40, "QUEUED", "TIME_CRITICAL")
	queue.Enqueue(100, "QUEUED", "TIME_CRITICAL")
	queue.Enqueue(2, "QUEUED", "TIME_CRITICAL")

	queue.Dequeue()

	peek := queue.Peek()
	if peek.Item() != 40 {
		t.Errorf("wrong dequeue")
	}

	queue.Dequeue()

	peek = queue.Peek()
	if peek.Item() != 100 {
		t.Errorf("wrong dequeue")
	}

	if peek.Status() != "QUEUED" {
		t.Errorf("got %v expected %v \n", peek.Status(), "QUEUED")
	}
}
