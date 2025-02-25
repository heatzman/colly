package main

import (
	"github.com/bits-and-blooms/bloom/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"log"
)

func main() {
	initRedis()
	// 初始化MySQL
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatal("数据库初始化失败：", err)
	}
	capacity := uint(1000000)
	fpRate := 0.001
	bloomFilter = bloom.NewWithEstimates(capacity, fpRate)

	//创建collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (windows Nt 10.0; win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"),
		colly.AllowedDomains("movie.douban.com"),
		colly.MaxDepth(MaxDepth),
		colly.Async(true),
	)

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher(
		"socks5://114.229.217.214:24029",
		"socks5://36.20.124.17:20352")
	if err != nil {
		log.Fatal(err)
	}
	// 【设置代理IP】 ，这里使用的是轮询ip方式
	c.SetProxyFunc(rp)

	// 配置爬虫参数
	c.Limit(&colly.LimitRule{
		//DomainGlob:  `movie\.douban\.com`,
		Parallelism: parallelism,
		RandomDelay: RandomDelay,
	})

	registerCallbacks(c)

	// 启动爬虫
	c.Visit("https://movie.douban.com/")
	c.Wait()

}
