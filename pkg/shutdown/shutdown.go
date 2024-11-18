package shutdown

import (
	"log"
	"os"
	"os/signal"
)

func Init() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server")
}
