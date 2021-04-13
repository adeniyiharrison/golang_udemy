## Simple Guide to setup Golang in your Mac and clone a repo.
* https://gist.github.com/sathiyaseelan/529695891e290991573d278a56180535

## Setup Go and workspace
* Create a Go workspace and set GO PATH
    * __GOPATH=/Users/adeniyiharrison/Documents/goworkspace__
    * __GOBIN=/Users/adeniyiharrison/Documents/goworkspace/bin__
        * The directory where 'go install' will install a command.
        * __Folder contains the Go packages you install__
    * __GOROOT=/usr/local/go__
        * The root of the go tree. // not sure what this means

```
mkdir -p $HOME/go
export GOPATH="$HOME/go"

# Folder contains your golang source codes
mkdir -p $GOPATH/src

# Folder contains the binaries when you install an go based executable
mkdir -p $GOPATH/bin

# Folder contains the Go packages you install
mkdir -p $GOPATH/pkg

# Folder contains the Github Source code for the repos you cloned
mkdir -p $GOPATH/src/github.com

```

## Cloning a Repo
* The folder structure to maintain the local repo is _$GOPATH/src/github.com/_
```
mkdir -p $GOPATH/src/github.com/keratin/authn-server
cd $GOPATH/src/github.com/keratin/authn-server

git clone git@github.com:keratin/authn-server.git
```

## Installing Go packages
```
go get github.com/benbjohnson/ego/cmd/ego

go get github.com/airbrake/gobrake
```
* Note: Depends on the package, it may install a binary under $GOPATH/bin or a package in $GOPATH/pkg

### Go Get vs Go Install
* it seems like `go get` is equivalent to pip installing something
* Go install creates a binary I believe

### go run vs go build vs go install
* https://levelup.gitconnected.com/go-run-vs-go-build-vs-go-install-c7c0fd135cf9
```
    // file name : hello.go
    package main
    import "fmt"
    func main(){
    
        fmt.Println("Hello, world!")
    }
```
* `Go Run`
    * If you run the command `go run hello.go` in your terminal/command prompt you should see the output as Hello, world! printed.
    * What happens when you run this command is Go compiles the code into a binary file, __But places this file in a temporary directory and executes this binary file from there, and deletes it after your program finishes.__ This command is useful for testing small programs during the initial development stage of your application.

* `Go Build`
    * So now you want to you have the binary you built for the application to use later or run this binary on a remote computer then you need to use go build command as below which creates a binary file named hello in the current directory(hello.exe in windows)

* `Go Install`
    * This command does perform the exact operation as go build but places the binary in `$GOPATH/bin` directory alongside the binaries of third-party tools installed via go get now if you run `$GOPATH/bin/hello` you will see Hello, world! printed on your terminal.
    * __Suggestion when using `go install` is to use this if you are planning to write tools that you want to use on your own computer or use `go build` if you are to run this application in elsewhere__

## Understanding Packages
* https://thenewstack.io/understanding-golang-packages/
### Workspace
* In Go, programs are kept in a directory hierarchy that is called a workspace. A workspace is simply a root directory of your Go applications. A workspace contains three subdirectories at its root:
    * `src` - This directory contains source files organized as packages. You will write your Go applications inside this directory.
    * `pkg` - This dirctory contains Go package objects
    * `bin` - This directory contains executable programs

### Packages
* __Package `main`__
    * When you develop executable programs, you will use the package “main” for making the package as an executable program. The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.The main function in the package “main” will be the entry point of our executable program
    * When you build shared libraries, you will not have any main package and main function in the package 
* __Importing Packages__
    * When you import packages, the Go compiler will look on the locations specified by the environment variable `GOROOT` and `GOPATH`
        * `GOROOT` contains the standard library that comes with go
    *  The packages that are created by yourself, and third-party packages which you have imported, are available in the `GOPATH` location.
* __Installing 3rd Party Packages__
    * We can download and install third-party Go packages by using “Go get” command. The Go get command will fetch the packages from the source repository and put the packages on the GOPATH location.
    * The following command in the terminal will install “mgo”,  a third-party Go driver package for MongoDB, into your GOPATH, which can be used across the projects put on the GOPATH directory:
    ```
    go get gopkg.in/mgo.v2
    ```
    * After installing the mgo, put the import statement in your programs for reusing the code, as shown below:
    ```
    import (        
        "gopkg.in/mgo.v2" 
        "gopkg.in/mgo.v2/bson"       
    )
    ```
    * The MongoDB driver, mgo,  provides two packages that we have imported in the above import statement.
