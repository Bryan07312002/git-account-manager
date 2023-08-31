package main

import (
	"flag"
	"os/exec"
	"path/filepath"
)

const GIT_ACCOUNT_PATH = "./git_accounts.json"
const REPOSITORIES = "./repositories.json"

func setGitConfig(path string) {
	file := openJson(REPOSITORIES)

	var repositories []Repository
	decodeJson(file, &repositories)

	fullPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	var accountId string = ""
	for _, repository := range repositories {
		if repository.RepositoryPath == fullPath {
			accountId = repository.AccountId
			break
		}
	}

	defer file.Close()
	if accountId == "" {
		panic("this path was not regitred")
	}

	file = openJson(GIT_ACCOUNT_PATH)
	var accounts []GitAccount
	decodeJson(file, &accounts)
	for _, account := range accounts {
		if account.Uuid == accountId {
			err := exec.Command("git", "config", "--global", "user.name", account.Username).Run()
			if err != nil {
				panic(err)
			}

			err = exec.Command("git", "config", "--global", "user.email", account.Email).Run()
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	register := flag.Bool("register", false, "register new resources")
	list := flag.Bool("list", false, "list resources")
	set := flag.Bool("set", false, "set account based in local repository")

	account := flag.Bool("account", false, "register new account")
	repository := flag.Bool("repository", false, "register new repository")

	path := flag.String("path", ".", "given path")
	accId := flag.String("accId", "", "account id")

	flag.Parse()

	if *register {
		if *account {
			// To register new user account
			RegisterNewAccount()
		}

		if *repository {
			if *path == "" {
				panic("path: should be given")
			}

			if *accId == "" {
				panic("accId: should be given")
			}

			RegisterNewRepository(*path, *accId)
		}
	}

	if *list {
		if *account {
			ListGitAccounts()
		}

		if *repository {
			ListRepositories()
		}
	}

	if *set {
		setGitConfig(*path)
	}
}
