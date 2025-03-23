# bubbletea-app-template

A template repository to create `[Bubbletea][bubbletea]` apps. using golang 1.24

```sh
go get -tool github.com/golangci/golangci-lint/cmd/golangci-lint
```

# run lint

````sh
go tool golangci-lint
```

# run release
````

go tool goreleaser --snapshot --clean

````
## Included

```azure
- a sample app that does nothing, so it includes all dependencies:
	- [bubbletea][]
	- [bubbles][]
	- [lipgloss][]
````

- github actions workflows for build, test, lint and release
- [GoReleaser][goreleaser] configs
- [golangci-lint][lint] configs

[bubbletea]: https://github.com/charmbracelet/bubbletea
[bubbles]: https://github.com/charmbracelet/bubbles
[lipgloss]: https://github.com/charmbracelet/lipgloss
[goreleaser]: https://goreleaser.com
[lint]: https://golangci-lint.run
