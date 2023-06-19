// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"

	"regexp"
	"fmt"
	"strings"
    "net/http"
)

var htmlString string
var htmlString1 string

func main() {

    
	

	// // write htmlString to a new file
	// f, err := os.Create("output1.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// f.WriteString(htmlString1)


    handleRequests()

}

func scrapeNparks() {
    fmt.Println("Scraping Nparks")
    // create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.nparks.gov.sg/activities/events-and-workshops`),

		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`.listing-content`, chromedp.ByQuery),
		chromedp.OuterHTML(`.listing-content`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(strings.TrimSpace(res + "result"))
    htmlString1 = `
    <head>
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.13.0/css/all.css">
    </head>
    <style>
        .event {
            display: flex;
        }
        .summary-wrapper-event {
            width: 100%;
            display: flex;
            flex-direction: column;
            padding: 20px;
            margin-left: 20px;
        }

        .thumbnail-event {
            width: 30%;
        }
        .listing-content img {
            width: 100%;
            max-height: 200px;
            padding: 20px;
        }
        .listing-content ul {
            padding-left: 0px;
        }
        li.event {
            list-style: none;
            border: 1px solid #ccc;
        }
        .icon-date:before {
            font-family: "Font Awesome 5 Free"; /* This is the correct font-family*/
            font-style: normal;
            font-size:25px;
            margin-right: 10px;
            margin-left: 10px;
            margin-top: 0;
            content: "";
        }

        .icon-location:before {
            font-family: "Font Awesome 5 Free"; /* This is the correct font-family*/
            font-style: normal;
            font-size:35px;
            margin-right: 10px;
            margin-left: 12px;
            margin-top: 0;
            content: "⚲";
        }

        .title-event a {
            color: #00833f;
        }
    </style>
    `
    htmlString1 += strings.ReplaceAll(res, `href="`, `target="_blank" href="https://www.nparks.gov.sg`)
    htmlString1 = strings.ReplaceAll(htmlString1, `src="`, `src="https://www.nparks.gov.sg/`)
}

func scrapeGBTB() {
    // create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.gardensbythebay.com.sg/en/things-to-do/calendar-of-events.html`),

		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`.gbb-location`, chromedp.ByQuery),
		chromedp.OuterHTML(`.programme-tiles`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(strings.TrimSpace(res + "result"))

	/*
		<div><a target="_self" href="/en/things-to-do/calendar-of-events/rose-romance-2023.html" class="programme-tile row-listing-tile a--tile small-image-tile"><div class="programme-tile__image"><img src="/content/dam/gbb-2021/image/things-to-do/events/rose-romance-2023/RoseRomance2023-Tile-Thumbnail-square-1080x1080.jpg" alt="Rose Romance 2023" class="object-fit-image"> <!----> <!----> <span class="tile--credit"></span></div> <div class="programme-tile__content"><div class="title"><h6>Rose Romance 2023</h6></div> <div class="detail">The Queen of Flowers takes pride of place in Flower Dome once more with the return of Rose Romance, presented in collaboration with the Embassy of Italy in Singapore.</div> <div class="info"><p><span class="gbb-event">Fri, 2 Jun 2023 - Sun, 16 Jul 2023</span></p>
<p><span class="gbb-hours">9:00am - 9:00pm</span></p>
<p><span class="gbb-location">Flower Dome</span></p>
</div></div></a></div>
	*/

	// <div><a target="_self"??????????????????????</a></div>
	regex := regexp.MustCompile(`(?s)<div><a target="_self".*?</a></div>`)
	// get all matches
	matches := regex.FindAllStringSubmatch(res, -1)

	htmlString = `
	<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1.0">

<head>
<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.13.0/css/all.css">
</head>

<style>
    
    a {
        text-decoration: none;
        /* dont change color on click */
        color: inherit;
    }

    .title {
        font-family: Lato,Arial,sans-serif;
        -webkit-font-smoothing: subpixel-antialiased;
        font-weight: 400;
        font-size: 30px;
    }
    


    .gbb-hours:before {
        font-family: "Font Awesome 5 Free"; /* This is the correct font-family*/
        font-style: normal;
        font-size:25px;
        margin-right: 10px;
        margin-left: 10px;
        margin-top: 0;
        content: "";
    }

    .gbb-event:before {
        font-family: "Font Awesome 5 Free"; /* This is the correct font-family*/
        font-style: normal;
        font-size:25px;
        margin-right: 10px;
        margin-left: 10px;
        margin-top: 0;
        content: "";
    }

    .gbb-location:before {
        font-family: "Font Awesome 5 Free"; /* This is the correct font-family*/
        font-style: normal;
        font-size:35px;
        margin-right: 10px;
        margin-left: 12px;
        margin-top: 0;
        content: "⚲";
    }

    .programme-tile {
        /* flex */
        display: flex;
        flex-direction: row;
        cursor: pointer;
        border: 1px solid #000000;
        margin-top: 40px;
        font: 12px/1.5 "Helvetica Neue", Helvetica, Arial, sans-serif;
    }

    .programme-tile__content {
        padding: 20px;
    }

    .programme-tile__image {
        max-width: 25%;
    }

    .object-fit-image {
        width: 100%;
        height: 100%;
        -o-object-fit: cover;
        object-fit: cover;
        font-family: "object-fit: cover;";
    }

    
</style>
`
	baseURL := "https://www.gardensbythebay.com.sg"

	// iterate through matches
	for _, match := range matches {
		// fmt.Println("\n\nmatch: ")
		// fmt.Println(match[0])
		// add baseURL to href and img src
		matchcopy := match[0]
		matchcopy = strings.Replace(matchcopy, `href="`, `href="` + baseURL, -1)
		matchcopy = strings.Replace(matchcopy, `src="`, `src="` + baseURL, -1)

		htmlString += matchcopy
	}
}

func gbtbPage(w http.ResponseWriter, r *http.Request){
    enableCors(&w)
    scrapeGBTB()
    
    w.Write([]byte(htmlString))
    // fmt.Println(htmlString)
    fmt.Println("Endpoint Hit: homePage")
}

func nparksPage(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    scrapeNparks()

    w.Write([]byte(htmlString1))

}

func handleRequests() {
    http.HandleFunc("/gbtb", gbtbPage)
    http.HandleFunc("/nparks", nparksPage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    }