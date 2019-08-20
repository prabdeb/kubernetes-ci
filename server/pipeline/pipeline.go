package pipeline

import (
	"github.com/ghodss/yaml"
	"github.com/prabdeb/kubernetes-cicd/server/scm"
)

// NewPipeline func
func NewPipeline(scmObj *scm.New, triggerObj *Trigger) ([]Stage, error) {
	var stages []Stage
	var err error
	pipelineYaml, err := scm.GetPipelineYml(scmObj)
	newPipelineObj, err := _parsePipeline(pipelineYaml)
	stages, err = _createStages(newPipelineObj, triggerObj)
	return stages, err
}

// _parsePipeline func
func _parsePipeline(pipelineYaml string) (*New, error) {
	newPipelineObj := New{}
	err := yaml.Unmarshal([]byte(pipelineYaml), &newPipelineObj)
	return &newPipelineObj, err
}

// _createStages func
func _createStages(newPipelineObj *New, trigger *Trigger) ([]Stage, error) {
	return ParseConditions(newPipelineObj, trigger)
}
