package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type stepStatus int

const (
	Pending stepStatus = iota
	Ready
	Running
	Completed
)

func (ss stepStatus) String() string {
	switch (ss) {
	case Pending:
		return "Pending"
	case Ready:
		return "Ready"
	case Running:
		return "Running"
	case Completed:
		return "Completed"
	default:
		return "***Unknown***"
	}
}

type step struct {
	name              string
	preRequisiteSteps []*step
	status            stepStatus
}

func NewStep(name string) *step {
	return &step{name, make([]*step, 0, 26), Ready}
}

func (s *step) addPreRequisite(stp *step) {
	s.status = Pending
	s.preRequisiteSteps = append(s.preRequisiteSteps, stp)
}

func (s *step) removePreRequisite(stp *step) {
	idx := findStep(s.preRequisiteSteps, stp)
	if idx != -1 {
		s.preRequisiteSteps = append(s.preRequisiteSteps[:idx], s.preRequisiteSteps[idx+1:]...)
		if len(s.preRequisiteSteps) == 0 {
			s.status = Ready
		}
	}
}

func (s *step) start() error {
	if s.status == Ready {
		s.status = Running
		return nil
	} else {
		return fmt.Errorf("invalid state for start: %s", s.status.String())
	}
}

func (s *step) completed() error {
	if s.status == Running {
		s.status = Completed
		return nil
	} else {
		return fmt.Errorf("invalid state for completed: %s", s.status.String())
	}
}

type worker struct {
	currentStep *step
	timeLeft    int
}

func NewWorker() *worker {
	return &worker{nil, -1}
}
func (w *worker) stepName() string {
	if w.currentStep != nil {
		return w.currentStep.name
	}
	return "."
}

func (w *worker) executeStep(stp *step) {
	w.currentStep = stp
	w.currentStep.start()
	w.timeLeft = 60 + int(stp.name[0]-'A'+1)
}

func (w *worker) tick() bool {
	if w.currentStep != nil {
		w.timeLeft -= 1
		if w.timeLeft == 0 {
			w.currentStep.completed()
		}
	}
	return w.timeLeft == 0
}

func (w *worker) complete() {
	w.currentStep = nil
	w.timeLeft = -1
}

func (w *worker) isBusy() bool {
	return w.currentStep != nil && w.timeLeft >= 0
}

type steps struct {
	steps     map[string]*step
	completed []*step
	workers   []*worker
}

func NewSteps(workers int) *steps {
	stps := &steps{make(map[string]*step), make([]*step, 0, 26), make([]*worker, workers)}
	for w := 0; w < workers; w++ {
		stps.workers[w] = NewWorker()
	}
	return stps
}

func (stps *steps) newOrExistingStep(stepName string) *step {
	stp, ok := stps.steps[stepName]
	if !ok {
		stp = NewStep(stepName)
		stps.steps[stepName] = stp
	}
	return stp
}

func (stps *steps) available() []*step {
	available := make([]*step, 0, len(stps.steps))
	for _, stp := range stps.steps {
		if stp.status == Ready {
			available = append(available, stp)
		}
	}
	sort.SliceStable(available, func(s1, s2 int) bool {
		return available[s1].name < available[s2].name
	})
	return available
}

func containsStep(stps []*step, stp *step) bool {
	return findStep(stps, stp) != -1
}

func findStep(stps []*step, stp *step) int {
	for idx, s := range stps {

		if s.name == stp.name {
			return idx
		}
	}
	return -1
}

func (stps *steps) completeStep(stp *step) error {
	if stp.status != Completed {
		return fmt.Errorf("expected step to be Completed was %s", stp.status.String())
	}
	stps.completed = append(stps.completed, stp)
	for _, stp2 := range stps.steps {
		if stp2.status != Pending {
			continue
		}
		stp2.removePreRequisite(stp)
	}
	return nil
}

func (stps *steps) stepsCompletedAsString() string {
	completed := ""
	for _, s := range stps.completed {
		completed = completed + s.name
	}
	return completed
}

func (stps *steps) execute() {
	second := 0
	fmt.Print("Second\t")
	for idx := range stps.workers {
		fmt.Printf("W%d\t", idx+1)
	}
	fmt.Println("Done\tCompleted")
	for {
		fmt.Printf("%-4d\t", second)
		busyWorkers := 0
		available := stps.available()
		for _, worker := range stps.workers {
			if worker.isBusy() {
				fmt.Printf("%s\t", worker.stepName())
				busyWorkers += 1
				if worker.tick() {
					stps.completeStep(worker.currentStep)
					worker.complete()
				}
			} else if len(available) > 0 {
				nextStep := available[0]
				available = available[1:]
				worker.executeStep(nextStep)
				worker.tick()
				busyWorkers += 1
				fmt.Printf("%s\t", worker.stepName())
			} else {
				fmt.Printf("%s\t", worker.stepName())
			}
		}
		completeString := stps.stepsCompletedAsString()
		fmt.Print("\t", completeString, "\n")
		if busyWorkers == 0 {
			break
		}
		second += 1
	}
}

func main() {
	stepsMap := NewSteps(5)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var stepName, nextStepName string
		_, err := fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.",
			&stepName, &nextStepName)
		if err != nil {
			fmt.Printf("ERROR: an error occurred while reading input (%s): %s", scanner.Text(), err.Error())
			os.Exit(-1)
		}
		//fmt.Printf("%s depends on %s\n", nextStepName, stepName)
		stp := stepsMap.newOrExistingStep(stepName)
		nstp := stepsMap.newOrExistingStep(nextStepName)
		nstp.addPreRequisite(stp)
	}
	stepsMap.execute()
	fmt.Print("\n")
}
