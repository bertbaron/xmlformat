package main

// A simple streaming XML formatter.
// Not as fast and versatile as xmllint, but doesn't run into memory issues for very big files.
//
// All whitespace-only tokens are assumed to be 'ignorable'. It is not possible to use a schema.

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
)

var indent = flag.String("indent", "tab", "the indentation string, or 'tab' for indenting with a tab character")
var outfile = flag.String("outfile", "", "the output file, leave blank to write to stdout")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: xmlformat [filename]\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
}

// all whitespace-only tokens are assumed to be ignorable
func ignorable(t xml.Token) bool {
	switch ele := t.(type) {
	case xml.CharData:
		for _, b := range ele {
			if b != 0x20 && b != 0x9 && b != 0xD && b != 0xA {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		flag.Usage()
		os.Exit(1)
	}

	xmlReader := os.Stdin
	if len(args) == 1 {
		inputFile := args[0]
		xmlFile, err := os.Open(inputFile)
		if err != nil {
			log.Fatal("Error opening input file:", err)
		}
		defer xmlFile.Close()
		xmlReader = xmlFile
	}
	decoder := xml.NewDecoder(xmlReader)
	decoder.CharsetReader = charset.NewReader

	out := os.Stdout
	if *outfile != "" {
		outFile, err := os.Create(*outfile)
		if err != nil {
			log.Fatal("Error opening output file:", err)
		}
		defer outFile.Close()
		out = outFile
	}
	encoder := xml.NewEncoder(out)
	if *indent == "tab" {
		*indent = "\t"
	}
	encoder.Indent("", *indent)

	for {
		t, err := decoder.Token()
		if t == nil {
			break
		}
		if err != nil {
			log.Fatalf("Failed to parse xml: %v", err)
		}
		if !ignorable(t) {
			switch t.(type) {
			default:
				encoder.EncodeToken(t)
			case xml.ProcInst:
				procInst := xml.ProcInst{"xml", []byte("version=\"1.0\" encoding=\"UTF-8\"")}
				encoder.EncodeToken(procInst)
				// For some reason the encoder does not write a newline after the ProcInst <?xml...>
				// We fix this by inserting a newline. This looks nice with go 1.5, but with 1.4
				// it will write an escaped newline...
				encoder.EncodeToken(xml.CharData([]byte{10}))
			}
		}
	}
	encoder.Flush()
	fmt.Fprintln(out, "")
}
