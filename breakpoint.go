// This is a PoC/illustrative code

package breakpoint

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/breakpoint", New())
}

type breakpoint struct{}

func (_ *breakpoint) Break(ctx context.Context, o *goja.Object) {
	rt := common.GetRuntime(ctx)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("breakpoint reached `continue` will continue script. Anything typed before that will be executed.\nStack:")
	{ // print stack trace
		var b bytes.Buffer
		stack := rt.CaptureCallStack(0, nil)
		for _, frame := range stack {
			b.WriteRune('\t')
			frame.Write(&b)
			b.WriteRune('\n')
		}
		fmt.Println(b.String())
	}
	for {
		fmt.Print("`continue` will continue script-> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		if "continue" == text {
			return
		}
		if strings.HasPrefix(text, "?") {
			text = strings.TrimPrefix(text, "?")
			v := o.Get(text)
			fmt.Printf("?%s= %s\n", text, v)
			continue
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
