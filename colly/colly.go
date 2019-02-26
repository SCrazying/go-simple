// colly_simple project main.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func MD5(text string) string {

	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

type Page struct {
	TotalPage int `json:"totalPage"`
	CurPage   int `json:"curPage"`
}

type House struct {
	Id         string `bson:"_id" json:"_id"`
	Url        string `bson:"url" json:"url"`
	Title      string `bson:"title" json:"title"`
	Address    string `bson:"address" json:"address"`
	HouseInfo  string `bson:"houseInfo" json:"houseInfo"`
	Image      string `bson:"image" json:"image"`
	Flood      string `bson:"flood" json:"flood"`
	Position   string `bson:"position" json:"position"`
	Star       string `bson:"star" json:"star"`
	Watch      string `bson:"watch" json:"watch"`
	PushTime   string `bson:"pushTime" json:"pushTime"`
	Tag        string `bson:"tag" json:"tag"`
	UnitPrice  string `bson:"unitPrice" json:"unitPrice"`
	TotalPrice string `bson:"totalPrice" json:"totalPrice"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		log.Fatal("err is ", err)
		panic(err)
	}
	defer session.Close()
	fmt.Println("connect mongo ok, port is 27017")
	session.SetMode(mgo.Monotonic, true)

	//使用数据库
	db := session.DB("mongodbtest")
	//使用house 集合
	c := db.C("house")

	url_map := make(map[string]bool, 0)

	control := colly.NewCollector()
	control.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"

	// control.AllowedDomains = []string{
	// 	"https://sh.lianjia.com",
	// 	"sh.lianjia.com",
	// 	"https://*.lianjia.com/ershoufang",
	// 	"mianyang.lianjia.com/ershoufang",
	// }
	control.OnError(func(r *colly.Response, err error) {
		fmt.Println("error ", err)
	})
	control.OnRequest(func(r *colly.Request) {
		//r.Headers.Set()
		r.Headers.Set("Host", "dig.lianjia.com")

		fmt.Println("visiting", r.URL.String())
	})
	control.OnResponse(func(r *colly.Response) {
		//fmt.Println("visited", string(r.Body))
	})

	control.OnHTML(".sellListContent", func(e *colly.HTMLElement) {

		pclass := e.Attr("class")
		if pclass == "sellListContent" {
			e.ForEach("li", func(zzz int, e *colly.HTMLElement) {
				url := e.ChildAttr("div.title>a", "href")
				title := e.ChildText("div.title>div>a")
				address := e.ChildText("div.address>div>a")
				houseInfo := e.ChildText("div.address>div.houseInfo")
				flood := e.ChildText("div.positionInfo")
				position := e.ChildText("div.positionInfo>a")
				followInfo := e.ChildText("div.followInfo")
				followInfos := strings.Split(followInfo, "/")

				var star, watch, pushTime string
				if len(followInfos) == 3 {
					star = strings.Replace(followInfos[0], " ", "", -1)
					watch = strings.Replace(followInfos[1], " ", "", -1)
					pushTime = strings.Replace(followInfos[2], " ", "", -1)
				}
				totalPrice := e.ChildText("div.totalPrice")
				unitPrice := e.ChildText("div.unitPrice")
				tag := e.ChildText("div.tag")

				image := e.ChildAttr("a>img.lj-lazy", "data-original")

				var t_house House
				t_house.Id = MD5(url)
				t_house.Url = url
				t_house.Title = title
				t_house.Address = address
				t_house.HouseInfo = houseInfo
				t_house.Image = image
				t_house.Flood = flood
				t_house.Position = position
				t_house.Star = star
				t_house.Watch = watch
				t_house.PushTime = pushTime
				t_house.Tag = tag
				t_house.UnitPrice = unitPrice
				t_house.TotalPrice = totalPrice
				err := c.Insert(t_house)
				if err != nil {
					//插入失败，更新数据
					//log.Fatal("insert false", err)

					selector := bson.M{"_id": bson.ObjectIdHex(t_house.Id)}
					fmt.Println(selector)
					//c.UpdateId(selector, t_house)
				}
			})

		}

	})
	control.OnHTML(".page-box.house-lst-page-box", func(e *colly.HTMLElement) {

		url := e.Request.URL

		url_slice := strings.Split(url.String(), "ershoufang")
		pageinfo_str := e.Attr("page-data")
		var pageinfo Page
		json.Unmarshal([]byte(pageinfo_str), &pageinfo)
		for i := 0; i < pageinfo.TotalPage; i++ {
			new_url := url_slice[0] + "ershoufang/pg" + strconv.FormatInt(int64(i), 10) + "/"
			url_map[new_url] = false
			control.Visit(new_url)
		}
	})

	control.Visit("https://mianyang.lianjia.com/ershoufang/")
	for true {

	}
}
