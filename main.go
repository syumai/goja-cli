package main

import (
	"fmt"
	"github.com/dop251/goja_nodejs/require"
	"log"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
)

var simplePrinter console.Printer = console.PrinterFunc(func(s string) {
	fmt.Println(s)
})

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) < 2 {
		fmt.Println("file name must be given. (e.g. example.js)")
		return nil
	}
	filePath := os.Args[1]
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	registry := new(require.Registry)
	vm := goja.New()
	registry.Enable(vm)
	require.RegisterNativeModule(console.ModuleName, func(runtime *goja.Runtime, module *goja.Object) {
		console.RequireWithPrinter(simplePrinter)(runtime, module)
	})
	console.Enable(vm)
	if _, err := vm.RunScript(filePath, string(b)); err != nil {
		return err
	}
	return nil
}
