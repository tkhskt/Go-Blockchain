package mining

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Go-Blockchain/data"
)

var rs1Letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString1(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func Mining(d []data.Trans) {
	fmt.Println("----------------------------\nStart mining")
	fmt.Println(data.AllNode.List)
	now := time.Now()
	layout := "Mon Jan 2 15:04:05 MST 2006"
	transes := []data.Trans{}
	bc := data.AllBlock.List[len(data.AllBlock.List)-1] //この段階での最新ブロック
	transData := ""                                     //文字列化したトランザクションデータ

	for _, v := range d {
		tm, _ := time.Parse(layout, v.Time)
		duration := now.Sub(tm)
		if 1 < duration.Minutes() && duration.Minutes() <= 2 { //2分前から1分前まで      duration.Hours() == 0 && 1 < duration.Minutes() &&
			transData = transData + v.ToCoin + v.FromCoin + v.Time
			transes = append(transes, v)

		}
	}
	empsign := []string{}
	newbc := data.Block{Datatype: "Block", Number: bc.Number + 1, Transaction: transes, Time: now.Format(layout), PrevHash: bc.Hash, Nonce: "", Hash: "", Sign: empsign, Miner: data.MyNode.Name}

	str := newbc.Datatype + transData + newbc.Time + newbc.PrevHash + newbc.Miner
	var hash [32]byte
	for {
		rand.Seed(time.Now().UnixNano())
		nonce := RandString1(20)
		hash = sha256.Sum256([]byte(str + nonce))
		var bar []byte = hash[:]
		bin := binary.BigEndian.Uint64(bar)
		key := true
		if bin < 1000000000000000 { //ハッシュ値1000000000000000以下なら成功
			newbc.Nonce = nonce
			newbc.Hash = strconv.FormatUint(bin, 10)
			if data.AllBlock.List[len(data.AllBlock.List)-1].Number+1 == newbc.Number { //自分がマイニングしたナンバーのデータがほかのノードによって既にマイニング済みでなかった場合
				fmt.Println(bin)
				n := 0
				for key {
					rand.Seed(time.Now().UnixNano())
					n = rand.Intn(len(data.AllNode.List) - 1)
					if data.AllNode.List[n].Name != data.MyNode.Name {
						key = false
					}
					time.Sleep(time.Microsecond * 3)
				}
				data.Dtype <- 2
				data.BlockSolo <- newbc
				data.PortNum <- data.AllNode.List[n].Port
			}
			break
		}
		end := time.Now()
		duration := end.Sub(now).Seconds()
		if duration >= 50 {
			break
		}
		time.Sleep(time.Millisecond)
	}
}
