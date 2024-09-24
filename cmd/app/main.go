package main

import (
	"github.com/sss7526/webshooter/internal/processor"
	"github.com/sss7526/webshooter/internal/cli"
)

func main() {
	parsedArgs := cli.ParseArgs()

	targets := parsedArgs["targets"].([]string)

	saveToImage, ok := parsedArgs["image"].(bool)
	if !ok && !saveToImage {
		saveToImage = false
	} else {
		saveToImage = true
	}

	saveToPDF, ok := parsedArgs["pdf"].(bool)
	if !ok && !saveToPDF {
		saveToPDF = false
	} else {
		saveToPDF = true
	}

	verbose, ok := parsedArgs["verbose"].(bool)
	if !ok && !verbose {
		verbose = false
	} else {
		verbose = true
	}

	translate, ok := parsedArgs["translate"].(bool)
	if !ok && !translate {
		translate = false
	}

	if len(targets) > 0 {
		processor.ProcessTargets(targets, verbose, saveToImage, saveToPDF, translate)
	}
}
