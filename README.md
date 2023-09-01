# Git account manager
A small and quick project just for easily manage git config name and email

## How to use
To register a new account
```sh
$ gam --register --account
username:
{account_username}
email:
{account_mail}
```

To list all registred accounts
```sh
$ gam --list --account
```

To register a new repository
```sh
$ gam --register --repository --accId {account_uuid_given_in_list_comand} --path[this flag is optional] {desired_path}
```

To list all registred repositories
```sh
$ gam --list --repository
```

To set a account in git based at current repository
```sh
$ gam --set
```

To set a account in git based at choosen repository
```sh
$ gam --set --path {desired_repository_path}
```

## How to install 

### Linux
Pre requisites
    - go
    - git
    
Download this repository

```sh
$ git clone https://github.com/Bryan07312002/git-account-manager.git
$ cd git-account-manager
```

Build the binary
```sh
$ go build ./src/*.go
```

Send to /bin
```sh
$ sudo mv accounts /bin/gam
```

create basic json folders and files
```sh
$ mkdir ~/.gam_config
$ touch ~/.gam_config/git_accounts.json
$ touch ~/.gam_config/repositories.json
$ echo [] > ~/.gam_config/git_accounts.json
$ echo [] > ~/.gam_config/repositories.json
```

add this to your .bashrc file
```bash
export GAM=~/.gam_config

function git() {
    command gam --set # Run gam set command to set right account
    command git "$@"  # Run the original git command
}
```
