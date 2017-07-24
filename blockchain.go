package main

import (
	"fmt"
	"time"

	"./data"
	"./mining"
	"./socket"
	"./transaction"
)

var nd []data.Node

var gosign = make(chan string, 10) //並行処理スタートのゴーサイン
var sn = make(chan int, 10)

var nodetype int
var joinreq string

func main() {
	fmt.Println("Enter your name")
	fmt.Scan(&data.MyNode.Name)
	fmt.Println("--------------------------------")
	fmt.Println("Enter your port")
	fmt.Scan(&data.MyNode.Port)
	fmt.Println("--------------------------------")
	fmt.Println("Are you first?\n(1)Yes (2)No")
	fmt.Scan(&nodetype)

	if nodetype == 1 { //最初のノードだった場合
		layout := "Mon Jan 2 15:04:05 MST 2006"
		times := time.Now().Format(layout)
		//var dt = []data.Trans{data.Trans{Datatype: "Trans", ToCoin: "nil", FromCoin: "nil", Sum: 0, Time: times}}
		var ph = []string{"nil", "nil", "nil"}
		dummyblock := data.Block{Datatype: "Block", Number: 1, Transaction: []data.Trans{}, Time: times, PrevHash: "000000", Nonce: "000000", Hash: "000000", Sign: ph, Miner: "nil"}
		data.AllBlock.List = append(data.AllBlock.List, dummyblock)
	} else if nodetype == 2 {
		fmt.Println("Request Node")
		fmt.Scan(&joinreq)
	}
	gosign <- "Start"
	fmt.Println(<-gosign)
	go socket.Server()
	go socket.Client()
	time.Sleep(time.Millisecond * 10)
	if nodetype == 2 {
		data.Dtype <- 3
		data.NodeSolo <- data.MyNode
		data.PortNum <- joinreq
	}
	//Server()とClient()は並列に処理しないと正常に動作しない！！
	now := time.Now().Minute()
	for { //mainプロセスが終了しないための無限ループ
		fornow := time.Now().Minute() //1分おきにマイニングを実行
		if fornow != now && len(data.AllNode.List) > 3 {
			sn <- 3
			<-sn
			go mining.Mining(data.AllTrans.List)
			now = fornow
		}
		if len(data.AllNode.List) > 3 {
			transaction.Send()
		}
		time.Sleep(time.Nanosecond * 10)
	}

}
