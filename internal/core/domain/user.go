package core_domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/defan6/listgo/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(ID int, Version int, fullname string, phoneNumber *string) User {
	return User{
		ID:          ID,
		Version:     Version,
		FullName:    fullname,
		PhoneNumber: phoneNumber,
	}
}

func NewUninitializedUser(fullname string, phoneNumber *string) User {
	return NewUser(UninitializedUserID, UninitializedVersion, fullname, phoneNumber)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid `FullName` length: %d : %w", fullNameLength, core_errors.ErrBadRequest)
	}
	if u.PhoneNumber != nil {
		phoneNumber := *u.PhoneNumber
		phoneNumberLength := len([]rune(phoneNumber))

		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("invalid `PhoneNumber` length: %d : %w", phoneNumberLength, core_errors.ErrBadRequest)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(phoneNumber) {
			return fmt.Errorf("invalid `PhoneNumber` format: %w", core_errors.ErrBadRequest)
		}
	}

	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (u *UserPatch) Validate() error {
	if u.FullName.Set && u.FullName.Value == nil {
		return fmt.Errorf("user `FullName` cannot patch to empty: %w", core_errors.ErrBadRequest)
	}
	return nil
}

func (u *User) ApplyPatch(patchUser UserPatch) error {
	if err := patchUser.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u
	if patchUser.FullName.Set {
		tmp.FullName = *patchUser.FullName.Value
	}

	if patchUser.PhoneNumber.Set {
		tmp.PhoneNumber = patchUser.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
