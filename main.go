package main

import (
	_ "github.com/ducdung17/my-awesome-extension/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
