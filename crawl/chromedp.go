// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package crawl

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

// RunCrawl ...
func RunCrawl() (string, error) {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithErrorf(log.Printf))
	if err != nil {
		log.Fatal(err)
	}
	// run task list
	var res string
	err = c.Run(ctxt, text(&res))
	if err != nil {
		log.Fatal(err)
	}
	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}
	// time.Sleep(2 * time.Second)
	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("overview: %s", res)
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Print("panjang huruf :")
	log.Println(len(res))
	var isValid = false
	for i1 := 0; i1 < len(res); i1++ {

		for i2 := 0; i2 < len(letterBytes); i2++ {
			if res[i1] == letterBytes[i2] {
				isValid = true
				break
			}
		}
		if isValid {
			break
		}
	}
	if isValid == false {
		return res, errors.New("len > 0 tapi tidak mengandung string a=Z")
	}
	if res == "" {
		return res, errors.New("Not Found")
	}
	// log.Println(res)
	// d1 := []byte(res)
	// // log.Println(d1)
	// errs := ioutil.WriteFile("fb.html", d1, 0644)
	// if errs != nil {
	// 	log.Fatal(errs)
	// }
	return res, nil

}

func text(res *string) chromedp.Tasks {
	randomId := time.Now().Nanosecond()
	id := strconv.Itoa(randomId)
	return chromedp.Tasks{
		// chromedp.Navigate(`https://golang.org/pkg/time/`),
		chromedp.Navigate(`https://www.facebook.com/groups/532806206744264?tes=` + id),
		// chromedp.Text(`content_container`, res, chromedp.NodeVisible, chromedp.ByID),
		chromedp.WaitVisible(`#content_container`),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			log.Printf(">>>>>>>>>>>>>>>>>>>> content_container IS VISIBLE")
			return nil
		}),
		chromedp.Text(`content_container`, res, chromedp.NodeVisible, chromedp.ByID),

		// chromedp.
		// chromedp.Text(`feed`, res, chromedp.NodeVisible, chromedp.BySearch),
	}
}
