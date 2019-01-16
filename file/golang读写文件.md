### golang读写文件操作

**func ReadFile**

```go
func ReadFile(filename string) ([]byte, error)
```

example

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func FatalError(hint string, err error) {
    if err != nil {
        log.Fatal(fmt.Sprintf("error : %s", hint), err)
    }
}
func main() {
    //读取文件,文件必须存在，不然会err出现文件找不到的错误

    bytes, err := ioutil.ReadFile("1.txt")
    FatalError("ReadFile ", err)
    fmt.Println(string(bytes))
}
```

**func WriteFile**

```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func FatalError(hint string, err error) {
    if err != nil {
        log.Fatal(fmt.Sprintf("error : %s", hint), err)
    }
}
func main() {
    //写数据到文件,通过iotuil写文件默认调用的是打开文件操作 os.O_WRONLY|os.O_CREATE|os.O_TRUNC,会清空缓存区
    err := ioutil.WriteFile("2.txt", []byte("hello go"), os.FileMode(os.ModeAppend))
    FatalError("ReadFile ", err)

}
```

使用os读文件

**func Open**

```go
func Open(name string) (*File, error)
//文件路径，打开模式，文件权限
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
)

func FatalError(hint string, err error) {
    if err != nil {
        log.Fatal(fmt.Sprintf("error : %s", hint), err)
    }
}
func main() {
    f, err := os.Open("1.txt")
    //f, err := os.OpenFile("1.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
    FatalError("os Open", err)
    defer f.Close()
    buffer, err := ioutil.ReadAll(f)
    FatalError("ReadAll", err)
    fmt.Println(string(buffer))

    //按行读取文件
    b := bufio.NewReader(f)
    FatalError("bufio NewReader", err)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil && err != io.EOF {
            FatalError("ReadAll", err)
        }
        fmt.Println(string(line))
        if err == io.EOF {
            break
        }
    }
}
```

使用os写文件，go写入文件是需要权限，因此可以使用OpenFile

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func FatalError(hint string, err error) {
    if err != nil {
        log.Fatal(fmt.Sprintf("error : %s", hint), err)
    }
}
func main() {
    //第二个参数控制权限，文件不存在就以perm的权限创建file，如果存在就会在写入之前清空内容
    //这里的权限是追加，创建，读写,缺少os.O_APPEND就会在文件头添加
    // os.O_CREATE 表示文件不存在就会创建
    // os.O_APPEND 表示以追加内容的形式添加
    // os.O_WRONLY 表示只写模式
    // os.O_RDONLY 表示只读模式
    // os.O_RDWR 表示读写模式
    // os.O_EXCL used with O_CREATE, file must not exist
    // os.O_SYNC I/O同步的方式打开
    // os.O_TRUNC if possible, truncate file when opened.
    f, err := os.OpenFile("1.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
    FatalError("os Open", err)
    defer f.Close()
    var bytes = []byte{76, 78, 85, 85, 64}
    siz, err := f.Write(bytes)
    FatalError("file write ", err)
    fmt.Println(siz)
    siz, err = f.WriteString("fakjfhlashfa")
    FatalError("file WriteString ", err)
    fmt.Println(siz)
    //相对文件头偏移120个字符添加，os.O_APPEND存在的时候优先文件尾部添加
    _, err = f.WriteAt([]byte("!!!!!!!!!!"), int64(120))
    FatalError("file WriteAt ", err)

    //使用bufio
    b := bufio.NewWriter(f)
    for i := 0; i < 10; i++ {
        b.Write([]byte(fmt.Sprintf("%d", i)))
        b.WriteByte('\n')
        b.WriteString("test")
        //将bufio缓存的文件刷新到文件中，不手动刷新就等程序结束调用
        b.Flush()
    }
}
```

读取文件速度对比

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "time"
)

func read0(path string) string {
    f, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Printf("%s\n", err)
        panic(err)
    }
    return string(f)
}

func read1(path string) string {
    fi, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fi.Close()

    chunks := make([]byte, 1024, 1024)
    buf := make([]byte, 1024)
    for {
        n, err := fi.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if 0 == n {
            break
        }
        chunks = append(chunks, buf[:n]...)
    }
    return string(chunks)
}

func read2(path string) string {
    fi, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    r := bufio.NewReader(fi)

    chunks := make([]byte, 1024, 1024)

    buf := make([]byte, 1024)
    for {
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if 0 == n {
            break
        }
        chunks = append(chunks, buf[:n]...)
    }
    return string(chunks)
}

func read3(path string) string {
    fi, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    fd, err := ioutil.ReadAll(fi)
    return string(fd)
}

func main() {

    file := "test.log"

    start := time.Now()

    read0(file)
    t0 := time.Now()
    fmt.Printf("Cost time %v\n", t0.Sub(start))

    read1(file)
    t1 := time.Now()
    fmt.Printf("Cost time %v\n", t1.Sub(t0))

    read2(file)
    t2 := time.Now()
    fmt.Printf("Cost time %v\n", t2.Sub(t1))

    read3(file)
    t3 := time.Now()
    fmt.Printf("Cost time %v\n", t3.Sub(t2))

}
```

```bash
Cost time 2.0343ms
Cost time 5.9848ms
Cost time 4.0146ms
Cost time 3.0041ms
```



创建文件

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Create("file.log")
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(file)
}
```

**判断文件或者文件夹是否存在**

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Stat("44.txt")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println("文件存在")

}
```

创建目录

```go

package main

import (
    "log"
    "os"
)

func main() {
    // 创建当个目录
    err := os.Mkdir("tmp", 0755)
    if err != nil {
        log.Fatalln(err)
    }

    // 递归创建目录
    err = os.MkdirAll("tmp/tmp1/tmp2", 0755)
    if err != nil {
        log.Fatalln(err)
    }
}
```

获取当前的软件路径

```go
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalln(err)
	}
	//绝对路径

	path, err := filepath.Abs(file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(path)

	//相对路径
	rst := filepath.Dir(path)

	fmt.Println(rst)
```

判断一个路径是文件还是文件夹

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Stat("1111")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}
	if os.IsNotExist(err) {
		log.Fatalln("文件或者文件夹不存在")
		return
	}

	if f.IsDir() {
		fmt.Println("这是一个文件夹")
	} else {
		fmt.Println("这是一个文件")
	}
}
```
