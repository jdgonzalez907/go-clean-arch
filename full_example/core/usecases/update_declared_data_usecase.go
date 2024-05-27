package usecases

import (
	"fmt"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain/changelog"
	declareddata "github.com/jdgonzalez907/go-patterns/full_example/core/domain/declared_data"
	"github.com/jdgonzalez907/go-patterns/full_example/core/domain/shared"
)

type (
	UpdateDeclaredDataUseCase struct {
		vaultUserDeclaredDataRepository    declareddata.UserDeclaredDataRepository
		apiUsersUserDeclaredDataRepository declareddata.UserDeclaredDataRepository
		changelogRepository                changelog.ChangelogRepository
		clock                              shared.Clock
	}

	UpdateDeclaredDataUseCaseInput struct {
		UserID      int64
		Nickname    *string
		CountryCode *string
		PhoneNumber *string
		RequestedBy string
	}

	UpdateDeclaredDataUseCaseOutput struct {
		Changelog changelog.Changelog
	}
)

func NewUpdateDeclaredDataUseCase(apiUsersUserDeclaredDataRepository, vaultUserDeclaredDataRepository declareddata.UserDeclaredDataRepository, changelogRepository changelog.ChangelogRepository, clock shared.Clock) UpdateDeclaredDataUseCase {
	return UpdateDeclaredDataUseCase{
		apiUsersUserDeclaredDataRepository: apiUsersUserDeclaredDataRepository,
		vaultUserDeclaredDataRepository:    vaultUserDeclaredDataRepository,
		changelogRepository:                changelogRepository,
		clock:                              clock,
	}
}

func (u *UpdateDeclaredDataUseCase) Execute(input UpdateDeclaredDataUseCaseInput) (UpdateDeclaredDataUseCaseOutput, error) {
	userBase, err := u.apiUsersUserDeclaredDataRepository.FindByID(input.UserID)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}
	fmt.Println("APIUser: \n" + userBase.String())

	userNew, err := u.vaultUserDeclaredDataRepository.FindByID(input.UserID)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}
	fmt.Println("VaultUser: \n" + userNew.String())

	userBase.Merge(userNew)
	fmt.Println()
	fmt.Println("MergedUser: \n" + userBase.String())

	err = u.apiUsersUserDeclaredDataRepository.Save(userBase)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}

	err = u.vaultUserDeclaredDataRepository.Save(userBase)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}

	ocurredOn := u.clock.Now()
	request, err := declareddata.NewDeclaredDataUpdateRequest(input.Nickname, input.CountryCode, input.PhoneNumber, input.RequestedBy, ocurredOn)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}
	fmt.Println()
	fmt.Println("Request: \n" + request.String())

	requestedOn := u.clock.Now()
	changelog, err := userBase.Update(request, requestedOn)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}
	fmt.Println()
	fmt.Println("Changelog: \n" + changelog.String())

	fmt.Println()
	fmt.Println("UserCurrent: \n" + userBase.String())

	err = u.changelogRepository.Save(changelog)
	if err != nil {
		return UpdateDeclaredDataUseCaseOutput{}, err
	}

	if userBase.PhoneChanged() {
		fmt.Println()
		associatedUsers, err := u.vaultUserDeclaredDataRepository.FindByPhone(*userBase.Phone())
		if err != nil {
			return UpdateDeclaredDataUseCaseOutput{}, err
		}

		for _, user := range associatedUsers {
			if user.ID() == userBase.ID() {
				continue
			}
			user.RecyclePhone()
			fmt.Println("Recicled Phone: \n" + user.String())
		}

		err = u.vaultUserDeclaredDataRepository.SaveAll(associatedUsers)
		if err != nil {
			return UpdateDeclaredDataUseCaseOutput{}, err
		}
	}

	return UpdateDeclaredDataUseCaseOutput{
		Changelog: changelog,
	}, nil
}
