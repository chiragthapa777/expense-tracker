package user

import (
	"fmt"

	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
)

func UpdateProfilePicture(body dto.UserProfileUpdateDto, currentUser models.User, option repository.Option) error {
	fileRepository := repository.NewFileRepository()

	// Handle profile removal
	if body.RemoveProfile {
		profile, err := fileRepository.FindByUserProfileId(currentUser.ID, option)
		if err != nil {
			return err
		}
		if profile != nil {
			if err := fileRepository.Delete(profile.ID, option); err != nil {
				return err
			}
		}
		return nil
	}

	// Handle profile update
	if body.ProfileId != nil {
		// Get current profile if exists
		currentProfile, err := fileRepository.FindByUserProfileId(currentUser.ID, option)
		if err != nil {
			return err
		}

		// Get new profile
		newProfile, err := fileRepository.FindByID(*body.ProfileId, option)
		if err != nil {
			return err
		}
		if newProfile == nil {
			return fmt.Errorf("profile with id %s not found", *body.ProfileId)
		}

		// If current profile exists and is different from new profile, delete it
		if currentProfile != nil && currentProfile.ID != *body.ProfileId {
			if err := fileRepository.Delete(currentProfile.ID, option); err != nil {
				return err
			}
		}

		// Update new profile with user ID
		newProfile.UserProfileID = &currentUser.ID
		if err = fileRepository.Update(newProfile, option); err != nil {
			return err
		}
	}

	return nil
}
