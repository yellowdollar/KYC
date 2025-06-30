package service

import (
	"KYC/iternals/models"
	"KYC/iternals/repository"
	"KYC/utils"
)

func CreateUser(u models.User) (*models.User, error) {
	// hashing pass
	u.Password = utils.GenerateHash(u.Password)

	// call function from repository
	result, err := repository.CreateUser(u)

	// checking for errors
	if err != nil {
		return nil, err
	}

	// After creating user row, creates account row
	err = repository.CreateUserProfile(int(result.ID))
	if err != nil {
		return nil, err
	}

	// return result
	return result, nil
}

func GetUserByLogin(userLogin string) (*models.User, error) {
	user, err := repository.GetUserByLogin(userLogin)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ComparePasswords(inputPassword string, hashedPassword string) bool {
	return utils.GenerateHash(inputPassword) == hashedPassword
}

// func CheckVerificationStatus(userID int) (string, error) {
// 	uVerification, err := repository.GetUserVerificationStatus(userID)
// 	if err != nil {
// 		return "", err
// 	}

// 	switch uVerification {
// 	case "unverified":
// 		return "unverified", nil
// 	case "pending":
// 		return "pending", nil
// 	}

// 	return "verified", nil
// }

func GetUserByID(userID int) (*models.User, error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserIdentification(userID int) (*models.User, error) {
	user, err := repository.UpdateUserIdentification(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
