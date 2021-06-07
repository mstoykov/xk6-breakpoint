// This is a PoC/illustrative code

package breakpoint

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/breakpoint", New())
}

type breakpoint struct{}

func (_ *breakpoint) Break(ctx context.Context) {
	rt := common.GetRuntime(ctx)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("breakpoint reached `continue` will continue script. Anything typed before that will be executed.")
	for {
		fmt.Print("`continue` will continue script-> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		if "continue" == text {
			return
		}
		v, err := rt.RunString(text)
		if err != nil {
			fmt.Printf("Got Error %s\n", err)
		}
		fmt.Printf(">>>> %s\n", v)
	}
}

func New() *breakpoint {
	return &breakpoint{}
}
