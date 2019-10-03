package main

import (
	"database/sql"
	"fmt"
	"github.com/jlaffaye/ftp"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type FtpConfig struct {
	Scan              int32
	Ftppath           string
	Temppath          string
	Localpath         string
	Ftpserverip       string
	Ftpserverport     string
	Ftpservername     string
	Ftpserverpassword string
	Suffflag          string
	Suff              string
}

func readFtpConfig() FtpConfig {
	data, _ := ioutil.ReadFile("config.yml")
	fmt.Println(string(data))
	t := FtpConfig{}
	//把yaml形式的字符串解析成struct类型

	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatal("配置文件错误")
	}
	return t
}
func copyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}
	//打开文件资源
	source_open, err := os.Open(source)
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()
	//只写模式打开文件 如果文件不存在进行创建 并赋予 644的权限。详情查看linux 权限解释
	dest_open, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	//进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
	}
}

func main() {

	conf := readFtpConfig()

	suffs := strings.Split(conf.Suff, ",")
	fmt.Println("开始ftp监控")
	wait := make(chan bool)
	go func() {
		for {
			fmt.Println("尝试重新下载ftp数据")
			//服务器ftp路径
			ftpPath := conf.Ftpserverip + ":" + conf.Ftpserverport
			//log.Println("-------------")
			c, err := ftp.Dial(ftpPath, ftp.DialWithTimeout(5*time.Second))
			//log.Println("-------------", err)
			if err != nil {
				log.Println("连接ftp服务器失败", err)
			} else {
				var user = conf.Ftpservername
				var password = conf.Ftpserverpassword
				err = c.Login(user, password)
				if err != nil {
					log.Println("用户名密码错误,", err)
				}
				ftppath := conf.Ftppath
				entries, err := c.List(ftppath)
				if err != nil {
					log.Println("获取FTP计划列表数据失败，error ---> ", err)
				} else {
					log.Println("获取FTP计划列表数据成功")
				}
				//连接sqlite数据库，将计划xml写入到sqlite中
				db, err := sql.Open("sqlite3", "./manager.db")
				if err != nil {
					log.Println("连接数据库失败, error --->", db)
				} else {
					for _, entry := range entries {
						if entry.Type == ftp.EntryTypeFile {
							//如果不是以.txt 就跳过
							if !strings.HasSuffix(entry.Name, ".txt") {
								continue
							}

							if conf.Suffflag == "true" {
								var containFlag = false
								for _, suff := range suffs {
									containFlag = strings.Contains(entry.Name, suff)
									if containFlag {
										break
									}
								}
								if !containFlag {
									fmt.Println("跳过文件 ", entry.Name)
									continue
								}
							}

							//判断是不是已经在数据库中了
							rows, err := db.Query("SELECT * FROM TaskInfo")
							if err != nil {
								log.Println("数据库语句错误, error --> ", err)
							} else {
								hasDownload := false
								for rows.Next() {
									var taskName string
									var taskCurrentTime string
									_ = rows.Scan(&taskName, &taskCurrentTime)
									if taskName == entry.Name {
										hasDownload = true
									}
								}
								if hasDownload {
									//log.Println("计划文件已经下载成功")
								} else {
									r, err := c.Retr(ftppath + entry.Name)

									if err != nil && err != io.EOF {
										log.Println("文件名：", entry.Name, "    error ----> ", err, "    ", r)
									} else {
										if r == nil {
											continue
										}
										//r.Read() && err != io.EOF
										buf, err := ioutil.ReadAll(r)
										r.Close()
										//如果读取的长度为0 就停止跳过读取文件
										if len(buf) == 0 {
											continue
										}

										//创建本地文件
										var temppath = conf.Temppath + entry.Name
										var localTempPath = temppath + ".temp"
										var localPath = conf.Localpath + entry.Name
										f, err := os.Create(localTempPath)
										flag := false
										//defer f.Close()
										if err != nil {
											log.Println("创建临时文件失败，文件名：", entry.Name, "   error is --->", err)
										} else {
											_, err = f.Write([]byte(buf))
											if err != nil {
												f.Close()
												log.Print("下载文件写入文件数据错误，文件名", entry.Name, "    error:", err)
											} else {
												f.Close()
												flag = true
											}
										}
										if flag {
											fmt.Println("下载文件成功,文件名：", entry.Name)

											time.Sleep(200 * time.Microsecond)
											//插入文件名称到数据库中
											stmt, err := db.Prepare("INSERT INTO TaskInfo(taskname,currenttime) values(?,?)")
											if err != nil {
												log.Println("数据库语句错误, error --> ", err)
											} else {
												t := time.Now()
												currenttime := t.Format("20060102150405")

												_, err := stmt.Exec(entry.Name, currenttime)
												if err != nil {
													log.Println("插入数据 error --> ", err)
												} else {
													//插入数据库成功后
													err = os.Rename(localTempPath, temppath)
													if err != nil {
														log.Println("重命名temp文件失败,error:", err)
													}
													err = os.Rename(temppath, localPath)
													if err != nil {
														log.Println("拷贝temp文件失败,error:", err)
													}
												}
											}

											//将下载的ftp文件放到本地数据库中，下一次读取就不在读取新的ftp计划

											//err = c.Delete(ftppath + entry.Name)
											//if err != nil {
											//	log.Fatal("删除ftp文件失败，文件名：", entry.Name, "    error", err)
											//}
										}
									}
								}
							}
						}
					}
					//断开数据库连接
					db.Close()
				}

				//断开连接
				if err := c.Quit(); err != nil {
					log.Print(err)
				}
			}

			//sleep 多少秒后重新下载文件
			scanMicrosecond := conf.Scan
			//fmt.Println(scanMicrosecond)

			time.Sleep(1000 * time.Duration(scanMicrosecond) * time.Microsecond)
		}
	}()
	<-wait
}
