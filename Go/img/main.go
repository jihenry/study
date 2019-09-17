package main

import (
"bytes"
"fmt"
"io/ioutil"
"net/http"
"os"
"regexp"
"strconv"
)

const (
PATH string = "~/meizi" //文件存储路径
STARTURL string = "http://www.mzitu.com/model" //妹子图模块列表页url
CONCURRENCY int = 3 //并发下载数
)

var (
c1 chan string //通道：负责STARTURL,以后可以做成从命令参数里读取哦
c2 chan string //通道: 负责传输所有的模块url
c3 chan []string //通道：负责传输imgUrl
c4 chan int //通道: 负责传输每张图片的下载完成状态
c5 chan int //通道：负责传输当前下载数
)

func init() {
c1 = make(chan string, 1)
c2 = make(chan string, 100)
c3 = make(chan []string, 1000)
c4 = make(chan int, 3)
c5 = make(chan int, 10)
go CgetList()
go Cdownload()
}
func main() {
c1 <- STARTURL
go CgetModel()
num := 0
for count := range c5 {
num = num + count
fmt.Println("已下载:", num)
}
}

//调度器, 拉取所有模块
func CgetModel() {
modelPage := getPage(<-c1)
for i := 1; i <= modelPage; i++ {
modelUrl := STARTURL + "/page/" + strconv.Itoa(i)
c2 <- modelUrl
}
}

//调度器拉取所有图片url,这里其实还可以多分一层
func CgetList() {
k := 0
tmp := make([]string, 3)
for modelUrl := range c2 {
imgLists := getList(modelUrl)
for _, imgList := range imgLists {
imgPage := getPage(imgList)
for j := 1; j <= imgPage; j++ {
imgUrl := imgList + "/" + strconv.Itoa(j)
if k < CONCURRENCY {
tmp[k] = imgUrl
k++
} else {
c3 <- tmp
k = 0
}
}
}
if k != 0 {
c3 <- tmp
k = 0
}
}

}

//调度器, 下载图片
func Cdownload() {
for imgUrls := range c3 {
if len(imgUrls) > 0 {
for _, imgUrl := range imgUrls {
go func() {
download(imgUrl)
c4 <- 1
}()
}
num := 0
for k := range c4 {
num = num + k
if num == len(imgUrls) {
c5 <- num
break
}
}
}
}
}

//图片列表
func getList(url string) (l []string) {
reg, _ := regexp.Compile(`<h2><a href="(http://www.mzitu.com/\d*)" title="(.*?)" target="_blank">.*?</a></h2>`)
_, html, _ := getHtml(url)
lists := reg.FindAllStringSubmatch(html, 1000)
for _, list := range lists {
l = append(l, list[1])
}
return
}

//下载html
func getHtml(url string) (error, string, error) {
response, err := http.Get(url)
defer response.Body.Close()
html, err1 := ioutil.ReadAll(response.Body)
return err, string(html), err1
}

//获取最大分页
func getPage(url string) (page int) {
_, html, _ := getHtml(url)
reg, _ := regexp.Compile(`<span>(\d*)</span>`)
s := reg.FindAllStringSubmatch(html, 200)
if len(s) < 2 {
fmt.Println("获取失败")
os.Exit(-1)
}
page, _ = strconv.Atoi(s[len(s)-1][1])
return

}

//下载图片
func download(url string) {
reg, _ := regexp.Compile(`<p><a href="http:\/\/www.mzitu.com/.*?" ><img src="(.*?)" alt="(.*?)" /></a></p>`)
reg1, _ := regexp.Compile(`http:\/\/pic\.dofay\.com/(.*)`)
_, html, _ := getHtml(url)
iterms := reg.FindAllStringSubmatch(html, 100)
for _, iterm := range iterms {
imgUrl := iterm[1]
imgPath := reg1.FindAllStringSubmatch(imgUrl, 100)
imgPaths := bytes.Split([]byte(imgPath[0][1]), []byte("/"))
path := PATH + "/" // + iterm[2]
imgResponse, _ := http.Get(imgUrl)
defer imgResponse.Body.Close()
imgByte, _ := ioutil.ReadAll(imgResponse.Body)
pInfo, pErr := os.Stat(path)
if pErr != nil || pInfo.IsDir() == false {
errDir := os.Mkdir(path, os.ModePerm)
if errDir != nil {
fmt.Println(errDir)
os.Exit(-1)
}
}
fn := path + "/" + string(imgPaths[len(imgPaths)-1])
_, fErr := os.Stat(fn)
var fh *os.File
if fErr != nil {
fh, _ = os.Create(fn)
} else {
fh, _ = os.Open(fn)
}
defer fh.Close()
fh.Write(imgByte)
}
}