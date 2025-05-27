package core

import (
	"errors"
	"fmt"
)

// TransferStep applies the transferFunction recursively to a FlowStep and its nested steps.
func TransferStep(step FlowStep, transferFunction func(step FlowStep) FlowStep) FlowStep {
	// Apply the transfer function to the current step
	updatedStep := transferFunction(step)

	switch updatedStep.Type {
	case FlowStepTypeLoop:
		if updatedStep.FirstLoopStep != nil {
			s := TransferStep(*updatedStep.FirstLoopStep, transferFunction)
			updatedStep.FirstLoopStep = &s
		}
	case FlowStepTypeStepRouter:
		if updatedStep.Children != nil {
			for i, child := range updatedStep.Children {
				if child != nil {
					s := TransferStep(*child, transferFunction)
					updatedStep.Children[i] = &s
				}
			}
		}
	}

	if updatedStep.NextStep != nil {
		s := TransferStep(*updatedStep.NextStep, transferFunction)
		updatedStep.NextStep = &s
	}

	return updatedStep
}

// GetAllSteps Utility function to get all steps in a FlowStep recursively
func GetAllSteps(step FlowStep) []FlowStep {
	var steps []FlowStep

	TransferStep(step, func(currentStep FlowStep) FlowStep {
		steps = append(steps, currentStep)
		return currentStep
	})
	return steps
}

// ErrorStepNotFound definition for a case where a step is not found
var ErrorStepNotFound = errors.New("step not found")

// GetStep retrieves a step by name in the flow hierarchy
func GetStep(name string, flowRoot FlowStep) *FlowStep {
	allSteps := GetAllSteps(flowRoot)
	for _, step := range allSteps {
		if step.Name == name {
			return &step
		}
	}
	return nil
}

// GetStepOrThrow retrieves a step by name or throws an error if not found
func GetStepOrThrow(name string, flowRoot FlowStep) (*FlowStep, error) {
	step := GetStep(name, flowRoot)
	if step == nil {
		return nil, fmt.Errorf("%w: stepName=%s", ErrorStepNotFound, name)
	}
	return step, nil
}
