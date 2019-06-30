// download package with `go get -v github.com/mactsouk/go/simpleGitHub`

// delete package from the system with
// `go clean -i -v -x github.com/mactsouk/go/simpleGitHub`
// &
// `rm -rf ~/go/src/github.com/mactsouk/go/simpleGitHub`

package main

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
