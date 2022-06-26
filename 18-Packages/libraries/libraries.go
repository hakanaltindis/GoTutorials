package main

import (
	"fmt"

	"utils"
)

/*
 * First, run "go mod init <module_name>" command in utils folder
 * Second, run "go mod init <module_name>" command in libraries folder
 * Third, run "go mod edit -replace utils=../utils" command in libraries folder
 */

func main() {

	// 3. Use local modules
	nl, length := utils.FullName("Hakan", "Altındiş")
	fmt.Printf("Fullname: %v, number of chars: %v\n", nl, length)
}
