package common

import (
	"fileManager2/pkg/models"
	"os"
	"sync"
)

type IFileManager interface {
	CopyOrMove(src, dst string, dt *models.DataTemplate, wgc *sync.WaitGroup, s chan *models.Stack)
	CheckExtAndCopy(entry os.FileInfo, dt *models.DataTemplate, wgp *sync.WaitGroup, s chan *models.Stack)
	StartProcessing(dt *models.DataTemplate, s chan *models.Stack)
	Process(fi []os.FileInfo, dt *models.DataTemplate, wg *sync.WaitGroup, s chan *models.Stack)
	MakeOutDirs(exts []string, dirOut string, wg *sync.WaitGroup)
}
