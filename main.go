package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"learngoframe/internal/cmd"
	_ "learngoframe/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
