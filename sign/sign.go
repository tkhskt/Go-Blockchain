package sign

import (
	"fmt"
	"math/rand"
	"time"

	"../data"
	"../slparse"
)

func listcheck(dt data.Node, bc data.Block) bool {
	if len(bc.Sign) > 0 {
		for _, v := range bc.Sign {
			if dt.Name == v {
				return false
			}
		}
	}
	if dt.Name == data.MyNode.Name { //ランダムに選ばれたノードが自分だった場合
		return false
	}
	if bc.Miner == dt.Name { //ランダムに選ばれたノードがマイナーだった場合
		return false
	}
	return true
}

func Sign(bc data.Block) { //署名
	key := true
	bc.Sign = append(bc.Sign, data.MyNode.Name) //署名する
	if len(bc.Sign) < 3 {
		for key { //署名に含まれていないノードからランダムに選んで送信
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(len(data.AllNode.List))
			if listcheck(data.AllNode.List[n], bc) {
				key = false
				data.Dtype <- 2
				data.BlockSolo <- bc
				data.PortNum <- data.AllNode.List[n].Port
			}
			time.Sleep(time.Microsecond * 5)
		}
	}

	if len(bc.Sign) == 3 { //三番目に署名した人が全員に送信
		for _, v := range data.AllNode.List {
			if v.Name != data.MyNode.Name {
				data.Dtype <- 2
				data.BlockSolo <- bc
				data.PortNum <- v.Port
			}
		}
		data.AllBlock.List = append(data.AllBlock.List, bc)
		slparse.Parse(2)
		fmt.Println(bc)
	}

}
