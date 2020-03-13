package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type arguments struct {
	regexp string
	output string
	knm    bool
}

func parseFlags() arguments {
	regularExpression := flag.String("r", `(.*)`, "PCRE Regular expression you want to match")
	outputFormat := flag.String("o", `<${0}>`, "Output format. $n or ${n} represents the n-th group matched")
	keepNotMatch := flag.Bool("knm", false, "If true keeps in the output the lines not matched")
	flag.Parse()
	return arguments{regexp: *regularExpression, output: *outputFormat, knm: *keepNotMatch}
}

func getParsedRegularExpression(rexp string) *regexp.Regexp {
	r, err := regexp.Compile(rexp)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func main() {
	args := parseFlags()

	reader := bufio.NewReader(os.Stdin)
	r := getParsedRegularExpression(args.regexp)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !r.MatchString(str) && !args.knm {
			continue
		}
		res := r.ReplaceAllString(str, args.output)
		fmt.Print(res)

	}
}
