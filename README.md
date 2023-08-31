# Git account manager
A small and quick project just for easily manage git config name and email

## How to use
To register a new account
```sh
# gam --register --account
username:
{account_username}
email:
{account_mail}
```

To list all registred accounts
```sh
# gam --list --account
```

To register a new repository
```sh
# gam --register --repository --accId {account_uuid_given_in_list_comand} --path[this flag is optional] {desired_path}
```

To list all registred repositories
```sh
# gam --list --repository
```

To set a account in git based at current repository
```sh
# gam --set
```

To set a account in git based at choosen repository
```sh
# gam --set --path {desired_repository_path}
```
