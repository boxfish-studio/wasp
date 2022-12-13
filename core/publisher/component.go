package publisher

import (
	"go.uber.org/dig"

	"github.com/iotaledger/hive.go/core/app"
	"github.com/iotaledger/wasp/packages/daemon"
	"github.com/iotaledger/wasp/packages/publisher"
)

func init() {
	CoreComponent = &app.CoreComponent{
		Component: &app.Component{
			Name:     "Publisher",
			DepsFunc: func(cDeps dependencies) { deps = cDeps },
			Provide:  provide,
			Run:      run,
		},
	}
}

var (
	CoreComponent *app.CoreComponent
	deps          dependencies
)

type dependencies struct {
	dig.In
	Publisher *publisher.Publisher
}

func provide(c *dig.Container) error {
	type subDeps struct {
		dig.In
	}

	type subResult struct {
		dig.Out

		Publisher *publisher.Publisher
	}

	if err := c.Provide(func(deps subDeps) subResult {
		publ := publisher.NewPublisher(
			CoreComponent.Logger(),
		)
		return subResult{
			Publisher: publ,
		}
	}); err != nil {
		CoreComponent.LogPanic(err)
	}

	return nil
}

func run() error {
	err := CoreComponent.Daemon().BackgroundWorker(
		"Publisher",
		deps.Publisher.Run,
		daemon.PriorityPublisher,
	)
	if err != nil {
		panic(err)
	}
	return nil
}