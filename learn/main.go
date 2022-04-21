package main

import (
	"fmt"

	_ "module.com/sample/learn/blank-identifier" // blank identifier
	"module.com/sample/learn/math-test"
	mathAlias "module.com/sample/learn/math-test"
)
func init() {
 fmt.Println("app/entry.go ==> init()")
}
//var myVersion = fetchVersion()
func main() {
 fmt.Println("Calling --> app/fetch-version.go ==> fetchVersion()")
 fmt.Println("version ===> ", fetchVersion())

 otherPackageAccess()
}


//math --> Package name, need not to be same as directory name
// depe ref we need to use directory name, access we need to user packagename/aliasname
func otherPackageAccess(){
    fmt.Println()
    fmt.Println("============ Other Oackage ===========")
    fmt.Println("MATH => ADD ---> " ,math.Add(12,10))
    fmt.Println("MATH => SUB ---> ", mathAlias.Sub(12,10))
}