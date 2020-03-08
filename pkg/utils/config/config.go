// Package config ...
package config

import (
	"sync"

	"github.com/dup2X/gomicrosvr/pkg/config"
	"github.com/dup2X/gomicrosvr/pkg/logger"
)

const (
	moduleSMS  = "Sms"
	modulePush = "Push"
)

type configItemType uint8

const (
	configItemTypeInt configItemType = iota
	configItemTypeBool
	configItemTypeString
)

var (
	configSet map[string]config.Configer
	lock      sync.Mutex
)

/**
* 不在自动加载rpc_server.conf等配置文件
* 加载策略：使用时加载，保证本地缓存
**/

/**
* 每次从缓存中获取命名空间的配置
* 从某个文件获取配置文件
* 单位是文件
**/
func getConfig(configPath string) (config.Configer, error) {
	configSet = make(map[string]config.Configer)
	if value, ok := configSet[configPath]; ok {
		// 配置文件已经加载过，直接返回
		return value, nil
	}
	// toml 格式
	conf, err := config.NewConfigWithFormatType(config.ConfFormatTypeToml, configPath)
	if nil != err {
		logger.Errorf(nil, logger.DLTagConfLoad, "load config[%s] failed,err:%v", configPath, err)
		return nil, wrapMissFileError(configPath)
	}
	lock.Lock()
	// 单位是文件，不是section
	configSet[configPath] = conf
	lock.Unlock()
	return conf, err

}

/**
* 获取配置文件内部内容，仅限一级目录
**/
func getConfigModule(moduleName string, filePath string) (config.Sectioner, error) {
	conf, err := getConfig(filePath)
	if nil != err {
		return nil, err
	}
	module, err := conf.GetSection(moduleName)
	if nil != err {
		return nil, wrapMissModuleError(filePath, moduleName)
	}
	return module, nil
}

/**
* 获取配置文件中module的bool型标签值
**/
func getConfigBoolItem(moduleName, itemName string, filePath string) (bool, error) {
	conf, err := getConfig(filePath)
	if nil != err {
		return false, err
	}
	item, err := conf.GetBoolSetting(moduleName, itemName)
	if nil != err {
		return false, wrapMissItemError(filePath, moduleName, itemName)
	}
	return item, nil
}

/**
* 获取配置文件中module的int型标签值
**/
func getConfigIntItem(moduleName, itemName string, filePath string) (int, error) {
	conf, err := getConfig(filePath)
	if nil != err {
		return 0, err
	}
	item, err := conf.GetIntSetting(moduleName, itemName)
	if nil != err {
		return 0, wrapMissItemError(filePath, moduleName, itemName)
	}
	return int(item), nil
}

/**
* 获取配置文件中module的string型标签值
**/
func getConfigStringItem(moduleName string, itemName string, configPath string) (string, error) {
	conf, err := getConfig(configPath)
	if nil != err {
		return "", err
	}
	item, err := conf.GetSetting(moduleName, itemName)
	if nil != err {
		return "", wrapMissItemError(configPath, moduleName, itemName)
	}
	return item, nil
}

// GetConfig ...
func GetConfig(filePath string) (config.Configer, error) {
	return getConfig(filePath)
}

// GetConfigModule ...
func GetConfigModule(moduleName string, filePath string) (config.Sectioner, error) {
	return getConfigModule(moduleName, filePath)
}

// GetConfigBoolItem ...
func GetConfigBoolItem(moduleName, itemName string, filePath string) (bool, error) {
	return getConfigBoolItem(moduleName, itemName, filePath)
}

// GetConfigIntItem ...
func GetConfigIntItem(moduleName, itemName string, filePath string) (int, error) {
	return getConfigIntItem(moduleName, itemName, filePath)
}

// GetConfigStringItem ...
func GetConfigStringItem(moduleName, itemName string, filePath string) (string, error) {
	return getConfigStringItem(moduleName, itemName, filePath)
}

// GetSmsConfig ...
func GetSmsConfig(filePath string) (config.Sectioner, error) {
	return getConfigModule(moduleSMS, filePath)
}

// GetPushConfig ...
func GetPushConfig(filePath string) (config.Sectioner, error) {
	return getConfigModule(modulePush, filePath)
}

// GetSDKConfig ...
func GetSDKConfig(moduleName string, filePath string) (config.Sectioner, error) {
	return getConfigModule(moduleName, filePath)
}
