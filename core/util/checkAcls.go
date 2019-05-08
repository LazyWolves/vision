package util
import (
	"vision/core/models"
)

func checkAcls(configJson *models.configModel) (error) {
	if configJson.AllowAll == true {
		blockFor := configJson.BlockFor
		if len(blockFor) == 0 {
			return nil
		}

		
	}
}
