package main

import (
	"github.com/jdgonzalez907/go-patterns/full_example/core/usecases"
	"github.com/jdgonzalez907/go-patterns/full_example/infrastructure"
)

func main() {
	serverClock := infrastructure.NewServerClock()
	inMemoryApiUsersUserDeclaredDataRepository := infrastructure.NewInMemoryApiUsersUserDeclaredDataRepository()
	inMemoryVaultUserDeclaredDataRepository := infrastructure.NewInMemoryVaultUserDeclaredDataRepository()
	inMemoryChangelogRepository := infrastructure.NewInMemoryChangelogRepository()

	updateDeclaredDataUseCase := usecases.NewUpdateDeclaredDataUseCase(inMemoryApiUsersUserDeclaredDataRepository, inMemoryVaultUserDeclaredDataRepository, inMemoryChangelogRepository, serverClock)

	tmpNickname := "jdgonzalez907"
	tmpCountryCode := "57"
	tmpPhoneNumber := "3147733992"

	input := usecases.UpdateDeclaredDataUseCaseInput{
		UserID:      1,
		Nickname:    &tmpNickname,
		CountryCode: &tmpCountryCode,
		PhoneNumber: &tmpPhoneNumber,
		RequestedBy: "auth-phone-enrollment-api",
	}

	_, err := updateDeclaredDataUseCase.Execute(input)
	if err != nil {
		panic(err)
	}
}
