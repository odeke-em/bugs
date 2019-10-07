package main

import (
	"debug/dwarf"
	"debug/macho"
	"encoding/json"
	"log"
)

func main() {
	ff, err := macho.Open("all1")
	if err != nil {
		log.Fatalf("Failed to open up binary: %v", err)
	}
	defer ff.Close()

	dd, err := ff.DWARF()
	if err != nil {
		log.Fatalf("Failed to get DWARF data: %v", err)
	}

	var entries []*dwarf.Entry
	var names []string
	dr := dd.Reader()
	for {
		entry, err := dr.Next()
		if err != nil {
			log.Fatal(err)
		}
		if entry == nil {
			break
		}

		if entry.Tag != dwarf.TagVariable {
			continue
		}

		name, ok := entry.Val(dwarf.AttrName).(string)
		if !ok {
			continue
		}
		entries = append(entries, entry)
		names = append(names, name)
	}
	blob, _ := json.MarshalIndent(entries, "", "  ")
	println(string(blob))
	if false {
		blob, _ = json.MarshalIndent(names, "", "  ")
		println(string(blob))
	}
}
