package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	//"path"
	"common"
	"flag"
	"os/exec"
	"path/filepath"
	//"strings"
	//"sync"
)

func compress(root, outDir, tooDir, libDir string) {
	//outMap := make(map[string]string)
	//toolPath := filepath.Dir(os.Args[0])
	//println(toolPath)
	//copy deps
	os.MkdirAll(libDir, os.ModePerm)
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		//println(p)
		curP := p[len(root):]
		//println("aa:", curP)
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".png" {
			return nil
		}
		//os.Exit(0)
		// get File md5 name
		bytes, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		md5List := md5.Sum(bytes)
		md5Str := fmt.Sprintf("%x", md5List)

		outFile := libDir + "/" + md5Str + "_" + info.Name()
		if _, err := os.Stat(outFile); os.IsNotExist(err) {
			c := exec.Command(tooDir+"/pngquant", "--force", p, "--output", p)
			//c := exec.Command(tooDir+"/pngquant", "--force", "--",p)
			//c := exec.Command("ls", "-hl", p)
			d, err := c.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				//os.Exit(0)
				return nil
			}
			//println("okkkk")
			fmt.Println(string(d))
			common.CopyFile(p, outFile)
		} else {
			//fmt.Println("已经存在")
			common.CopyFile(outFile, p)
		}

		println("compress:" + p)
		//fmt.Printf("pngquant --output %s", p)
		//os.Exit(0)
		return nil
	})
	//println("ok")
	/*
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
	*/
	println("导出完毕!")
}

func main() {
	var root string
	var outDir string
	var toolDir string
	var picLibPath string
	flag.StringVar(&root, "i", "./temp/res_src/", "input")
	flag.StringVar(&outDir, "o", "./temp/res_src/", "output")
	flag.StringVar(&toolDir, "t", ".", "tool")
	flag.StringVar(&picLibPath, "lib", ".", "pic lib")
	flag.Parse()
	//root := "../test/pc_test/sanguohero/scripts"
	compress(root, outDir, toolDir, picLibPath)
}
