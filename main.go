package main
import (
    "os"
    "cloudgo/service"
    flag "github.com/spf13/pflag"
)
const (
    //使用端口9000
    PORT string = "9000" 
)
func main() {
    port := os.Getenv("PORT") 
    if len(port) == 0 {
        port = PORT
    }
    //端口解析
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    //启动server
    service.NewServer(port)
}