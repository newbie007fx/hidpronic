package usecase_test

import (
	"hidroponic/internal/module/automation/ports"
	"hidroponic/internal/module/automation/ports/mocks"
	"hidroponic/internal/module/automation/usecase"
	installationConfMocks "hidroponic/internal/module/installationconfig/ports/mocks"
	plantMocks "hidroponic/internal/module/plants/ports/mocks"
	"os"
	"testing"
)

var automationUsc ports.Usecase

func setup() {
	automationUsc = usecase.New(&mocks.MockRepository{}, &installationConfMocks.MockRepository{}, &plantMocks.MockRepository{})
}

func TestMain(m *testing.M) {
	setup()

	// Run all the tests
	code := m.Run()

	// Exit with the received code
	os.Exit(code)
}
