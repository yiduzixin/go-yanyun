package libs

import (
	"dropzone-go/utils"
	"errors"
)

func ValidateAccess(k, s string) error {
	access, secret := utils.GetAccessInfo()
	if k != access || s != secret {
		return errors.New("验证不通过")
	}
	return nil
}
