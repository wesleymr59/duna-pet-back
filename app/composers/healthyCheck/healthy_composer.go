package composers

import (
	usecases "duna-pet-back/app/usecases/healthyCheck"
)

type HealthyComposer struct{}

func NewHealthyComposer() *HealthyComposer {
	return &HealthyComposer{}
}

func (c *HealthyComposer) ComposeHelloService() *usecases.HealthyUsecase {
	return usecases.NewHealthyService()
}
