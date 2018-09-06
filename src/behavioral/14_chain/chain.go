package chain

import "fmt"

type Todolist struct {
	name string
	info string
	next *Todolist
}

func (t *Todolist) dojob(jobname string) string {
	if jobname == t.name {
		fmt.Println(t.info);
		return t.info + " do this job"
	} else {
		if t.next == nil {
			return "finish all jobs"
		} else {
			return t.next.dojob(jobname)
		}
	}
}

func (t *Todolist) addjob(job *Todolist) {
	if t.next == nil {
		t.next = job
	} else {
		t.next.addjob(job)
	}
}

func main() {
}
