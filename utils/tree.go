package tree

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
)

// type FileInfo struct {
// 	name         string
// 	size         int64
// 	isDir        bool
// 	isLast       bool
// 	parentIsLast bool
// 	level        int
// }

type FileMeta struct {
	file         fs.FileInfo
	isLast       bool
	parentIsLast bool
	level        int
}

func DirTree(out *bytes.Buffer, path string, printFiles bool, parentIsLast bool, level int) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	// filesInLevel := make([]FileInfo, 0, 5)
	filesInLevel := make([]FileMeta, 0, 5)

	for _, file := range files {
		if !printFiles && !file.IsDir() {
			continue
		}

		fileInfo := FileMeta{file, false, parentIsLast, level}
		filesInLevel = append(filesInLevel, fileInfo)
	}

	// for i, file := range files {
	// 	if !printFiles && !file.IsDir() {

	// 		continue
	// 	}

	// 	isLast := i == len(files)-1

	// 	fileInfo := FileInfo{file.Name(), file.Size(), file.IsDir(), isLast, parentIsLast, level}
	// 	filesInLevel = append(filesInLevel, fileInfo)
	// }

	// sort.Slice(filesInLevel, func(i, j int) bool {
	// 	return filesInLevel[j].name > filesInLevel[i].name
	// })

	// for _, v := range filesInLevel {
	// 	fmt.Println(getGraphicLine(v))
	// 	out.Write([]byte(getGraphicLine(v) + "\n"))
	// 	if v.isDir {

	// 		f := v.parentIsLast && v.isLast

	// 		DirTree(out, path+"/"+v.name, printFiles, f, level+1)
	// 	}
	// }

	for i, v := range filesInLevel {
		isLast := i == len(filesInLevel)-1
		v.isLast = isLast
		f := v.parentIsLast && isLast
		v.parentIsLast = f

		fmt.Println(getGraphicLine(v))
		out.Write([]byte(getGraphicLine(v) + "\n"))

		if v.file.IsDir() {
			DirTree(out, path+"/"+v.file.Name(), printFiles, isLast, level+1)
		}
	}

	return nil
}

func getGraphicLine(fileInfo FileMeta) string {

	var tabStr string
	symbol := "├───"

	for i := 0; i < fileInfo.level; i++ {

		// if fileInfo.level > 1 {
		// 	tabStr = tabStr + "│\t"
		// 	continue
		// }

		if fileInfo.isLast && fileInfo.parentIsLast {
			tabStr = tabStr + "\t"
			continue
		}

		tabStr = tabStr + "│\t"
	}

	if fileInfo.isLast {
		symbol = "└───"
	}

	if fileInfo.file.IsDir() {
		return fmt.Sprintf("%s%s%s %t", tabStr, symbol, fileInfo.file.Name(), fileInfo.isLast)
	}

	return fmt.Sprintf("%s%s%s (%d) %t", tabStr, symbol, fileInfo.file.Name(), fileInfo.file.Size(), fileInfo.isLast)
}

// func getGraphicLine(fileInfo FileInfo) string {

// 	var tabStr string
// 	symbol := "├───"

// 	for i := 0; i < fileInfo.level; i++ {

// 		// if fileInfo.level > 1 {
// 		// 	tabStr = tabStr + "│\t"
// 		// 	continue
// 		// }

// 		if fileInfo.isLast && fileInfo.parentIsLast {
// 			tabStr = tabStr + "1\t"
// 			continue
// 		}

// 		tabStr = tabStr + "│\t"
// 	}

// 	if fileInfo.isLast {
// 		symbol = "└───"
// 	}

// 	if fileInfo.isDir {
// 		return fmt.Sprintf("%s%s%s %t", tabStr, symbol, fileInfo.name, fileInfo.isLast)
// 	}

// 	return fmt.Sprintf("%s%s%s (%d) %t", tabStr, symbol, fileInfo.name, fileInfo.size, fileInfo.isLast)
// }
