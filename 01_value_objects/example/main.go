package main

import (
	"fmt"
	"time"

	"github.com/jdgonzalez907/go-patterns/01_value_objects/example/domain"
)

func strToStrPointer(s string) *string {
	return &s
}

func main() {
	juandaBase, err := domain.NewUserDeclaredData(1, nil, nil, strToStrPointer("juandaelteso@gmail.com"), nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("API USER: ")
	fmt.Println(juandaBase)

	juandaNew, err := domain.NewUserDeclaredData(1, nil, strToStrPointer("@devjuanda"), nil, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("VAULT: ")
	fmt.Println(juandaNew)

	juandaBase.Merge(juandaNew)

	fmt.Println()
	fmt.Println("USER MERGED:")
	fmt.Println(juandaBase)

	request, err := domain.NewDeclaredDataUpdateRequest(strToStrPointer("meli dev Inc"), strToStrPointer("@jdgonzalez907"), strToStrPointer("jdgonzalez907@gmail.com"), strToStrPointer("57"), strToStrPointer("3147733992"), "miintegrador", time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("REQUEST: ")
	fmt.Println(request)

	changelog, err := juandaBase.Update(request, "zxc-123-zxc-123", time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("CHANGELOG: ")
	fmt.Println(changelog)

	// domain.UserDeclaredDataRepository.Save(juandaBase) -> API USERS
	// domain.UserDeclaredDataRepository.Save(juandaBase) -> API VAULT
	// domain.ChangelogRepository.Save(changelog)
}
