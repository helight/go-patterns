package chain

import (
	"testing"
)

func TestChain(t *testing.T) {
	job1 := new(Todolist)
	job1.name = "job1"
	job1.info = "job1 info"
	job2 := new(Todolist)
	job2.name = "job2"
	job2.info = "job2 info"
	job3 := new(Todolist)
	job3.name = "job3"
	job3.info = "job3 info"
	job1.addjob(job2)
	job1.addjob(job3)

	job1.dojob("job2")
	job1.dojob("job3")
}

func init() {
}
