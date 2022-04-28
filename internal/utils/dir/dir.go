package dir

import (
	"github.com/alkaidchen/Dubhe/internal/log"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

const DubheDir = ".dubhe"

func GetOrCreateDubheHome() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Error(err)
		return "", err
	}
	dubheHome := path.Join(homeDir, DubheDir)
	_, err = os.Stat(dubheHome)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(dubheHome, os.ModePerm); err != nil {
				log.Error(err)
				return "", err
			}
		} else {
			log.Error(err)
			return "", err
		}
	}
	return dubheHome, nil
}
