- `go mod init` creates a new module, initializing the go.mod file that describes it.
- `go build`, `go test`, and other package-building commands add new dependencies to go.mod as needed.
- `go list -m all` prints the current module’s dependencies.
- `go list -m -versions github.com/shall-we-go/somemodule` changes the required version of a dependency (or adds a new dependency).
- `go mod tidy` removes unused dependencies.
- `git tag v0.1.0` creates tag for versioning
- `git push origin v0.1.0` push tag to origin
- `go get -u` upgrade its dependency`