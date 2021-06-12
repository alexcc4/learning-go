package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"
)

func serverApp(server *http.Server) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		w.Write([]byte("pong"))
	})

	server.Handler = mux

	fmt.Printf("Starting server at port 8000\n")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	server := &http.Server{
		Addr: ":8000",
	}

	g.Go(func() error {
		fmt.Println("staring start app goroutine")
		return serverApp(server)
	})

	// Print time every second
	//g.Go(func() error {
	//	fmt.Println("staring time print goroutine")
	//	ticker := time.NewTicker(1 * time.Second)
	//	defer ticker.Stop()
	//	for {
	//		select {
	//		case <-ticker.C:
	//			fmt.Println(time.Now())
	//		case <-ctx.Done():
	//			fmt.Println("time print exit...")
	//			return ctx.Err()
	//		}
	//	}
	//})
	g.Go(func() error {
		fmt.Println("staring time print goroutine")
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println(time.Now())
				time.Sleep(1)
			case <-ctx.Done():
				fmt.Println("time print exit...")
				return ctx.Err()
			}
		}
	})

	g.Go(func() error {
		fmt.Println("staring shutdown app goroutine")
		select {
		case <-ctx.Done():
			fmt.Println("http server stop")
			return server.Shutdown(ctx)

		}

	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case s := <-quit:
			return errors.Errorf("get os signal: %v", s)
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("group error: ", err)
	} else {
		fmt.Println("group all done successfully!")
	}
}
