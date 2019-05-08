package util
import (
	"vision/core/models"
	"strings"
	"errors"
	//"fmt"
)

func CheckAcls(path string, configJson *models.ConfigModel) (error) {
	if configJson.AllowAll == true {
		blockFor := configJson.BlockFor
		if len(blockFor) == 0 {
			return nil
		}
		for _, blockedEntity := range configJson.BlockFor {
			if strings.HasPrefix(path, blockedEntity) {
				return errors.New("FILE_NOT_ALLOWED_TO_BE_VIEWED")
			}
		}

		return nil
	} else {
		allowFor := configJson.AllowFor
		if len(allowFor) == 0 {
			return errors.New("FILE_NOT_ALLOWED_TO_BE_VIEWED")
		}

		for _, allowedEntity := range configJson.AllowFor {
			if strings.HasPrefix(path, allowedEntity) {
				return nil
			}
		}

		return errors.New("FILE_NOT_ALLOWED_TO_BE_VIEWED")
	}
}
