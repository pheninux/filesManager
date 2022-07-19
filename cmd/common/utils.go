package common

import (
	"fileManager2/pkg/models"
	"fmt"
	"os"
)

type Utils struct {
	Slash string
}

func (u *Utils) ValidateArgs(p *models.DataTemplate) error {
	var err error
	//var regex = "[a-zA-Z]:\\\\(((?![<>:\"/\\\\|?*]).)+((?<![ .])\\\\)?)*$"
	////var regex = "^[a-zA-Z]:\\\\(((?![<>:\"/\\\\|?*]).)+((?<![ .])\\\\)?)*$"
	//
	//b, err := regexp.MatchString(regex, p.DirIn)
	//u.CheckErr(err)
	//if b {
	//	return fmt.Errorf("répertoire cible format incorrect :[%s]", err)
	//
	//}
	//b, err = regexp.MatchString(regex, p.DirOut)
	//u.CheckErr(err)
	//if b {
	//	return fmt.Errorf("répertoire sortie format incorrect :[%s]", err)
	//
	//}

	err = u.ValidateDir(p.DirOut, "output")
	err = u.ValidateDir(p.DirIn, "input")

	if err != nil {
		return err
	}
	return nil
}

func (u *Utils) CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (u *Utils) ValidateDir(dir, mode string) error {
	_, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("failed to find the %s directory, error: %s", mode, dir)
	}
	return nil
}
