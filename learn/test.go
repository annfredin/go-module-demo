package main

import (
	"fmt"
)
func init() {
 fmt.Println("app/fetch-version.go ==> init()")
}
func fetchVersion() string {
 fmt.Println("app/fetch-version.go ==> fetchVersion()")
 return "1.0.1"
}