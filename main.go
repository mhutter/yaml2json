package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mhutter/yaml2json/convert"
)

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), `
Convert YAML to JSON

USAGE:
  $ yaml2json [<flags>] [infile [outfile]]

ARGUMENTS:
    infile
        File to read from (defaults to STDIN)

    outfile
        File to write to (defaults to STDOUT)

FLAGS:`)
	flag.PrintDefaults()
}

var (
	infile      = os.Stdin
	outfile     = os.Stdout
	pretty      = false
	versionFlag = false
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.CommandLine.Usage = usage
	flag.BoolVar(&pretty, "p", pretty, "Pretty-print (indent) output")
	flag.BoolVar(&versionFlag, "v", versionFlag, "Print version and exit")
	flag.Parse()
}

func main() {
	if versionFlag {
		printVersion()
		os.Exit(0)
	}

	var err error
	if flag.NArg() > 0 {
		infile, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalln(err)
		}
	}
	if flag.NArg() > 1 {
		outfile, err = os.OpenFile(
			flag.Arg(1),
			// Open write-only, create file if it does not exist, truncate
			// content if it DOES exist
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0644,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if err := convert.YAML2JSON(infile, outfile, pretty); err != nil {
		log.Fatalln("Error encoding data:", err)
	}
}
