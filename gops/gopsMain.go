package main

import (
	"time"

	"github.com/google/gops/agent"
	"github.com/rs/zerolog/log"
)

func main(){
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal().Msg(err.Error())
	}

	time.Sleep(time.Hour)
}
