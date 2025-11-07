package cli

import (
	"path/filepath"

	"github.com/segmentfault/pacman/log"
)

const (
	DefaultConfigFileName = "config.yaml"
)

var (
	ConfigFileDir  = "/conf/"
	UploadFilePath = "/uploads/"
	//I18nPath       = "/i18n/"
	//CacheDir       = "/cache/"
)

func FormatAllPath(dataDirPath string) {
	ConfigFileDir = filepath.Join(dataDirPath, ConfigFileDir)
	UploadFilePath = filepath.Join(dataDirPath, UploadFilePath)
	//I18nPath = filepath.Join(dataDirPath, I18nPath)
	//CacheDir = filepath.Join(dataDirPath, CacheDir)
	log.Info("config file path:", ConfigFileDir)
	log.Info("upload file path:", UploadFilePath)
}

// GetConfigFilePath get config file path
func GetConfigFilePath() string {
	return filepath.Join(ConfigFileDir, DefaultConfigFileName)
}
