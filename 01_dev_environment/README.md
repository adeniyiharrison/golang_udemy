# Section 3 Notes

### Go Workspace
* One folder // any name or location
    * `Bin`
        * Binary will be stored here
    * `Pkg`
        *  Archived files, when you compile code from other packages and those packages havent changed they will be placed here. So they dont need to be recomplied saving time and compliation
    * `Src`
        * Name spacing and package mangagement

### Initial Go Commands
* `go fmt`: auto formatting your code
    * `go fmt ./...`: format recursively through a directory
* `go run`: Builds and runs code, similar to `python app.py` or whatever
    * `go run <file_name>.go`
* `go build`
    * For an executable
        * Builds the file. Good for checking that your code can build properly. If there are no errors, it puts an executable into the current folder
        * will create something like `001-hello-world.exe`
        * you can run this executable like `./001-hello-world.exe`
    * For a package
        * Build the file
        * Reports any errors
        * Throws away binary
* `go install`
    * For an executable
        * compiles the program (builds it)
        * names the executable the folder name holding the code
        * puts the executable in `workspace/bin`
            * `$GOPATH/bin`
    * For a package
        * compiles the code
        * puts the executable in `worspace/pkg`
            * `$GOPATH/pkg`
        * majes it an archived file

### Managing Go Modules
* `go mod init`
    * creates new module, initializing the `go.mod` that describes it
* `go build, go test`
    * and other package-building commands adds new dependencies to `go.mod`
* `go list -m all`
    * prints all of the current modules dependencies (both direct and indirect)
    * `go list -m versions <package_name>`
        * list all the available versions of package
* `go get`
    * changes the required version of dependency or adds new dependency
    * `go get <package_name>@<version_number>`
        * specify the specific version you want
* `go mod tidy`
    * removes any unused dependency from module