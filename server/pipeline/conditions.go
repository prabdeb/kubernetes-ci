package pipeline

import (
	"fmt"
	"strings"

	"github.com/prabdeb/go-utils/contains"
	"github.com/ryanuber/go-glob"
)

// List of supported_conditions
var supportedConditions = []string{}

// ParseConditions and expand stages accordingly
func ParseConditions(pipelineObj *New, pipelineTrigger *Trigger) ([]Stage, error) {
	stages := []Stage{}
	var err error
	// Priority 1: check for stages having explicit condition
	if pipelineObj.Pipeline.Stages != nil && len(pipelineObj.Pipeline.Stages) > 0 {
		for _, stage := range pipelineObj.Pipeline.Stages {
			stageIsAdded := false
			if stage.Conditions != nil && len(stage.Conditions) > 0 {
				for _, condition := range stage.Conditions {
					var shallExecute bool
					var serversideCondition string
					var serversideConditionParameter string
					shallExecute, serversideCondition, serversideConditionParameter, err = _evaluateCondition(condition, pipelineTrigger)
					if !stageIsAdded && shallExecute {
						if serversideCondition != "" && serversideConditionParameter != "" {
							stage.ServersideCondition = serversideCondition
							stage.ServersideConditionParameter = serversideConditionParameter
						}
						stages = append(stages, stage)
						stageIsAdded = true
					}
					// Having multiple condition along with AND strict check
					if stageIsAdded && !shallExecute {
						if strings.ToLower(stage.ConditionsType) == "and" {
							if contains.Contains(stages, stage) {
								stages = stages[:len(stages)-1]
							}
						}
					}
				}
			}
		}
	}
	// Priority 2: check for stages inside condition blocks
	// any conditions defined inside this stages, will be ignored
	if pipelineObj.Pipeline.Conditions != nil && len(pipelineObj.Pipeline.Conditions) > 0 {
		for _, condition := range pipelineObj.Pipeline.Conditions {
			var shallExecute bool
			var serversideCondition string
			var serversideConditionParameter string
			shallExecute, serversideCondition, serversideConditionParameter, err = _evaluateCondition(condition, pipelineTrigger)
			if shallExecute {
				for _, stage := range condition.Stages {
					if serversideCondition != "" && serversideConditionParameter != "" {
						stage.ServersideCondition = serversideCondition
						stage.ServersideConditionParameter = serversideConditionParameter
					}
					stages = append(stages, stage)
				}
			}
		}
	}
	return stages, err
}

func _evaluateCondition(pipelineConditionObj Condition, pipelineTrigger *Trigger) (bool, string, string, error) {
	shallExecute := false
	var serversideCondition string
	var serversideConditionParameter string
	var err error
	if strings.ToLower(pipelineConditionObj.Type) == "pull-request" || strings.ToLower(pipelineConditionObj.Type) == "pullrequest" {
		// Check if pull request
		if pipelineTrigger.SCM.IsPullRequest {
			shallExecute = true
		}
	} else if strings.ToLower(pipelineConditionObj.Type) == "git-branch" || strings.ToLower(pipelineConditionObj.Type) == "gitbranch" {
		// Check if expected git branch
		if glob.Glob(pipelineConditionObj.Check, pipelineTrigger.SCM.TargetBranch) {
			shallExecute = true
		}
	} else if strings.ToLower(pipelineConditionObj.Type) == "git-ref" || strings.ToLower(pipelineConditionObj.Type) == "gitref" {
		// Check if expected git ref
		if glob.Glob(pipelineConditionObj.Check, pipelineTrigger.SCM.TargetRef) {
			shallExecute = true
		}
	} else if strings.ToLower(pipelineConditionObj.Type) == "on-demand" || strings.ToLower(pipelineConditionObj.Type) == "ondemand" {
		// Check if	custom stages needs to be called
		if strings.ToLower(pipelineConditionObj.Check) == strings.ToLower(pipelineTrigger.Argument) {
			shallExecute = true
		}
	} else if strings.ToLower(pipelineConditionObj.Type) == "file-exists" || strings.ToLower(pipelineConditionObj.Type) == "fileexists" {
		// Check if a specific file exists in workspace
		shallExecute = true
		serversideCondition = "fileexists"
		serversideConditionParameter = pipelineConditionObj.Check
	} else if strings.ToLower(pipelineConditionObj.Type) == "execute-local" || strings.ToLower(pipelineConditionObj.Type) == "executelocal" {
		// Check by custom executable scripts return status in workspace
		shallExecute = true
		serversideCondition = "executelocal"
		serversideConditionParameter = pipelineConditionObj.Check
	} else if strings.ToLower(pipelineConditionObj.Type) == "status" {
		// Check by status
		shallExecute = true
		serversideCondition = "status"
		serversideConditionParameter = pipelineConditionObj.Check
	} else {
		// Not supported condition error
		err = fmt.Errorf("NotSupportedCondition: %s is not supported, check for all supported conditions", pipelineConditionObj.Type)
	}
	return shallExecute, serversideCondition, serversideConditionParameter, err
}