* __Creating and Reusing Packages__
    * Let’s create a sample program for demonstrating a package. Create a source file “languages.go” for the package “lib”
        * at the location ` github.com/shijuvar/go-samples-thenewstack/packagedemo/lib`
    * Since the “languages.go” belongs to the folder “lib”, the package name will be named “lib”.  All files created inside the lib folder belongs to lib package.
    ```
        package lib
        var languages map[string]string
        func init(){
            languages= make(map[string]string)
            languages["cs"] = "C #"
            languages["js"] = "JavaScript"
            languages["rb"] = "Ruby"
            languages["go"] = "Golang"
        }
        func Get(key string) (string){
            return languages[key]
        }
        func Add(key,value string){
            languages[key]=value
        }
        func GetAll() (map[string]string){
            return languages
        }    
    ```
    * In the above program, we included an init method so that this method will be invoked at the beginning of execution. Let’s build our Go package for reusing with other packages. In the terminal window, go to the location `lib` folder, and run the following command: `go install`
    * The go install command will build the package `lib` which will be available at the `pkg` subdirectory of `GOPATH` (`$GOPATH/pkg`)
    * Let’s create a `main.go` for making an executable program where we want to reuse the code specified in the package `lib`
    ```
        package main
        
        import (
            "fmt"
            "github.com/shijuvar/go-samples-thenewstack/packagedemo/lib"
        )
        
        func main() {
            lib.Add("dr","Dart")
            fmt.Println(lib.Get("dr"))
            languages:=lib.GetAll()
            for _, v := range languages {
                fmt.Println(v)
            }
        }
    ```

## Using Go Modules
* https://blog.golang.org/using-go-modules
* Go's dependency management system that makes dependency version information explicit and easier to manage.
* A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root.

### Creating a new module
* Create a new, empty directory somewhere outside `$GOPATH/src`, cd into that directory, and then create a new source file, hello.go:
```
package hello

func Hello() string {
    return "Hello, world."
}
```
* Let's write a test, too, in hello_test.go:
```
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```
* At this point, the directory contains a package, but not a module, because there is no go.mod file. If we were working in `/home/gopher/hello` and ran go test now, we'd see:
```
$ go test
PASS
ok  	_/home/gopher/hello	0.020s
$
```
* The last line summarizes the overall package test. Because we are working outside $GOPATH and also outside any module, the go command knows no import path for the current directory and makes up a fake one based on the directory name: `_/home/gopher/hello`
* Let's make the current directory the root of a module by using go mod init and then try go test again:
```
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
$ go test
PASS
ok  	example.com/hello	0.020s
$
```
* Congratulations! You’ve written and tested your first module.
```
$ cat go.mod
module example.com/hello

go 1.12
$
```
* The `go.mod` file only appears in the root of the module. Packages in subdirectories have import paths consisting of the module path plus the path to the subdirectory. For example, if we created a subdirectory world, we would not need to (nor want to) run `go mod init` there. The package would automatically be recognized as part of the `example.com/hello` module, with import path `example.com/hello/world`

### Adding a dependancy
* Let's update our hello.go to import rsc.io/quote and use it to implement Hello:
```
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```
* Now let’s run the test again
```
$ go test
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok  	example.com/hello	0.023s
$
```
* The `go` command resolves imports by using the specific dependency module versions listed in `go.mod`. When it encounters an import of a package not provided by any module in `go.mod`, the `go` command automatically looks up the module containing that package and adds it to `go.mod`, using the latest version. (“Latest” is defined as the latest tagged stable (non-prerelease) version, or else the latest tagged prerelease version, or else the latest untagged version.)
```
$ cat go.mod
module example.com/hello

go 1.12

require rsc.io/quote v1.5.2
$
```
* As we saw above, adding one direct dependency often brings in other indirect dependencies too. The command `go list -m all` lists the current module and all its dependencies:
```
$ go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
$
```