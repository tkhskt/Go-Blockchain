[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9ec239a01af5482a9590b8e6e75cb735)](https://www.codacy.com/app/gericass/Go-Blockchain?utm_source=github.com&utm_medium=referral&utm_content=gericass/Go-Blockchain&utm_campaign=badger)
# Go-Blockchain [![Go Report Card](https://goreportcard.com/badge/github.com/gericass/go-blockchain)](https://goreportcard.com/report/github.com/gericass/go-blockchain) [![Build Status](https://travis-ci.org/gericass/Go-Blockchain.svg?branch=master)](https://travis-ci.org/gericass/Go-Blockchain)
## Go言語によるブロックチェーンシミュレーター
ブロックチェーンの学習を目的としたブロックチェーンのシミュレーターです

## 仕様
- Socket通信を用いたアプリケーション単位での通信でネットワークを形成します
- トランザクションはノード数4以上で発生します
- マイニングは１分おきに実行されます。全てのノードは同時にマイニングを開始します
- トランザクションはランダムに発生します
- ブロックには以下の情報が格納されています
   1. ブロック（リスト）
   2. トランザクション（リスト）
   3. タイムスタンプ
   4. 直前のブロックのハッシュ
   5. ナンス
   6. 署名（リスト）
   7. マイニングをしたノード名
- チェーン分岐時の処理は未実装
- 署名時のハッシュの再計算は行わない
- 第三者によるトランザクションの承認は行わない
- ブロック情報、ノード情報、トランザクション情報は全ノードで共有しています
- データの送受信にはJSONを使用しています

## 使用方法
### エントリーポイントはblockchain.go

`go run blockchain.go -n node1 -p 8888 -t 1`

#### コマンドオプション(Required)

|オプション||
|:--:|:--:|
|-n|ノードの名前|
|-p|ポート番号|
|-t|ノードタイプ|

`Request Node`

- ブロック、トランザクション、ノードの情報を送信をリクエストするノードのポート番号を入力

##### ノードタイプ

|入力値|意味|
|:--:|:--:|
|1|ネットワークの最初のノード|
|2|その他のノード|


## 技術仕様
1. 使用言語
   - Go
2. 通信
   - トランスポート層(TCP)での通信(Socket通信)
3. ハッシュ計算アルゴリズム
   - SHA-256
   
## 解決すべき課題
 - ノード数が増えた際に全ノードが全ノードの情報を保持しているのは非常に非効率
    - ChordやKademliaなどのDHTアルゴリズムを用いて分散的にノードを管理する
    - Gossip Protocol(SWIM Protocol)を使うのが一般的
 - 並行処理について
    - 並行処理が多い(パケットの送受信をしながらマイニングをする　等)
    - 並行処理について、Go言語では文法レベルでの強力なサポートがあるが、他の言語でも果たして同じようにプログラム出来るのか
    - Shared memory型とMessage passing型の並行処理が混在している
    
## 参考リンク
[ブロックチェーンの基本的な仕組み](https://blockchain-jp.com/guides/4) 

[平成２７年度  我が国経済社会の情報化・サービス化に係る基盤整備
（ブロックチェーン技術を利用したサービスに関する国内外動向調査） 
報告書（pdf）](http://www.meti.go.jp/press/2016/04/20160428003/20160428003-2.pdf) 


    
