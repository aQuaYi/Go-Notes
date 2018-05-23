// Command simple is a chromedp example demonstrating how to do a simple google
// search.
package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var function string
	err = c.Run(ctxt, getFunction(&function))
	if err != nil {
		log.Println("Run error")
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Println("Shutdown error")
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Println("Wait error")
		log.Fatal(err)
	}

	log.Println(function)
}

func getFunction(function *string) chromedp.Tasks {
	textarea := `//textarea`
	btn := `#question-detail-app > div > div:nth-child(3) > div > div > div.row.control-btn-bar > div > div > div > div > span.Select-arrow-zone`
	goSel := `#react-select-2--option-9`
	return chromedp.Tasks{
		// chromedp.Navigate(`https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/`),
		chromedp.Navigate(`https://leetcode.com/problems/k-th-symbol-in-grammar/description/`),
		chromedp.Click(btn, chromedp.ByID),
		chromedp.Click(goSel, chromedp.ByID),
		chromedp.Text(textarea, function),
	}
}
