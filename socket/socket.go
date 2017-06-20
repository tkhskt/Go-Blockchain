package socket

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"../data"
	"../join"
	"../sign"
	"../slparse"
	"../transaction"
)

type un struct {
	Height string
	View   string
}

type jsons struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Dt   un
}

// Server はサーバー側
func Server() {
	service := ":" + data.MyNode.Port
	tcpAddrs, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listner, err := net.ListenTCP("tcp", tcpAddrs)
	checkError(err)
	for {
		conn, err := listner.Accept() //これより下にClient()があると動かない
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	//fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	messageLen, err := conn.Read(messageBuf)
	checkError(err)

	message := string(messageBuf[:messageLen])
	/*-----------------------データのパース--------------------*/
	cpTr := strings.Index(message, "Trans")
	cpBl := strings.Index(message, "Block")
	cpNo := strings.Index(message, "Node")
	cpTrLi := strings.Index(message, "TransList")
	cpBlLi := strings.Index(message, "BlockChain")
	cpNoLi := strings.Index(message, "NodeList")

	if cpTr >= 0 && cpBl < 0 && cpNo < 0 && cpTrLi < 0 { //Transデータの受け取り
		jsonBytes := ([]byte)(message)
		var js data.Trans
		if err := json.Unmarshal(jsonBytes, &js); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
		}
		transaction.Transaction(js)
	}
	if cpBl >= 0 && cpNo < 0 && cpBlLi < 0 { //Blockデータの受け取り
		//fmt.Println(message)
		jsonBytes := ([]byte)(message)
		var js data.Block
		if err := json.Unmarshal(jsonBytes, &js); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
		}
		if len(js.Sign) < 3 { //署名数が3未満だった時処理
			sign.Sign(js)
		}
		if len(js.Sign) == 3 { //署名数が3あったときの処理
			if data.AllBlock.List[len(data.AllBlock.List)-1].Number+1 == js.Number {
				data.AllBlock.List = append(data.AllBlock.List, js)
				slparse.Parse(2)
				if js.Miner == data.MyNode.Name { //マイナーが自分だった時自分のValueに報酬分を追加
					data.MyNode.Value = data.MyNode.Value + 10
				} else {
					for i := 0; i < len(data.AllNode.List); i++ { //マイナーの報酬分を追加
						if js.Miner == data.AllNode.List[i].Name {
							data.AllNode.List[i].Value = data.AllNode.List[i].Value + 10
						}
					}
				}
			}

			fmt.Println(js)
		}
	}
	if cpNo >= 0 && cpNoLi < 0 { //Nodeデータの受け取り && cpTr < 0 && cpNo < 0 && cpNoLi < 0
		jsonBytes := ([]byte)(message)
		var js data.Node
		if err := json.Unmarshal(jsonBytes, &js); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
		}
		join.Join(js)
	}
	if cpTrLi >= 0 { //TransListの受け取り
		jsonBytes := ([]byte)(message)
		var js data.TransList
		if err := json.Unmarshal(jsonBytes, &js); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
		}
		data.AllTrans = js
		slparse.Parse(1)
		fmt.Println(js)
	}
	if cpBlLi >= 0 { //BlockChainの受け取り
		jsonBytes := ([]byte)(message)
		var js data.BlockChain
		if err := json.Unmarshal(jsonBytes, &js); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
		}
		data.AllBlock = js
		slparse.Parse(2)
		fmt.Println(js)
	}
	if cpNoLi >= 0 { //NodeListの受け取り
		jsonBytes := ([]byte)(message)
		var js data.NodeList
		if err := json.Unmarshal(jsonBytes, &js); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
		}
		data.AllNode = js
		fmt.Println(js)
	}

	/*---------------------------------------------------------*/

}

// Client はクライアント側
func Client() {
	var jb []byte
	for {
		num := <-data.Dtype
		if num == 1 { //Trans(単品)データの送信
			d := <-data.TransSolo
			jb, _ = json.Marshal(d)
		} else if num == 2 { //Block(単品)データの送信
			d := <-data.BlockSolo
			jb, _ = json.Marshal(d)
		} else if num == 3 { //Node(単品)データの送信
			d := <-data.NodeSolo
			jb, _ = json.Marshal(d)
		} else if num == 4 { //Trans(リスト)データの送信
			d := data.AllTrans
			jb, _ = json.Marshal(d)
		} else if num == 5 { //Block(リスト)データの送信
			d := data.AllBlock
			jb, _ = json.Marshal(d)
		} else if num == 6 { //Node(リスト)データの送信
			p := append(data.AllNode.List, data.MyNode)
			d := data.NodeList{Datatype: "NodeList", List: p}
			jb, _ = json.Marshal(d)
		}
		pn := <-data.PortNum
		serverIP := "localhost" //サーバ側のIP
		serverPort := pn        //サーバ側のポート番号
		tcpAddr, _ := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
		myAddr := new(net.TCPAddr)
		conn, _ := net.DialTCP("tcp", myAddr, tcpAddr)

		/*-----json処理----*/

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		conn.Write([]byte(jb))
		/*-------------*/

		conn.Close() //TCPなので送信が終わったらCloseして再びコネクションを生成

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
