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

	_, err := os.Stat(p.DirIn)
	if err != nil {
		return fmt.Errorf("failed to find the input directory, error: %w : %s", err, p.DirIn)
	}
	_, err = os.Stat(p.DirOut)
	if err != nil {
		return fmt.Errorf("failed to find the output directory, error: %w : %s", err, p.DirOut)
	}

	return nil
}

func (u *Utils) CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
