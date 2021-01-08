# Section 3 Notes

### Go Workspace
* One folder // any name or location
    * `Bin`
        * Binary will be stored here
    * `Pkg`
        *  Archived files, when you compile code from other packages and those packages havent changed they will be placed here. So they dont need to be recomplied saving time and compliation
    * `Src`
        * Name spacing and package mangagement

### Go Commands
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
