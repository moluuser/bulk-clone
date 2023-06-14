# bulk-clone
## Prepare
Before using this script, you need to install [gh](https://cli.github.com/manual/installation) and [configure](https://docs.github.com/en/github-cli/github-cli/quickstart) it with your GitHub account.

## Usage
Write down the **repositories** you want to clone in `./input/repos`, one repository per line. Such as:
```
gin-gonic/gin
```

Write down the **organization or anyone’s username** you want to clone in `./input/prefixes`. Such as:
```
qor5
go-gorm
```
For the above example, will clone all repositories under `qor5` and `go-gorm` organization. You must make sure you have access to these repositories before cloning.

Then run `main.go`. All repositories will be cloned to `./output` directory.
