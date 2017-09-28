package join

import (
	"fmt"

	"time"

	"github.com/gericass/Go-Blockchain/data"
)

func Join(nd data.Node) {
	if nd.Sign == "" { //まだ新ノードが全ノードに送信されていなかった場合
		/*----新規参加ノードのNodeデータを全ノードに送信-----*/
		fmt.Println(nd)
		nd.Sign = data.MyNode.Name
		for _, v := range data.AllNode.List {
			data.Dtype <- 3
			data.NodeSolo <- nd
			data.PortNum <- v.Port
		}
		/*----Nodeのリストを新規参加ノードに送信------*/
		data.Dtype <- 6
		data.PortNum <- nd.Port
		/*----Blockのリストを新規参加ノードに送信-----*/
		data.Dtype <- 5
		data.PortNum <- nd.Port
		/*----Transのリストを新規参加ノードに送信-----*/
		data.Dtype <- 4
		data.PortNum <- nd.Port
	}
	time.Sleep(time.Millisecond * 5) //リスト送信処理中にAllNode.Listに新ノードの情報が格納されないようにする
	data.AllNode.List = append(data.AllNode.List, nd)

}
