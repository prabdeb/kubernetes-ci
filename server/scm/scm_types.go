package scm

import (
	"github.com/prabdeb/kubernetes-cicd/secrets"
)

type (
	// New SCM type
	New struct {
		Repository       *Repository
		PipelineRmoteURL string
		SourceBranch     string
		TargetBranch     string
		TargetRef        string
		TriggerType      string
		IsPullRequest    bool
	}

	// Repository type
	Repository struct {
		RepositoryURL string
		Credential    *secrets.Credential
	}
)
