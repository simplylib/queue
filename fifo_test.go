package queue

import "testing"

func TestFIFO(t *testing.T) {
	q := FIFO[int]{}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	if q.All()[0] != 1 && q.All()[1] != 2 && q.All()[2] != 3 {
		t.Fatal("all does not contain expected data")
	}

	if s := q.Size(); s != 3 {
		t.Fatalf("size should be 3, have (%v)\n", s)
	}

	if !q.More() {
		t.Fatal("no more in queue after adding")
	}

	if n := q.Next(); n != 1 {
		t.Fatalf("first in queue is not 1, got (%v)\n", n)
	}

	if !q.More() {
		t.Fatal("should be 2 in queue")
	}

	if n := q.Next(); n != 2 {
		t.Fatalf("first in queue is not 2, got (%v)\n", n)
	}

	if !q.More() {
		t.Fatal("should be 1 in queue")
	}

	if n := q.Next(); n != 3 {
		t.Fatalf("first in queue is not 3, got (%v)\n", n)
	}

	if q.More() {
		t.Fatal("should be 0 in queue")
	}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	q.Clear()

	if q.More() {
		t.Fatal("queue has data after being cleared")
	}
}
