package main

import (
	"common"
	"fmt"
	"io/ioutil"
	"os"
	//"path"
	"bytes"
	"encoding/json"
	"flag"
	"github.com/xxtea"
	"path/filepath"
	"sync"
)

func encrypt_res(root, outDir, key string) {
	//root = filepath.Dir(root)
	var wg sync.WaitGroup
	os.RemoveAll(outDir)
	dst := new(bytes.Buffer)
	os.MkdirAll(outDir, os.ModePerm)
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		wg.Add(1)
		defer wg.Done()
		if info.IsDir() {
			return nil
		}
		extStr := filepath.Ext(info.Name())
		fileName := p[len(root):]
		outFile := fmt.Sprintf("%s%s", outDir, fileName)
		if extStr != ".json" &&
			extStr != ".proto" &&
			extStr != ".pb" &&
			extStr != ".plist" &&
			//extStr != ".txt" &&
			extStr != ".png" {
			//println(p, outFile)
			common.CopyFile(p, outFile)
			return nil
		}

		bytes, err := ioutil.ReadFile(p)
		if err != nil {
			println("read error", p)
			return nil
		}
		if extStr == ".json" {
			dst.Reset()
			err = json.Compact(dst, bytes)
			if err != nil {
				println("json compact error", p)
				return nil
			}
			bytes = dst.Bytes()
		}
		//fmt.Printf("%x\n", md5Str)
		encrypt_data := xxtea.Encrypt(bytes, []byte(key))
		dst.Reset()
		dst.Write([]byte("ZY"))
		dst.Write(encrypt_data)
		encrypt_data = dst.Bytes()
		println(outFile)
		//print(encrypt_data)
		os.MkdirAll(filepath.Dir(outFile), os.ModePerm)
		//encrypt_data = append([]byte("ZY"), encrypt_data)
		ioutil.WriteFile(outFile, encrypt_data, os.ModePerm)
		//common.CopyFile(p, outFile)
		return nil
	})
	wg.Wait()
	//println("ok")
	println("导出完毕!")
}

func main() {
	var root string
	var key string
	var outDir string
	flag.StringVar(&root, "i", "../test/pc_test/sanguohero/", "input")
	flag.StringVar(&outDir, "o", "temp2", "output")
	flag.StringVar(&key, "k", "6d2a4052c2", "key")
	flag.Parse()
	//root := "../test/pc_test/sanguohero/scripts"
	//outDir := "./temp/"
	println(root)
	encrypt_res(root, outDir, key)
}
