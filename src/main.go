package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"path/filepath"
)

const GIT_ACCOUNT_PATH = "./git_accounts.json"
const REPOSITORIES = "./repositories.json"

type GitAccount struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Repository struct {
	Uuid           string `json:"uuid"`
	AccountId      string `json:"AccountId"`
	RepositoryPath string `json:"repositoryPath"`
}

func openJson(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return file
}

func saveJson(encoded []byte, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write(encoded)
}

func decodeJson[T any](file *os.File, accounts *[]T) {
	err := json.NewDecoder(file).Decode(&accounts)
	if err != nil {
		panic(err)
	}
}

func encodeToJson[T any](accounts []T) []byte {
	jsonData, err := json.Marshal(accounts)
	if err != nil {
		panic(err)
	}

	return jsonData
}

func RegisterNewAccount() {
	var username string
	var email string

	fmt.Println("username: ")
	fmt.Scan(&username)
	fmt.Println("email: ")
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

func getDefaultTable() table.Writer {
	t := table.NewWriter()

	t.SetStyle(table.StyleLight)
	t.Style().Color.Header = text.Colors{text.FgGreen}
	t.Style().Color.Row = text.Colors{text.FgWhite}
	t.Style().Color.Separator = text.Colors{text.FgYellow}
	t.Style().Format.Footer = text.FormatLower
	t.Style().Options.DrawBorder = false

	t.SetOutputMirror(os.Stdout)
	return t
}

func main() {
	// To register new user
	register := flag.Bool("register", false, "register new resources")
	list := flag.Bool("list", false, "list resources")
	// set := flag.Bool("set", false, "set account based in local repository")

	account := flag.Bool("account", false, "register new account")
	repository := flag.Bool("repository", false, "register new repository")

	path := flag.String("path", "", "given path")
	accId := flag.String("accId", "", "account id")

	flag.Parse()

	if *register {
		if *account {
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
}
