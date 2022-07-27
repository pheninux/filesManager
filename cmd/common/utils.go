package common

import (
	"errors"
	"fileManager2/pkg/models"
	"fmt"
	"os"
	"path/filepath"
)

type Utils struct {
	Slash string
}

func (u *Utils) ValidateArgs(dt *models.DataTemplate, s chan *models.Stack) error {
	var err error
	//var regex = "[a-zA-Z]:\\\\(((?![<>:\"/\\\\|?*]).)+((?<![ .])\\\\)?)*$"
	////var regex = "^[a-zA-Z]:\\\\(((?![<>:\"/\\\\|?*]).)+((?<![ .])\\\\)?)*$"
	//
	//b, err := regexp.MatchString(regex, dt.DirIn)
	//u.CheckErr(err)
	//if b {
	//	return fmt.Errorf("répertoire cible format incorrect :[%s]", err)
	//
	//}
	//b, err = regexp.MatchString(regex, dt.DirOut)
	//u.CheckErr(err)
	//if b {
	//	return fmt.Errorf("répertoire sortie format incorrect :[%s]", err)
	//
	//}

	fmt.Println(dt)
	err = u.ValidateDir(dt.DirOut, "output")
	if err != nil {
		//fmt.Println(err)
		dt.Stack.Err = err.Error()
		s <- dt.Stack
		return err
	}
	err = u.ValidateDir(dt.DirIn, "input")
	if err != nil {
		//fmt.Println(err)
		dt.Stack.Err = err.Error()
		s <- dt.Stack
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
		return fmt.Errorf("failed to find the %s \n directory : %s", mode, dir)
	}
	return nil
}

func CountFileWithExtention(dt *models.DataTemplate, s chan *models.Stack) {
	var filesCheck []string

	dt.Stack.Ffound = map[string]int{}
	for _, e := range dt.Exts {
		pattern := filepath.Join(dt.DirIn, "*."+e)
		dt.Stack.Pattern = append(dt.Stack.Pattern, pattern)
		files, err := filepath.Glob(pattern)
		filesCheck = append(filesCheck, files...)
		if err == nil {
			//fmt.Printf("Found %d files.%s .\n", len(files), e)
			dt.Stack.Ffound[e] = len(files)

		}
	}
	if len(filesCheck) == 0 {
		dt.Stack.Err = errors.New("no files found with \nthese extensions").Error()
	}
	s <- dt.Stack
}
