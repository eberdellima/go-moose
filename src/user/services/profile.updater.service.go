package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/user/inputs"
	"go-moose/src/user/utils"

	AuthorizationServices "go-moose/src/authorization/services"
)

// UpdateUserProfile tries to update user's data partially and quick returns if error
// is encountered. If no error is found, it updates the user's data successfully
func UpdateUserProfile(user *models.User, payload inputs.ProfileInput) error {

	var err error

	if err = partialUpdateEmail(user, payload.Email); err != nil {
		return err
	}

	if err = partialUpdateUsername(user, payload.Username); err != nil {
		return err
	}

	if err = partialUpdatePassword(user, payload); err != nil {
		return err
	}

	database.DB.Save(user)

	return nil
}

// partialUpdateEmail checks if a user with the provided email already exists
// and returns error if true. Otherwise nil is returned instead
func partialUpdateEmail(providedUser *models.User, email string) error {

	if email == "" {
		return nil
	}

	user := models.User{}

	database.DB.Where(models.User{Email: email}).First(&user)

	if user.ID != 0 {
		return &utils.EmailIsTakenError{}
	}

	providedUser.Email = email

	return nil
}

// partialUpdateUsername checks if a user with the provided username already exists
// and returns error if true. Otherwise nil is returned instead
func partialUpdateUsername(providedUser *models.User, username string) error {

	if username == "" {
		return nil
	}

	user := models.User{}

	database.DB.Where(models.User{Username: username}).First(&user)

	if user.ID != 0 {
		return &utils.UsernameIsTakenError{}
	}

	providedUser.Username = username

	return nil
}

// partialUpdatePassword checks if provided old password is correct before updating user's password
// In case it's not correct error is returned
func partialUpdatePassword(providedUser *models.User, payload inputs.ProfileInput) error {

	if payload.OldPassword == "" || payload.NewPassword == "" {
		return nil
	}

	ok, err := AuthorizationServices.ComparePasswords(payload.OldPassword, providedUser.Password)
	if !ok || err != nil {
		return &utils.IncorrectPasswordError{}
	}

	hashedPassword, _ := AuthorizationServices.GeneratePassword(payload.NewPassword)

	providedUser.Password = hashedPassword

	return nil
}
