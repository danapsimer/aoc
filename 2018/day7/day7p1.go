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
}

func NewStep(name string) *step {
	return &step{name, make([]*step, 0, 50), make([]*step, 0, 50)}
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

func (stps *steps) execute() {
	for _, stp := range stps.steps {
		if len(stp.preRequisiteSteps) == 0 {
			stps.availableSteps = append(stps.availableSteps, stp)
		}
	}
	stps.executeNext()
}

func containsStep(stps []*step, stp *step) bool {
	for _, s := range stps {
		if s.name == stp.name {
			return true
		}
	}
	return false
}

func (stps *steps) executeNext() {
	if len(stps.availableSteps) == 0 {
		return
	}
	sort.SliceStable(stps.availableSteps, func(s1, s2 int) bool {
		return stps.availableSteps[s1].name < stps.availableSteps[s2].name
	})
	nextStep := stps.availableSteps[0]
	fmt.Print(nextStep.name)
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
	stps.executeNext()
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
	stepsMap.execute()
	fmt.Print("\n")
}
