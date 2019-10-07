package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	//"path"
	"common"
	"flag"
	"path/filepath"
	"strings"
	"sync"
)

func Create_md5(root, outDir string) {
	outMap := make(map[string]string)
	var wg sync.WaitGroup
	os.RemoveAll(outDir)
	os.MkdirAll(outDir, os.ModePerm)
	os.MkdirAll(outDir+"deps", os.ModePerm)
	//copy deps
	println(root+"/deps", "|", outDir+"deps/")
	common.CopyDir(root+"/deps", outDir+"deps/")
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
		common.CopyFile(p, outFile)
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

func main() {
	var root string
	flag.StringVar(&root, "i", "./temp/", "导出")
	flag.Parse()
	//root := "../test/pc_test/sanguohero/scripts"
	outDir := "./out/"
	Create_md5(root+"scripts/", outDir+"scripts/")
}
