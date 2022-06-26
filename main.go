package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	ctx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	// run task list
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.baidu.com`),
		chromedp.SendKeys(`kw`, "杨明明", chromedp.ByID),
		chromedp.Click(`su`, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
