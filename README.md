# bubbletea-app-template

A template repository to create `[Bubbletea][bubbletea]` apps. using golang 1.24

```sh
go get -tool github.com/golangci/golangci-lint/cmd/golangci-lint
go get -tool golang.org/x/tools/cmd/goimports@latest
go get -tool github.com/goreleaser/goreleaser/v2@latest
go get -tool mvdan.cc/gofumpt@latest
```

# run lint

````sh
go tool golangci-lint run


# run release
```sh
go tool goreleaser --snapshot --clean
```
# tidy imports
```sh 
go tool goimports -l -w .
```
````
# fmt
```sh
gofumpt -l -w .
```

## Included

```azure
- a sample app that does nothing, so it includes all dependencies:
	- [bubbletea][]
	- [bubbles][]
	- [lipgloss][]
```

- github actions workflows for build, test, lint and release
- [GoReleaser][goreleaser] configs
- [golangci-lint][lint] configs

[bubbletea]: https://github.com/charmbracelet/bubbletea
[bubbles]: https://github.com/charmbracelet/bubbles
[lipgloss]: https://github.com/charmbracelet/lipgloss
[goreleaser]: https://goreleaser.com
[lint]: https://golangci-lint.run
