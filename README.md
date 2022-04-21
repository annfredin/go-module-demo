## MODULE - GOLANG:

define multiple file with same/one package, then you can access all identiers(variable,ufunction, struct) even with private access scope(small case name)

but when do go run <main_file> then you face issue like identifier undefined, since it does not have idea about other file, so we can test this with <main_file>

- go run \*.go -> will look all the file under directory
- go build or go install -> then you get access..

Within same package Puble scope not required for access...

#### Directory Name vs Pakage Name [ref math-test dire]

- Package name need not to be same as directory name.
- For dependency ref we need to use directory name, and for access identifier we need to user packagename/aliasname

#### Blank Identifier (\_)[ ref blank-dentifer dire]

- we use blank identifer(\_) in order to bypass/avoid compilation error
- in Dependency ref blank identifier used to intitlize/execute init function, where we used write code for any external system connections like DB Connection etc. if blank identifier file init will get more preference in execution order.

#### Add a dependency to your project

- Directly adding it to the go.mod file
- Do a go get
- Add the dependency to your source code and do a go mod tidy

- Directly adding it to the go.mod file:
  Add below dependency to the go.mod file
  `require github.com/pborman/uuid v1.2.1`
  then run the below command `go mod download`, this will download all direct & indirect dependencies and update "go.sum" file

- Do a go get:
  `export GO111MODULE=on go get github.com/pborman/uuid`

- Add the dependency to your source code and do a go mod tidy:
  - Add any dependency which is imported in the source files
  - Remove any dependency which is mentioned in the go.mod file but not imported in any of the source files.

`go mod tidy`

#### Adding a vendor directory:

    It will create a vendor directory inside your project directory, and hold all the dependencies,
    `go mod vendor`

#### Module Import Path --> Naming:

- The module is a utility module and you plan to publish your module:
  - Module name should match the URL of the repo which host that module
- The module is a utility module and you don’t plan to publish your module:
  - use the utility module locally only -> name canbe anything
- The module is a executable module:
  - In this case also module import path can be anything. The module import path can be a non-url even if you plan to commit your module into VCS as it will not be used by any other module

Also couple of points to note about upgrading module

- To upgrade a dependency to its latest patch version only, use below command
  `go get -u=patch`
- To upgrade a dependency to a specific version, use below command
  `go get dependency@version`
- To upgrade a dependency to a specific commit, use below command
  `go get @commit_number`
- To upgrade all dependency to their latest minor and patch version, use below command
  `go get ./...`
