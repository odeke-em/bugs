package main

import (
	"bufio"
	"common"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	//"path"
	"flag"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func GenNew(root, outDir string, version int) map[string]string {
	//println(flag.Args())
	//root := "../../test/pc_test/sanguohero/"
	//outDir := "./scripts/"
	//outDir := "./"
	root = filepath.Dir(root)
	outMap := make(map[string]string)
	var wg sync.WaitGroup
	filepath.Walk(root+"/scripts", func(p string, info os.FileInfo, err error) error {
		wg.Add(1)
		defer wg.Done()
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".lua" {
			return nil
		}
		bytes, err := ioutil.ReadFile(p)
		md5Str := md5.Sum(bytes)
		//println(p, " ", root)
		outMap[p[len(root)+1:]] = fmt.Sprintf("%x", md5Str)
		return nil
	})
	wg.Wait()
	filepath.Walk(root+"/res", func(p string, info os.FileInfo, err error) error {
		wg.Add(1)
		defer wg.Done()
		if info.IsDir() {
			return nil
		}
		bytes, err := ioutil.ReadFile(p)
		md5Str := md5.Sum(bytes)
		outMap[p[len(root)+1:]] = fmt.Sprintf("%x", md5Str)
		return nil
	})
	//println("ok")
	return outMap
}

func main() {
	var version int
	var mode int
	var root string
	flag.IntVar(&version, "version", 0, "版本号")
	flag.IntVar(&mode, "mode", 0, "0导出全部版本，1导出差异版本")
	flag.StringVar(&root, "i", "./temp/", "input")
	flag.Parse()
	//root := "../../test/pc_test/sanguohero/"
	outDir := "./"
	outMap := GenNew(root, outDir, version)
	if mode == 0 {
		resultStr := ""
		for k, v := range outMap {
			//kStr := k[len(root):]
			//kStr = strings.Replace(kStr, "/", ".", -1)
			resultStr += fmt.Sprintf("%s,%s\n", k, v)
		}
		//println("version:", version)
		ioutil.WriteFile(outDir+"res_"+strconv.Itoa(version)+".pack", []byte(resultStr), os.ModePerm)
		//println(resultStr)
		println("导出完毕!")
	} else {
		oldMap := make(map[string]string)
		versionFile := outDir + "res_" + strconv.Itoa(version) + ".pack"
		f, err := os.OpenFile(versionFile, os.O_RDONLY, 0)
		if err != nil {
			println("版本不存在!")
			os.Exit(1)
		}
		br := bufio.NewReader(f)
		lineNum := 0
		for {
			//每次读取一行
			line, err := br.ReadString('\n')
			lineNum++
			if err == io.EOF {
				break
			} else {
				strList := strings.Split(line, ",")
				if len(strList) != 2 {
					println("版本文件错误:", lineNum)
					return
				}
				md5Str := strList[1]
				//fmt.Printf("%v", md5Str[0:len(md5Str)-1])
				oldMap[strList[0]] = md5Str[0 : len(md5Str)-1]
				//fmt.Printf("%v", line)
			}
		}
		oldDele := make(map[string]string)
		for k, v := range oldMap {
			//print(k, "|", outMap[k], "|", v, "|", strings.EqualFold(outMap[k], v))
			if outMap[k] == v {
				delete(outMap, k)
			} else {
				oldDele[k] = v
			}
		}
		for k, _ := range outMap {
			common.CopyFile(root+k, outDir+"/update/"+k)
		}
		println("导出update完毕")
	}
}
