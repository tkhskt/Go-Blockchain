package transaction

import (
	"github.com/gericass/Go-Blockchain/data"
	"github.com/gericass/Go-Blockchain/slparse"
)

func Transaction(d data.Trans) {
	if d.ToCoin == data.MyNode.Name { //自分に送金された場合
		data.MyNode.Value = data.MyNode.Value + d.Sum

	}
	for i := 0; i < len(data.AllNode.List); i++ {
		if data.AllNode.List[i].Name == d.FromCoin {
			data.AllNode.List[i].Value = data.AllNode.List[i].Value - d.Sum
		}
		if data.AllNode.List[i].Name == d.ToCoin {
			data.AllNode.List[i].Value = data.AllNode.List[i].Value + d.Sum
		}
	}
	data.AllTrans.List = append(data.AllTrans.List, d)
	slparse.Parse(1)
	//fmt.Println(data.AllTrans)
}
