// 爬蟲(大樂透)
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly/v2"
)

var __VIEWSTATEGENERATOR string
var __EVENTVALIDATION string
var __VIEWSTATE string

func main() {
	var url = "https://www.taiwanlottery.com.tw/3th_Lotto/Lotto649/history.aspx" // 爬取目標網址
	c := colly.NewCollector()                                                    // 爬蟲套件

	c.OnHTML("input[name='__VIEWSTATEGENERATOR']", func(e *colly.HTMLElement) {
		__VIEWSTATEGENERATOR = e.Attr("value") // 網頁標籤內容
		// fmt.Println(__VIEWSTATEGENERATOR)
	})

	c.OnHTML("input[name='__EVENTVALIDATION']", func(e *colly.HTMLElement) {
		__EVENTVALIDATION = e.Attr("value")
		// fmt.Println(__EVENTVALIDATION)
	})

	c.OnHTML("input[name='__VIEWSTATE']", func(e *colly.HTMLElement) {
		__VIEWSTATE = e.Attr("value")
		//fmt.Println(__VIEWSTATE)
	})

	c.Visit(url) // 訪問網址

	// 帶入Header資料
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Chrome/84.0.4147.89 Safari/537.36")
	})

	// 整個body內容
	c.OnScraped(func(r *colly.Response) {
		// fmt.Println(string(r.Body))
	})

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo1_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo2_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo3_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo4_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo5_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo6_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnHTML("span#Lotto649Control_history1_dlQuery_No7_0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.OnHTML("span#Lotto649Control_history1_dlQuery_L649_DrawTerm_0", func(e *colly.HTMLElement) {
		fmt.Print("期別:")
		fmt.Print(e.Text)
	})

	c.OnHTML("span#Lotto649Control_history1_dlQuery_L649_DDate_0", func(e *colly.HTMLElement) {
		fmt.Print("  開獎日期:")
		fmt.Print(e.Text)
	})

	c.OnHTML("span#Lotto649Control_history1_dlQuery_L649_EDate_0", func(e *colly.HTMLElement) {
		fmt.Print("  兌獎截止:")
		fmt.Print(e.Text)
	})

	fmt.Println()
	fmt.Print("開獎號碼:")

	// 爬取網頁標籤內容
	for i := 1; i <= 7; i++ {
		c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo"+strconv.Itoa(i)+"_0", func(e *colly.HTMLElement) {
			fmt.Print(e.Text, ",")
		})
	}

	c.OnHTML("span#Lotto649Control_history1_dlQuery_SNo_0", func(e *colly.HTMLElement) {
		fmt.Print("  特別號:")
		fmt.Print(e.Text)
	})

	// post要帶入的資料
	var formData = map[string]string{
		"__EVENTTARGET":                          "",
		"__EVENTARGUMENT":                        "",
		"__LASTFOCUS":                            "",
		"__VIEWSTATE":                            __VIEWSTATE,
		"__VIEWSTATEGENERATOR":                   __VIEWSTATEGENERATOR,
		"__VIEWSTATEENCRYPTED":                   "",
		"__EVENTVALIDATION":                      __EVENTVALIDATION,
		"Lotto649Control_history1$DropDownList1": "2",
		"Lotto649Control_history1$chk":           "radYM",
		"Lotto649Control_history1$dropYear":      "102", // 中華民國年分
		"Lotto649Control_history1$dropMonth":     "12",  // 月份
		"Lotto649Control_history1$btnSubmit":     "查詢",
	}

	if err := c.Post(url, formData); err != nil { // 進到該url 執行POST
		log.Println(err)
	}
}
