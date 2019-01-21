package run

import (
	"github.com/lygo/runner"
	"context"
	"time"
	"fmt"
)

func New() (app *runner.App, err error) {
	app = runner.New()

	defer func() {
		if err != nil {
			app.Shutdown()
		}
	}()

	ctx, cancelGraceful := context.WithCancel(context.Background())

	app.Runners = append(app.Runners, func() error {
		for {
			select {
			case <-time.After(time.Second*3): {
				fmt.Println("log runner")
			}
			case <-ctx.Done(): return nil
			}
		}
	})

	app.Slams = append(app.Slams, func() error {
		go cancelGraceful()
		<-ctx.Done()
		return nil
	})

	return
}
