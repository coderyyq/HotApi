package main

import (
	"HotApi/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hot/", api.GetHotListHandler)
	fmt.Println("      _ooOoo_")
	fmt.Println("     o8888888o")
	fmt.Println(`                           88" . "88`)
	fmt.Println("                           (| -_- |)")
	fmt.Println("                           O\\  =  /O")
	fmt.Println("                        ____/`---'\\____")
	fmt.Println("                      .'  \\||     ||/  `.")
	fmt.Println("                     /  \\||||  :  ||||/  \\ ")
	fmt.Println("                    /  _||||| -:- |||||-  \\ ")
	fmt.Println("                    |   | |||  -  ||| |   |")
	fmt.Println("                    | \\_|  ''\\---/''  |   |")
	fmt.Println("                    \\  .-\\__  `-`  ___/-. /")
	fmt.Println("                  ___`. .'  /--.--\\  `. . __")
	fmt.Println("               .'''< ```.___\\_<|>_/___. ```>'''.")
	fmt.Println("              | | :  `- \\`.;`\\ _ /`;.`/ - ` : | |")
	fmt.Println("              \\  \\ `-.   \\_ __\\ /__ _/   .-` /  /")
	fmt.Println("         ======`-.____`-.___\\_____/___.-`____.-'======")
	fmt.Println("         ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	fmt.Println("                         接口服务启动成功")
	fmt.Println("                      http://127.0.0.1:9000")
	log.Fatal(http.ListenAndServe(":9000", mux))
}
