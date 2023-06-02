# GO

## GO commands
- go build -> compiles a bunch of go source code files
- go run -> compiles and executes on or two files
- go fmt -> formats all the code in each file in the current direcotry
- go install -> compiles and "installs" a package.
- go get -> downloads the raw source code of someone else's package
- go test -> runs any tests associated with the current project

## GO packages
* package can be interpreted as project/workspace
* it is a colleciton of common source files
* all source files within same dir should have same package

* there are two types of pacakges in Go:
    1. Executable - generates a file that we can run
    2. Reusable - Code used as "helpres". Good place to put reusable logic

* All executable pacakges have same name -> <b>main</b>. Other names for pacakges makes them reusable

* import keyword does importing of mentioned packages, which basically means import all of codes, functions and other from this package so I can use it within mine
* <b>fmt</b> package is Standard lib
* <a href="hhtps://golang.org/pkg/">here</a> we can find a list of all packages included in standard library which we can import and use. Every package has a separate package with doc explanation of usable functions.
