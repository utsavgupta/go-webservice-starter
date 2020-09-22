package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/utsavgupta/go-webservice-starter/globals"
	"github.com/utsavgupta/go-webservice-starter/handlers"
)

func main() {
	start := time.Now()
	globals.InitConfig()
	globals.InitLogger(globals.APPLICATIONNAME, globals.Config.Stage)

	r := httprouter.New()

	r.GET("/.well-known/live", handlers.Wrap(handlers.Ok))
	r.GET("/.well-known/ready", handlers.Wrap(handlers.Ok))

	r.GET("/", handlers.Wrap(handlers.Default))

	intr := make(chan os.Signal)
	err := make(chan error)

	go func(e chan error) {
		err := http.ListenAndServe(fmt.Sprintf(":%d", globals.Config.Port), r)
		e <- err
	}(err)

	signal.Notify(intr, os.Interrupt)

	globals.Logger.Infof(context.Background(), "Started server in %dms", time.Since(start)/time.Millisecond)

	select {
	case i := <-intr:
		globals.Logger.Infof(context.Background(), "Received interrupt %+v", i)
	case e := <-err:
		globals.Logger.Error(context.Background(), e)
	}

	globals.Logger.Infof(context.Background(), "Stopping server ...")

	globals.Logger.Infof(context.Background(), "Bye!")
}
