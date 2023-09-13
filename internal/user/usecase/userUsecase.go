package usecase

import adapter "persona/internal/user/adapter/in"

type UserUseCase struct{}

func (usecase UserUseCase) RegisterMembership(command *adapter.UserRegisterCommand) error {
	return nil
}

func (usecase UserUseCase) DropdownMembership(command *adapter.UserDropdownCommand) error {
	return nil
}
