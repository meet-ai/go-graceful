package main

import (
	"fmt"

	"github.com/samber/do"
)
func main() {
    injector := do.New()

    // provides CarService
    do.Provide(injector, NewCarService)

    // provides EngineService
    do.Provide(injector, NewEngineService)

    car := do.MustInvoke[*CarService](injector)
    car.Start()
    // prints "car starting"

    do.HealthCheck[EngineService](injector)
    // returns "engine broken"

    // injector.ShutdownOnSIGTERM()    // will block until receiving sigterm signal
    injector.Shutdown()
    // prints "car stopped"
}

type EngineService interface{}

func NewEngineService(i *do.Injector) (EngineService, error) {
    return &engineServiceImplem{}, nil
}

type engineServiceImplem struct {}

// [Optional] Implements do.Healthcheckable.
func (c *engineServiceImplem) HealthCheck() error {
	return fmt.Errorf("engine broken")
}

func NewCarService(i *do.Injector) (*CarService, error) {
    engine := do.MustInvoke[EngineService](i)
    car := CarService{Engine: engine}
    return &car, nil
}

type CarService struct {
	Engine EngineService
}

func (c *CarService) Start() {
	println("car starting")
}

// [Optional] Implements do.Shutdownable.
func (c *CarService) Shutdown() error {
	println("car stopped")
	return nil
}