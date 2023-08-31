package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
)

type GitAccount struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func RegisterNewAccount() {
	var username string
	var email string

	fmt.Println("username:")
	fmt.Scan(&username)
	fmt.Println("email:")
	fmt.Scan(&email)

	var newAccount = GitAccount{
		Uuid:     uuid.New().String(),
		Username: username,
		Email:    email,
	}

	file := openJson(GIT_ACCOUNT_PATH)
	var accounts []GitAccount
	decodeJson(file, &accounts)

	accounts = append(accounts, newAccount)

	encoded := encodeToJson(accounts)
	saveJson(encoded, GIT_ACCOUNT_PATH)
}

func ListGitAccounts() {
	file := openJson(GIT_ACCOUNT_PATH)
	var accounts []GitAccount
	decodeJson(file, &accounts)

	t := getDefaultTable()
	t.AppendHeader(table.Row{"Uuid", "Usename", "Email"})
	for _, account := range accounts {
		t.AppendRow([]interface{}{account.Uuid, account.Username, account.Email})
	}
	t.Render()
}

func ListRepositories() {
	file := openJson(REPOSITORIES)
	var repositories []Repository
	decodeJson(file, &repositories)

	t := getDefaultTable()
	t.AppendHeader(table.Row{"Uuid", "Account Uuid", "path"})
	for _, repository := range repositories {
		t.AppendRow([]interface{}{repository.Uuid, repository.AccountId, repository.RepositoryPath})
	}
	t.Render()
}
