package slparse

import (
	"sort"
	"time"

	"../data"
)

func Parse(a int) {
	if a == 1 {
		sort.Slice(data.AllTrans.List, func(i, j int) bool { //トランザクション配列の中身をソート
			layout := "Mon Jan 2 15:04:05 MST 2006"
			ti, _ := time.Parse(layout, data.AllTrans.List[i].Time)
			tj, _ := time.Parse(layout, data.AllTrans.List[j].Time)
			ti.Before(tj)
			return ti.Before(tj)
		})
	}
	if a == 2 {
		sort.Slice(data.AllBlock.List, func(i, j int) bool { //ブロックチェーン配列の中身をソート
			layout := "Mon Jan 2 15:04:05 MST 2006"
			ti, _ := time.Parse(layout, data.AllBlock.List[i].Time)
			tj, _ := time.Parse(layout, data.AllBlock.List[j].Time)
			ti.Before(tj)
			return ti.Before(tj)
		})
	}
}
