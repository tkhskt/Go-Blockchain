package transaction

import (
	"math/rand"
	"time"

	"../data"
)

func Send() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10000)
	if num == 1 && len(data.AllNode.List) > 2 && data.MyNode.Value > 0 { //ノード数が4以上の時取引をする
		random := rand.Intn(len(data.AllNode.List))
		to := data.AllNode.List[random]
		layout := "Mon Jan 2 15:04:05 MST 2006"
		times := time.Now().Format(layout)
		tr := data.Trans{Datatype: "Trans", ToCoin: to.Name, FromCoin: data.MyNode.Name, Sum: 5, Time: times}
		data.AllTrans.List = append(data.AllTrans.List, tr)
		data.MyNode.Value = data.MyNode.Value - 5
		data.AllNode.List[random].Value = data.AllNode.List[random].Value + 5
		for _, v := range data.AllNode.List {
			data.Dtype <- 1
			data.TransSolo <- tr
			data.PortNum <- v.Port
		}
	}

}
