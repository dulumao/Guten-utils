package dump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

// 格式化打印变量(类似于PHP-vardump)
func DD(i ...interface{}) {
	fmt.Println()

	for _, v := range i {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "\t")
		if err := encoder.Encode(v); err == nil {
			fmt.Print(buffer.String())
		} else {
			fmt.Errorf("%s", err.Error())
		}
	}

	fmt.Println()
}

func DD2(a ...interface{}) {
	spew.Dump(a...)
}

func Printf(i ...interface{}) {
	fmt.Println()

	for _, v := range i {
		fmt.Printf("%c[1;40;32m%+v%c[0m\n", 0x1B, v, 0x1B)
	}

	fmt.Println()
}

func Printf2(format string, a ...interface{}) {
	spew.Printf(format, a...)
}
