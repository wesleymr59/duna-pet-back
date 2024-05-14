package usecases

type HealthyUsecase struct{}

func NewHealthyService() *HealthyUsecase {
	return &HealthyUsecase{}
}

func (s *HealthyUsecase) GetHelloMessage() string {
	return "Hello, World!"
}

func (s *HealthyUsecase) GetIndexMessage() string {
	message := "Api is Running"
	return message
}
