package main

import (
	log "kcp-demo/app/module"
	"time"
)

func main() {
	log.Start("127.0.0.1", "micro", "./golang.log")

	log.Req("aa", "200", 1000, nil)
	log.Req("aa", "200", 600, nil)
	log.Req("aa", "200", 1300, nil)
	log.Req("bb", "300", 3000, nil)


	log.Rpc("aa", "200", 1000, "www.baidu.com","199.99.99.99")
	log.Rpc("aa", "200", 600, "www.baidu.com","199.99.99.99")
	log.Rpc("aa", "200", 1300, "www.baidu.com","199.99.99.99")
	log.Rpc("bb", "300", 3000, "www.baidu.com","199.99.99.99")

	log.Src("mysql_pool_usage", "Mysql_Pool", 12, 30)

	time.Sleep(6*time.Second)
	log.Stop()
	time.Sleep(30*time.Second)

}
