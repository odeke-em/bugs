package main

import (
	"bufio"
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

func main() {
	var root string
	var mode string
	var outDir string
	var action string
	var curVersion int
	var oldVersion int
	flag.StringVar(&root, "i", "./temp/", "input")
	flag.StringVar(&outDir, "o", "./out/", "output")
	flag.StringVar(&mode, "m", "all", "all|code|res")
	flag.IntVar(&curVersion, "v", 1, "res version")
	flag.StringVar(&action, "action", "md5_name", "action: md5_name|md5_content|export_update")
	flag.IntVar(&oldVersion, "old_version", 1, "old version")
	flag.Parse()
	//root := "../test/pc_test/sanguohero/scripts"
	//outDir := "./out/"
	if action == "md5_name" {
		if mode == "all" || mode == "code" {
			codeMap := md5_code(root+"scripts/", outDir+"scripts/")
			writeMap(codeMap, "code_md5.txt")
		}
		if mode == "all" || mode == "res" {
			resMap := md5_res(root+"res/", outDir+"res/")
			writeMap(resMap, "res_md5.txt")
		}
	} else if action == "md5_content" {
		versionFile := fmt.Sprintf("%s/hd_res_md5_%d.txt", outDir, curVersion)
		if _, err := os.Stat(versionFile); os.IsNotExist(err) {
			outMap := make(map[string]string)
			// path/to/whatever does not exist
			md5_content(root, "/res", outMap)
			md5_content(root, "/scripts", outMap)
			writeMap(outMap, versionFile)
			println(curVersion, "版本生成完成!")
		} else {
			println(curVersion, "版本已经存在!")
		}
	} else if action == "export_update" {
		//println(oldVersion, curVersion)
		if oldVersion < curVersion {
			oldFile := fmt.Sprintf("%s/hd_res_md5_%d.txt", outDir, oldVersion)
			newFile := fmt.Sprintf("%s/hd_res_md5_%d.txt", outDir, curVersion)
			if _, err := os.Stat(oldFile); os.IsNotExist(err) {
				println(oldVersion, "旧版本不存在!")
				return
			}
			oldMap := getOldVersionMap(oldFile)
			if _, err := os.Stat(newFile); os.IsNotExist(err) {
				newMap := make(map[string]string)
				md5_content(root, "/res", newMap)
				md5_content(root, "/scripts", newMap)
				exportUpdate(oldVersion, curVersion, oldMap, newMap, root, outDir)
				writeMap(newMap, newFile)
				println(curVersion, "倒出更新包!")
			} else {
				println(curVersion, "新版本已经存在!")
			}
		} else {
			println("版本错误(old大于new):", oldVersion, curVersion)
		}
	}
}

func exportUpdate(oldVersion, newVersion int, oldMap, newMap map[string]string, root, outDir string) {
	updateDir := fmt.Sprintf("%s/update_%d_to_%d", outDir, oldVersion, newVersion)
	os.MkdirAll(updateDir, os.ModePerm)
	updateMap := make(map[string]string)
	for k, v := range newMap {
		//println(k, v)
		if oldMap[k] != v {
			updateMap[k] = v
			//println("ttttttt", oldMap[k], k)
		}
	}
	for k, _ := range updateMap {
		common.CopyFile(root+"/"+k, updateDir+"/"+k)
	}
}

func getOldVersionMap(filePath string) map[string]string {
	strMap := make(map[string]string)
	inFile, _ := os.Open(filePath)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		list := strings.Split(scanner.Text(), ",")
		if len(list) == 2 {
			pathStr := strings.TrimSpace(list[1])
			//println(pathStr)
			strMap[pathStr] = list[0]
		}
		//fmt.Println(scanner.Text())
	}
	println(len(strMap))
	return strMap
}

