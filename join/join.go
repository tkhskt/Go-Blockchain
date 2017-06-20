package join

import (
	"fmt"

	"../data"
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

	data.AllNode.List = append(data.AllNode.List, nd)

}
