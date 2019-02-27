// mongodb project main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	Id    int
	CName string
}

//这种没有标记说明json或者bson的映射
type User struct {
	Uid      int
	Username string
	Password string
	Age      int
	Email    string
	Hobby    []string
	Course   []int
}
type UserBson struct {
	Id       string   `bson:"_id" json:"_id"`
	Uid      int      `bson:"uid"  json:"uid"`
	Username string   `bson:"username"  json:"username"`
	Password string   `bson:"password"  json:"password"`
	Age      int      `bson:"age"  json:"age"`
	Email    string   `bson:"email"  json:"email"`
	Hobby    []string `bson:"hobby"  json:"hobby"`
	Course   []int    `bson:"cource"  json:"cource"`
}

func main() {
	//mongodb://127.0.0.1:27017/?gssapiServiceName=mongodb
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

	c := db.C("course")
	c.DropCollection()
	c.Insert(Course{Id: 1, CName: "chinese"})
	c.Insert(Course{Id: 2, CName: "English"})
	c.Insert(Course{Id: 3, CName: "math"})
	c = db.C("user")
	//清空数据库
	c.DropCollection()

	//插入数据

	{

		user := User{Uid: 1, Username: "xiaoming", Password: "xiaoming", Email: "xiaoming@outlook.com", Age: 10, Hobby: []string{"swim", "game"}, Course: []int{1, 2}}
		user1 := User{Uid: 2, Username: "xiaohong", Password: "xiaohong", Email: "xiaohong@outlook.com", Age: 12, Hobby: []string{"shopping", "read"}, Course: []int{1}}
		c.Insert(user)
		c.Insert(user1)

		//批量插入
		{
			var users []interface{}
			users = append(users, User{Uid: 3, Username: "xiaogang", Password: "xiaogang", Email: "xiaogang@outlook.com", Age: 15, Hobby: []string{"boll", "game"}, Course: []int{2, 3}})
			users = append(users, User{Uid: 4, Username: "xiaoxin", Password: "xiaoxin", Email: "xiaoxin@outlook.com", Age: 13, Hobby: []string{"read", "game"}, Course: []int{1, 3}})
			err := c.Insert(users...)
			fmt.Println(err)

		}
	}

	//获取集合信息
	{

		//查询所有的用户
		{
			result := make([]User, 0)
			c.Find(bson.M{}).All(&result)

			//			data, err := json.MarshalIndent(result, "", "\t")
			//			if err != nil {
			//				log.Fatal(err)
			//				panic(err)
			//			}
			fmt.Println(result)

			//简单分页 ,页数1，每一页数据3
			var pagenum = 1
			var pagesize = 3
			c.Find(bson.M{}).Limit(pagesize).Skip(pagenum * pagesize).All(&result)
			fmt.Println(result)
		}

		//使用bson获取数据
		{
			c := db.C("user")
			result := make([]UserBson, 0)
			c.Find(bson.M{}).All(&result)
			data, err := json.MarshalIndent(result, "", "\t")
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			fmt.Print(string(data))
		}
		//条件查询
		{
			var oneResult User
			//使用find方法的时候数据类型要一一对应，不然无法查询
			c.Find(bson.M{"uid": 1}).One(&oneResult)
			fmt.Println(oneResult)
		}
		{
			var oneResult User
			//无法查询到数据
			//c.FindId("5c3b52f76af98c72e2b30666").One(&oneResult)

			//可以查询到数据
			objid := bson.ObjectIdHex("5c3b52f76af98c72e2b3066a")
			c.FindId(objid).One(&oneResult)
			//等价
			c.Find(bson.M{"_id": bson.ObjectIdHex("5c3b52f76af98c72e2b3066a")}).One(&oneResult)
			fmt.Println(oneResult)
			//查询喜欢read的同学，字段为切片 可以直接包含
			result := make([]User, 0)
			c.Find(bson.M{"hobby": "read"}).All(&result)
			fmt.Println(result)
			//查询喜欢read的同学且选了数学课程的同学并且只解析用户名和课程
			//select{"filedname":0}，表示忽略该字段则结果不返回此字段
			//select{"filedname":1}，表示关注该字段则只返回关注字段
			result = make([]User, 0)
			c.Find(bson.M{"hobby": "read", "course": 3}).Select(bson.M{"username": 1, "course": 1}).All(&result)
			fmt.Println(result)
			//查找小于12岁的同学
			c.Find(bson.M{"age": bson.M{"$lt": 12}}).All(&result)
			fmt.Println(`age 小于12的同学`, result)
			//查找小于等于12岁的同学
			c.Find(bson.M{"age": bson.M{"$lte": 12}}).All(&result)
			fmt.Println(`age 小于等于12的同学`, result)
			//查找不等于12岁的同学
			c.Find(bson.M{"age": bson.M{"$ne": 12}}).All(&result)
			fmt.Println(`age 不等于12岁的同学`, result)
			//查找爱好包含shopping 或game的同学
			c.Find(bson.M{"hobby": bson.M{"$in": []string{"shopping", "game"}}}).All(&result)
			fmt.Println(`查找爱好包含shopping 和game的同学`, result)
			//查找爱好不包含shopping 或game的同学
			c.Find(bson.M{"hobby": bson.M{"$nin": []string{"shopping", "game"}}}).All(&result)
			fmt.Println(`查找爱好不包含shopping 或game的同学`, result)
			//查找爱好包含swim和game的同学
			c.Find(bson.M{"hobby": []string{"swim", "game"}}).All(&result)
			fmt.Println(`查找爱好包含shopping 和game的同学`, result)

			//$project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
			//$match：用于过滤数据，只输出符合条件的文档。match：用于过滤数据，只输出符合条件的文档。match使用MongoDB的标准查询操作。
			//$limit：用来限制MongoDB聚合管道返回的文档数。
			//$skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。
			//$unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
			//$group：将集合中的文档分组，可用于统计结果。
			//$sort：将输入文档排序后输出。
			//$geoNear：输出接近某一地理位置的有序文档。
			//匹配优先级，优先匹配切片前面的条件
			//简单分页 ,页数0，每一页数据3,页数从第一页开始
			var pagenum = 0
			var pagesize = 3
			pipe := c.Pipe([]bson.M{

				bson.M{"$match": bson.M{"hobby": "game"}}, //匹配喜欢读书的
				bson.M{"$limit": pagesize},
				bson.M{"$skip": pagenum * pagesize},
				//bson.M{"$project": bson.M{"username": 1}}, //只输出username
			})
			resp := []bson.M{}
			pipe.All(&resp)
			fmt.Println(resp)
			pipe.All(&result)
			fmt.Println(result)
		}
	}
	//更新用户信息
	//修改xiaoming的密码为password
	//没有set标记会删除其他数据
	//c.Update(bson.M{"username": "xiaoming"}, bson.M{"password": "password"})
	//正确更新方式
	c.Update(bson.M{"username": "xiaoming"}, bson.M{"$set": bson.M{"password": "password"}})
}
