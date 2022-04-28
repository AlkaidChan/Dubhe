package main

import (
	"github.com/alkaidchen/Dubhe/internal/app"
	"github.com/alkaidchen/Dubhe/internal/log"
)

func main() {
	dubhe := app.New()
	if err := dubhe.Init(); err != nil {
		log.Fatal("application init failed, ", err)
	}
	if err := dubhe.Run(); err != nil {
		log.Fatal("application run failed, ", err)
	}
}
