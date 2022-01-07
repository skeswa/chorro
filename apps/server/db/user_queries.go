package db

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/skeswa/chorro/apps/server/db/model"
	"gorm.io/gorm"
)

// Returns true if a user associated with the specified userID exists.
func (db *DB) DoesUserExist(userID uint) (bool, error) {
	doesExist := false

	if err := db.client.Model(&model.User{}).
		Select("count(*) > 0").
		Where(fmt.Sprintf("%s = ?", model.UserColumnNames.ID), userID).
		Find(&doesExist).
		Error; err != nil {
		return doesExist, errors.Wrap(err, "Failed to check if user exists")
	}

	return doesExist, nil
}

// Returns the user identified by the specified googleUserID, or nil if no such
// user exists.
func (db *DB) GoogleUser(googleUserID string) (*model.User, error) {
	user := model.User{}

	if err := db.client.Where(
		fmt.Sprintf("%s = ?", model.UserColumnNames.Google.UserID),
		googleUserID,
	).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "Failed to read Google user")
	}

	return &user, nil
}

// Returns the ID of the user identified by the specified googleUserID, or nil
// if no such user exists.
func (db *DB) UserIDOfGoogleUser(googleUserID string) (*uint, error) {
	user := model.User{}

	if err := db.client.Select(model.UserColumnNames.ID).
		Where(
			fmt.Sprintf("%s = ?", model.UserColumnNames.Google.UserID),
			googleUserID,
		).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "Failed to read Google user")
	}

	return &user.ID, nil
}

// Returns the user identified by the specified userID, or nil if no such user
// exists.
func (db *DB) User(userID uint) (*model.User, error) {
	user := model.User{}

	if err := db.client.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "Failed to read user")
	}

	return &user, nil
}
