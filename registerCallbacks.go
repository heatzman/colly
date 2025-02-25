package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"strings"
	//"github.com/gocolly/colly/v2/proxy"
)

func registerCallbacks(c *colly.Collector) {
	// 添加请求前回调
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "*.douban.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("origin", "")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")

		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) { fmt.Println("error:", err) })
	movie := Movie{}

	// 提取标题
	c.OnHTML(`h1 span[property="v:itemreviewed"]`, func(e *colly.HTMLElement) {
		//movie := e.Request.Ctx.GetAny("movie").(*Movie)
		movie.Title = e.Text
		fmt.Println("标题:", movie.Title)
	})

	// 提取年份
	c.OnHTML(`.year`, func(e *colly.HTMLElement) {
		re := regexp.MustCompile(`\d+`)
		//movie := e.Request.Ctx.GetAny("movie").(*Movie)
		movie.Year = re.FindString(e.Text)
		fmt.Println("年份:", movie.Year)
	})

	// 提取评分
	c.OnHTML(`strong[property="v:average"]`, func(e *colly.HTMLElement) {
		//movie := e.Request.Ctx.GetAny("movie").(*Movie)
		movie.Rating = e.Text
		fmt.Println("评分:", movie.Rating)
	})

	// 提取导演和演员信息
	c.OnHTML(`div#info`, func(e *colly.HTMLElement) {
		movie.Directors = []string{}
		movie.Actors = []string{}
		movie.Genre = []string{}
		//movie := e.Request.Ctx.GetAny("movie").(*Movie)
		// 导演（使用XPath）
		e.ForEach("a[rel='v:directedBy']", func(_ int, el *colly.HTMLElement) {
			if name := el.Text; name != "" {
				movie.Directors = append(movie.Directors, name)
			}
		})

		fmt.Println("导演:", movie.Directors)

		// 演员（使用CSS选择器）
		e.ForEach("a[rel='v:starring']", func(_ int, el *colly.HTMLElement) {
			if name := el.Text; name != "" {
				movie.Actors = append(movie.Actors, name)
			}
		})

		fmt.Println("主演:", movie.Actors)

		// 类型（使用属性选择器）
		e.ForEach(`span[property="v:genre"]`, func(_ int, el *colly.HTMLElement) {
			movie.Genre = append(movie.Genre, el.Text)
		})

		fmt.Println("类型:", movie.Genre)

		//保存到数据库
		if err := saveMovie(movie); err != nil {
			log.Printf("保存失败: %v", err)
		}

	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("href"))
		if url == "" {
			return
		}

		bloommutex.Lock()
		defer bloommutex.Unlock()

		// Bloom Filter检查
		if !bloomFilter.TestString(url) {
			// 添加到Bloom Filter并访问
			if strings.Contains(url, "/subject/") {
				bloomFilter.AddString(url)
				e.Request.Visit(url)
			}
		}
	})

}
