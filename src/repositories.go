package main

import (
	"github.com/google/uuid"
	"path/filepath"
)

type Repository struct {
	Uuid           string `json:"uuid"`
	AccountId      string `json:"AccountId"`
	RepositoryPath string `json:"repositoryPath"`
}

func RegisterNewRepository(path string, accountId string) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	var newRepository = Repository{
		Uuid:           uuid.New().String(),
		AccountId:      accountId,
		RepositoryPath: fullPath,
	}

	file := openJson(REPOSITORIES)
	var repositories []Repository
	decodeJson(file, &repositories)

	repositories = append(repositories, newRepository)

	encoded := encodeToJson(repositories)
	saveJson(encoded, REPOSITORIES)
}
