package scm

import (
	"io/ioutil"
	"net/http"
)

// NewRepository func
func NewRepository() *Repository {
	return &Repository{}
}

// NewScm func
func NewScm(repository *Repository) *New {
	var newScm New
	newScm.Repository = repository
	return &newScm
}

// GetPipelineYml func
func GetPipelineYml(scmObj *New) (string, error) {
	var pipelineYaml string
	var err error
	client := &http.Client{}
	req, err := http.NewRequest("GET", scmObj.PipelineRmoteURL, nil)
	if err != nil {
		return pipelineYaml, err
	}
	req.SetBasicAuth(scmObj.Repository.Credential.UserName, scmObj.Repository.Credential.Password)
	resp, err := client.Do(req)
	if err != nil {
		return pipelineYaml, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return pipelineYaml, err
	}
	pipelineYaml = string(body)
	return pipelineYaml, err
}
