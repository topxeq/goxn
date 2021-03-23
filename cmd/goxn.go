package main

import (
	"os"

	"github.com/topxeq/goxn"
	"github.com/topxeq/tk"
)

func main() {
	tk.Pl("running %v ...", os.Args[1])

	resultT, errT := goxn.RunScript(tk.LoadStringFromFile(os.Args[1]), tk.GetSwitch(os.Args, "-input="), os.Args, nil)

	tk.CheckErrf("error: %v", errT)

	tk.Pl("%v", resultT)

}
