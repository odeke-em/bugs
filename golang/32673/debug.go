package main

import (
	"debug/dwarf"
	"debug/macho"
	"flag"
	"log"
	"os"
)

func main() {
	srcFile := flag.String("file", "", "The file to examine")
	flag.Parse()
	log.SetFlags(0)

	f, err := os.Open(*srcFile)
	if err != nil {
		log.Fatalf("Failed to open the subject file: %v", err)
	}
	defer f.Close()
	mf, err := macho.NewFile(f)
	if err != nil {
		log.Fatalf("Failed to read the Macho binary: %v", err)
	}
	dwd, err := mf.DWARF()
	if err != nil {
		log.Fatalf("Failed to get the DWARF debug information: %v", err)
	}
	rd := dwd.Reader()

	for {
		entry, err := rd.Next()
		if entry == nil {
			break
		}
		if err != nil {
			log.Printf("Got error: %v", err)
		}
		lnr, err := dwd.LineReader(entry)
		if err != nil {
			log.Printf("LineReader error: %v", err)
			continue
		}
		if lnr == nil {
			continue
		}
		lnPos := lnr.Tell()
                declFile := entry.AttrField(dwarf.AttrDeclFile)
                if false && declFile == nil {
                    continue
                }
                declLine := entry.AttrField(dwarf.AttrDeclLine)
                log.Printf("%v: %v:%v %#v\n\n", entry.AttrField(dwarf.AttrName), declFile, declLine, lnPos)
	}

	if false {
		syms := mf.Symtab.Syms
		for _, sym := range syms {
			log.Printf("Symbol: %#v", sym)
		}
	}
}
