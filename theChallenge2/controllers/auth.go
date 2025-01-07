package controllers

import (
	"theChallenge2/utils"

	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

func Auth() project.Project {

	var projects = project.Project{
		Id:          utils.GetId(),
		PrivateKey:  checks.CheckPrivateKey(utils.GetPrivateKey()),
		Environment: checks.CheckEnvironment("sandbox"),
	}

	return projects
}
