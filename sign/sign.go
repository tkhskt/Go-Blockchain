package sign

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gericass/Go-Blockchain/data"
	"github.com/gericass/Go-Blockchain/slparse"
)

func listcheck(dt data.Node, bc data.Block) bool {
	if len(bc.Sign) > 0 {
		for _, v := range bc.Sign {
			if dt.Name == v {
				return false
			}
		}
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
			data.Dtype <- 2
			data.BlockSolo <- bc
			data.PortNum <- v.Port
		}
		for i := 0; i < len(data.AllNode.List); i++ {
			if bc.Miner == data.AllNode.List[i].Name {
				data.AllNode.List[i].Value = data.AllNode.List[i].Value + 10
			}
		}
		data.AllBlock.List = append(data.AllBlock.List, bc)
		slparse.Parse(2)
		fmt.Println(bc)
	}

}
