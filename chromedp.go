package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

// 使用chromedp处理JavaScript渲染
func renderPage(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#content`), // 等待关键元素加载
		chromedp.OuterHTML(`html`, &html),
	)
	if err != nil {
		log.Printf("渲染失败: %v", err)
	}
	return html
}
