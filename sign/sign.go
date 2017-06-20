package sign

import (
	"fmt"
	"math/rand"
	"time"

	"../data"
	"../slparse"
)

func Sign(bc data.Block) { //署名
	bc.Sign = append(bc.Sign, data.MyNode.Name) //署名する
	key := true
	rand.Seed(time.Now().UnixNano())
	for key { //署名に含まれていないノードからランダムに選んで送信
		n := rand.Intn(len(data.AllNode.List) - 1)
		for _, v := range bc.Sign {
			if v != data.AllNode.List[n].Name {
				key = false
				data.Dtype <- 2
				data.BlockSolo <- bc
				data.PortNum <- data.AllNode.List[n].Port
				break
			}
		}
	}
	if len(bc.Sign) == 3 { //三番目に署名した人が全員に送信
		for _, v := range data.AllNode.List {
			data.Dtype <- 2
			data.BlockSolo <- bc
			data.PortNum <- v.Port
		}
		data.AllBlock.List = append(data.AllBlock.List, bc)
		slparse.Parse(2)
		fmt.Println(bc)
	}

}
