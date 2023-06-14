# bulk-clone
## Prepare
Before using this script, you need to install [gh](https://cli.github.com/manual/installation) and [configure](https://docs.github.com/en/github-cli/github-cli/quickstart) it with your GitHub account.

## Usage
Write down the **repository name**, **organization name** or anyone's **username** you want to clone in `./input/repos`.

Such as:
```
gin-gonic/gin
qor5
go-gorm
```

For the above example, will clone `gin-gonic/gin` and all repositories under `qor5` and `go-gorm` organization. You must make sure you have access to these repositories before cloning.

Then run `main.go`. All repositories will be cloned to `./output` directory.
