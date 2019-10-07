package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	//"path"
	"io"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	root := "../test/pc_test/sanguohero/scripts"
	outDir := "./out/"
	outMap := make(map[string]string)
	var wg sync.WaitGroup
	os.RemoveAll(outDir)
	os.MkdirAll(outDir, os.ModePerm)
	os.MkdirAll(outDir+"deps", os.ModePerm)
	//copy deps
	CopyDir(root+"/deps", outDir+"deps/")
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		wg.Add(1)
		defer wg.Done()
		if info.IsDir() {
			return nil
		}
		if info.Name() == "re.lua" {
			return nil
		}
		if filepath.Ext(info.Name()) != ".lua" {
			return nil
		}
		if strings.Index(p, "deps") != -1 {
			//println("dddd:", p)
			return nil
		}

		bytes, err := ioutil.ReadFile(p)
		md5Str := md5.Sum(bytes)
		//fmt.Printf("%x\n", md5Str)
		outFile := fmt.Sprintf("%s%x%s", outDir, md5Str, ".lua")
		if info.Name() == "main.lua" {
			outFile = outDir + "main.lua"
		} else {
			outMap[p] = outFile
		}
		CopyFile(p, outFile)
		return nil
	})
	wg.Wait()
	//println("ok")
	resultStr := "local resTable = {\n"
	for k, v := range outMap {
		kStr := k[len(root)+1 : len(k)-4]
		kStr = strings.Replace(kStr, "/", ".", -1)
		resultStr += fmt.Sprintf("[\"%s\"] = \"%s\", \n", kStr, v[len(outDir):len(v)-4])
	}
	resultStr += "}\n"
	resultStr += "return resTable \n"
	ioutil.WriteFile(outDir+"res.lua", []byte(resultStr), os.ModePerm)
	//println(resultStr)
	println("导出完毕!")
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create
			// sub-directories
			// -
			// recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform
			// copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
