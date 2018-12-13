package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type step struct {
	name              string
	nextSteps         []*step
	preRequisiteSteps []*step
	completed         bool
}

func NewStep(name string) *step {
	return &step{name, make([]*step, 0, 50), make([]*step, 0, 50), false}
}

func (s *step) addPreRequisite(stp *step) {
	s.preRequisiteSteps = append(s.preRequisiteSteps, stp)
}

func (s *step) addNextStep(stp *step) {
	s.nextSteps = append(s.nextSteps, stp)
}

type steps struct {
	steps          map[string]*step
	availableSteps []*step
	completedSteps []*step
}

func NewSteps() *steps {
	return &steps{make(map[string]*step), make([]*step, 0, 101), make([]*step, 0, 101)}
}

func (stps *steps) newOrExistingStep(stepName string) *step {
	stp, ok := stps.steps[stepName]
	if !ok {
		stp = NewStep(stepName)
		stps.steps[stepName] = stp
	}
	return stp
}

func (stps *steps) execute() string {
	for _, stp := range stps.steps {
		if len(stp.preRequisiteSteps) == 0 {
			stps.availableSteps = append(stps.availableSteps, stp)
		}
	}
	return stps.executeNext("")
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

func (stps *steps) executeNext(compleated string) string {
	if len(stps.availableSteps) == 0 {
		return compleated
	}
	sort.SliceStable(stps.availableSteps, func(s1, s2 int) bool {
		return stps.availableSteps[s1].name < stps.availableSteps[s2].name
	})
	nextStep := stps.availableSteps[0]
	compleated = compleated + nextStep.name
	stps.completedSteps = append(stps.completedSteps, nextStep)
	stps.availableSteps = stps.availableSteps[1:]
	for _, nstp := range nextStep.nextSteps {
		if !containsStep(stps.availableSteps, nstp) {
			completedPreReqCount := 0
			for _, cstp := range stps.completedSteps {
				if containsStep(nstp.preRequisiteSteps, cstp) {
					completedPreReqCount += 1
				}
			}
			if completedPreReqCount == len(nstp.preRequisiteSteps) {
				stps.availableSteps = append(stps.availableSteps, nstp)
			}
		}
	}
	return stps.executeNext(compleated)
}

func (stps *steps) stepReady(name string) bool {
	stp := stps.steps[name]
	return !stp.completed && len(stp.preRequisiteSteps) == 0
}

func (stps *steps) completeStep(stp *step) {
	stp.completed = true
	stps.completedSteps = append(stps.completedSteps, stp)
	for _, pstp := range stps.steps {
		idx := findStep(pstp.preRequisiteSteps, stp)
		if idx != -1 {
			pstp.preRequisiteSteps = append(pstp.preRequisiteSteps[:idx], pstp.preRequisiteSteps[idx+1:]...)
		}
	}
}

func (stps *steps) executeWorkers(order string, numWorkers int) int {
	workerTimes := make([]int, numWorkers)
	workerSteps := make([]*step, numWorkers)
	pos := 0
	second := 0
	fmt.Printf("order: %s\n", order)
	for {
		fmt.Printf("%d:\n", second)
		busyWorker := 0
		for i := range workerTimes {
			if pos < len(order) {
				nextStepName := order[pos : pos+1]
				if workerTimes[i] == 0 && stps.stepReady(nextStepName) {
					fmt.Printf("Starting %s on worker #%d\n", nextStepName, i)
					workerTimes[i] = 60 + int(rune(order[pos])-rune('A')+1)
					workerSteps[i] = stps.steps[nextStepName]
					pos += 1

				}
			}
			if workerTimes[i] > 0 {
				workerTimes[i] -= 1
				if workerTimes[i] == 0 {
					stps.completeStep(workerSteps[i])
				}
				busyWorker += 1
			}
		}
		second += 1
		if busyWorker == 0 {
			break
		}
	}
	return second
}

func main() {
	stepsMap := NewSteps()
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

		stp.addNextStep(nstp)
		nstp.addPreRequisite(stp)
	}

	fmt.Println(stepsMap.executeWorkers(stepsMap.execute(), 5))
}
