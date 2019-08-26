# goCleanse

  

This project's goal is to help users cleanse their repositories of sensitive strings such as credentials, tokens or pretty much any thing you need gone.

  

### Note:

If you have a repository with credentials or other sensitive information, the first thing you should do is **change them**!
 

___

## Getting Started
These instructions will help you get started on removing credentials from your git repository.

### Installing
```sh
go get github.com/Krlier/goCleanse
```

### Running
#### Note : In order to completely remove strings from git history we need to rewrite git history, which could possibly break things. Be sure to have a backup of your repository before proceeding!
`goCleanse` reads all the information it needs from environment variables. You can create an `.env` file with all of them just like the example below:
```sh
export GCLEANSE_REPURL="git@github.com:Example/myExampleRepo.git"
export GCLEANSE_CLONEPATH="./tmp"
export GCLEANSE_NEWBRANCHNAME="gCleanseBranch"
export GCLEANSE_OLDSTRING="superSecretPassword"
export GCLEANSE_NEWSTRING="removedPassword"
export GCLEANSE_FILEEXTENSION="txt"
export GCLEANSE_REMOTE="origin"
```
After that, simply source the file through the command:
```sh
. .env
```
Now that we have our environment variables properly set up, it's time to run `goCleanse` through the commands:
```sh
cd app/
go run app.go
```
If everything went ok, you should now have a brand new branch in your repository without the string you removed! If you want to, you can rename this branch and make it your new `master`.

### Signing Commits
If your commits were GPG signed, after running `goCleanse` they will come out unsigned. In order to re-sign them, you can simply run `goCleanse` again, but this time providing the e-mail address of the user who would like to have his commits signed.
```sh
go run app.go daniel@example.com
```

After running the command, you should now see all your commits have been GPG signed. 😁