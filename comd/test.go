package main

import (
	"fmt"

	"github.com/slusenti/noskull/lib/base"
	"github.com/slusenti/noskull/lib/utils"
)

func main() {
	id := utils.GenerateID()
	fmt.Println(id)
	fmt.Println(utils.TimeFromID(id))
	b := base.NewConfig()
	b.RootPWD = "ok"
	b.ToGOB64()
	c := base.ConfigFromGOB64()
	fmt.Println(c)
}