func md5_content(root, subDir string, outMap map[string]string) {
	curPath := root + subDir
	filepath.Walk(curPath, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		extStr := filepath.Ext(info.Name())
		if extStr != ".json" &&
			extStr != ".zip" &&
			extStr != ".mp3" &&
			extStr != ".mp4" &&
			extStr != ".txt" &&
			extStr != ".ttf" &&
			extStr != ".fnt" &&
			extStr != ".plist" &&
			extStr != ".pb" &&
			extStr != ".proto" &&
			extStr != ".lua" &&
			extStr != ".png" {
			return nil
		}
		bytes, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		md5List := md5.Sum(bytes)
		absFile := p[len(root)+1:]
		md5Str := fmt.Sprintf("%x", md5List)
		outMap[absFile] = md5Str
		return nil
	})
}

func writeMap(outMap map[string]string, path string) {
	resultStr := ""
	for k, v := range outMap {
		//println(k, v)
		resultStr += fmt.Sprintf("%s,%s \n", v, k)
	}
	resultStr += "\n"
	//println("write path:", path)
	ioutil.WriteFile(path, []byte(resultStr), os.ModePerm)
}

func md5_res(root, outDir string) map[string]string {
	root = filepath.Dir(root)
	outMap := make(map[string]string)
	var wg sync.WaitGroup
	os.RemoveAll(outDir)
	os.MkdirAll(outDir, os.ModePerm)
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		wg.Add(1)
		defer wg.Done()
		if info.IsDir() {
			return nil
		}

		extStr := filepath.Ext(info.Name())
		if extStr != ".json" &&
			extStr != ".zip" &&
			extStr != ".mp3" &&
			extStr != ".mp4" &&
			extStr != ".txt" &&
			extStr != ".ttf" &&
			extStr != ".fnt" &&
			extStr != ".plist" &&
			extStr != ".pb" &&
			extStr != ".proto" &&
			extStr != ".png" {
			return nil
		}
		//println("eeee", p, root)

		absFile := p[len(root)+1:]
		//fmt.Printf("%s, %s\n", p, root)
		md5Str := md5.Sum([]byte(strings.ToLower(absFile) + "test"))
		outMap[absFile] = fmt.Sprintf("%x%s", md5Str, extStr)
		if absFile == "ui_layout/ui/mainpage/num-vip.png" {
			//fmt.Printf("%s, %x\n", absFile, md5Str)
		}
		outFile := fmt.Sprintf("%s%x%s", outDir, md5Str, extStr)
		common.CopyFile(p, outFile)
		return nil
	})
	wg.Wait()
	println("资源导出完毕!")
	return outMap
}

func md5_code(root, outDir string) map[string]string {
	root = filepath.Dir(root)
	outMap := make(map[string]string)
	var wg sync.WaitGroup
	os.RemoveAll(outDir)
	os.MkdirAll(outDir, os.ModePerm)
	//os.MkdirAll(outDir+"deps", os.ModePerm)
	//copy deps
	//println(root+"/deps", "|", outDir+"deps/")
	//common.CopyDir(root+"/deps", outDir+"deps/")
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		wg.Add(1)
		defer wg.Done()
		if info.IsDir() {
			return nil
		}
		//bytes, err := ioutil.ReadFile(p)
		absFile := p[len(root)+1:]
		//fmt.Printf("%s, %s\n", p, root)
		md5Str := md5.Sum([]byte(strings.ToLower(absFile) + "test"))
		//fmt.Printf("%s, %x\n", absFile, md5Str)
		outFile := fmt.Sprintf("%s%x%s", outDir, md5Str, ".lua")
		outMap[absFile] = fmt.Sprintf("%x", md5Str)
		/*
			if info.Name() == "main.lua" {
				outFile = outDir + "main.lua"
			} else {
				outMap[absFile] = fmt.Sprintf("%x%s", md5Str, ".lua")
			}
		*/
		common.CopyFile(p, outFile)
		return nil
	})
	wg.Wait()
	//println(resultStr)
	println("代码导出完毕!")
	return outMap
}
