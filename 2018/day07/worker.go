package day07

type Worker struct {
	delay     int
	working   string
	remaining int
	history   string
}

func NewWorker(delay int) *Worker {
	return &Worker{delay: delay}
}

func (this *Worker) IsIdle() bool {
	return len(this.working) == 0
}

func (this *Worker) Accept(task string) {
	this.remaining = this.delay + int(task[0]-'A') + 1
	this.working = task
}

func (this *Worker) DoWork() (finished string) {
	this.history += this.working
	if len(this.working) == 0 {
		this.history += " "
	}

	if len(this.working) > 0 {
		if this.remaining--; this.remaining == 0 {
			finished = this.working
			this.working = ""
		}
	}
	return finished
}
