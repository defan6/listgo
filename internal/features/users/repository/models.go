package repository

import core_domain "github.com/defan6/listgo/internal/core/domain"

type UserModel struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUserDomainFromUserModel(userModel UserModel) core_domain.User {
	return core_domain.User{
		ID:          userModel.ID,
		Version:     userModel.Version,
		FullName:    userModel.FullName,
		PhoneNumber: userModel.PhoneNumber,
	}
}

func NewUserDomainsFromUserModels(userModels []UserModel) []core_domain.User {
	userDomains := make([]core_domain.User, 0, len(userModels))
	for _, model := range userModels {
		userDomain := NewUserDomainFromUserModel(model)
		userDomains = append(userDomains, userDomain)
	}

	return userDomains
}
