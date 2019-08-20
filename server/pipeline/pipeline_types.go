package pipeline

import "github.com/prabdeb/kubernetes-cicd/server/scm"

type (
	// New is the is the main type
	New struct {
		Pipeline struct {
			PersistantEnvVariables []string    `yaml:"environment"`
			PersistantPrivileged   bool        `yaml:"privileged"`
			Stages                 []Stage     `yaml:"stages"`
			Conditions             []Condition `yaml:"conditions"`
		} `yaml:"pipeline"`
	}
	// Stage is the pipeline stage type
	Stage struct {
		Name                         string      `yaml:"name"`
		Image                        string      `yaml:"image"`
		Commands                     []string    `yaml:"commands"`
		Entrypoints                  []string    `yaml:"entrypoints"`
		EnvVariables                 []string    `yaml:"environment"`
		Privileged                   bool        `yaml:"privileged"`
		Secrets                      []string    `yaml:"secrets"`
		ManualOrdering               int         `yaml:"order"`
		ParallelGroup                string      `yaml:"parallel"`
		Conditions                   []Condition `yaml:"conditions"`
		ConditionsType               string      `yaml:"conditions-type"`
		ServersideCondition          string
		ServersideConditionParameter string
	}
	// Condition type
	Condition struct {
		Type   string  `yaml:"type"`
		Check  string  `yaml:"check"`
		Stages []Stage `yaml:"stages"`
	}
	// Trigger type
	Trigger struct {
		Type         string
		Argument     string
		EnvVariables []string
		SCM          *scm.New
	}
)
