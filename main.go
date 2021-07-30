package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	debugLevels := []string{"ERROR", "WARN", "INFO", "DEBUG"}
	logPath := flag.String("logPath", "file.log", "Path of the log file to analyze")
	debugLevel := flag.Int("debugLevel", 1, "Debug level. Valid numbers are 1=ERROR, 2=ERROR+WARN, 3=ERROR+WARN+INFO, 4=3=ERROR+WARN+INFO")
	flag.Parse()
	logFile, err := os.Open(*logPath)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	logReader := bufio.NewReader(logFile)
	for {
		logString, err := logReader.ReadString('\n')
		if err != nil {
			break
		}
		if *debugLevel >= 1 && *debugLevel <= (len(debugLevels)) {
			for i := 0; i <= (*debugLevel - 1); i++ {
				if strings.Contains(logString, debugLevels[i]) {
					fmt.Print(logString)
				}
			}
		} else {
			log.Fatal("Not a valid debugLevel")
			break
		}
	}
}
