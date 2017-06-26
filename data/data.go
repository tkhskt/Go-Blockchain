package data

type Trans struct { //トランザクションデータ
	Datatype string `json:"datatype"`
	ToCoin   string `json:"tocoin"`   //誰へコインを渡したか
	FromCoin string `json:"fromcoin"` //誰がコインを受け取ったか
	Sum      int    `json:"sum"`      //金額
	Time     string `json:"time"`     //タイムスタンプ
}

type Block struct { //ブロック
	Datatype    string   `json:"datatype"`
	Number      int      `json:"number"`      //ブロックナンバー
	Transaction []Trans  `json:"transaction"` //トランザクションデータを格納するスライス []Trans
	Time        string   `json:"time"`        //タイムスタンプ
	PrevHash    string   `json:"prevhash"`    //前のブロックのハッシュ
	Nonce       string   `json:"nonce"`       //ナンス
	Hash        string   `json:"hash"`        //ハッシュ
	Sign        []string `json:"sign"`        //署名
	Miner       string   `json:"miner"`       //マイナーのName
}

type Node struct {
	Datatype string `json:"datatype"`
	Name     string `json:"name"`  //ノード名
	Value    int    `json:"value"` //ノードの保有コイン
	Port     string `json:"port"`  //ポート番号
	Sign     string `json:"sign"`  //全ノードに送信したノードのName
}

type TransList struct {
	Datatype string  `json:"datatype"`
	List     []Trans `json:"list"`
}
type BlockChain struct {
	Datatype string  `json:"datatype"`
	List     []Block `json:"list"`
}
type NodeList struct {
	Datatype string `json:datatype`
	List     []Node `json:"list"`
}

var MyNode = Node{Datatype: "Node", Name: "", Value: 10, Port: "", Sign: ""} //自分のノード情報

var a []Trans
var b []Block
var c []Node

var AllTrans = TransList{Datatype: "TransList", List: a}   //すべてのトランザクションを保存する
var AllBlock = BlockChain{Datatype: "BlockChain", List: b} //すべてのブロックを保存する
var AllNode = NodeList{Datatype: "NodeList", List: c}      //自分以外のすべてのノードを保存する

var PortNum = make(chan string, 10) //Client()で送信する送信先のポート番号
var Dtype = make(chan int, 10)      //Client()で送信するデータのタイプ(6パターン)

var TransSolo = make(chan Trans, 10)    //Client()にTransデータを送るためのChannel
var BlockSolo = make(chan Block, 10000) //Client()にBlockデータを送るためのChannel
var NodeSolo = make(chan Node, 10)      //Client()にNodeデータを送るためのChannel
