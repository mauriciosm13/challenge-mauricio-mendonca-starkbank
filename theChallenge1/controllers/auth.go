package controllers

import (
	"challenge-mauricio-mendonca-starkbank/theChallenge1/utils"

	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

func Auth() project.Project {
	//funcao para realizar a conexao com o projeto no ambiente sandbox

	var projects = project.Project{
		Id:          utils.GetId(),
		PrivateKey:  checks.CheckPrivateKey(utils.GetPrivateKey()),
		Environment: checks.CheckEnvironment("sandbox"),
	}

	return projects
}
