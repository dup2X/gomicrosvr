/**
*  Author: gongfengguang
*  Description: ---
*  Create: 2017-05-10 14:45:47
*  Last Modified: 2017-05-10 14:45:47
**/

package config

import (
	"fmt"
)

func wrapMissFileError(fileName string) error {
	return fmt.Errorf("[config] can not find config file: %s ", fileName)
}

func wrapMissModuleError(fileName, moduleName string) error {
	return fmt.Errorf("[config] can not find config module: %s %s", fileName, moduleName)
}

func wrapMissItemError(fileName, moduleName, itemName string) error {
	return fmt.Errorf("[config] can not find config item: %s %s %s ", fileName, moduleName, itemName)
}

func wrapGetPwdError() error {
	return fmt.Errorf("[config] get pwd error")
}

func wrapGetGoPathError() error {
	return fmt.Errorf("[config] get GOPATH error")
}

func wrapMissGoCommonPathError() error {
	return fmt.Errorf("[config] can not find 'go-common' path")
}
