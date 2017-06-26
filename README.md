# Go-Blockchain
## Go言語によるブロックチェーンシミュレーター
ブロックチェーンの学習を目的としたブロックチェーンのシミュレーターです

## 仕様
- Socket通信を用いたアプリケーション単位での通信でネットワークを形成します
- マイニングは１分おきに実行されます。全てのノードは同時にマイニングを開始します
- トランザクションはランダムに発生します
- ブロックには以下の情報が格納されています
   1. ブロック番号
   2. トランザクション
   3. タイムスタンプ
   4. 直前のブロックのハッシュ
   5. ナンス
   6. 署名（リスト）
   7. マイニングをしたノード名
- チェーン分岐時の処理は未実装（面倒臭かったから）
- 署名時のハッシュの再計算は行わない（面倒臭かったから）
- 第三者によるトランザクションの承認は行わない（公開鍵暗号とか面倒臭そうだったから）
- ブロック情報、ノード情報、トランザクション情報は全ノードで共有しています
- データの送受信にはJSONを使用しています

## 使用方法
#### エントリーポイントはblockchain.go

`Enter your name`

ノード名を入力してください

`Enter your port`

使用するポート番号を入力してください

`Are you first?`

一番最初のノードであれば1を入力

それ以外は2を入力してください


## 技術仕様
1. 使用言語
   - Go
2. 通信
   - トランスポート層(TCP)での通信(Socket通信)
   - UDPでマルチキャストするのも便利そうだけどパケットロスとか困りますし。。。
3. ハッシュ計算アルゴリズム
   - sha256(ビットコインと同じ)
   
## 解決すべき課題
 - ノード数が増えた際に全ノードが全ノードの情報を保持しているのは非常に非効率
    - ChordやKademliaなどのDHTアルゴリズムを用いて分散的にノードを管理すればいいかもしれない
    - Gossip Protocol(SWIM Protocol)を使うのが一般的かも
 - 並行処理について色々
    - 並行処理がマジ多い(パケットの送受信をしながらマイニングをする　等)
    - 並行処理について、Go言語では文法レベルでの強力なサポートがあるが、他の言語でも果たして同じようにプログラム出来るのかしら？
    - Shared memory型とMessage passing型の並行処理が混在してるので色々なタイミングで処理の競合が起きそうで危ない
    
## 参考リンク
[ブロックチェーンの基本的な仕組み](https://blockchain-jp.com/guides/4) ←ざっくり書いてあり、分かりやすい

[平成２７年度  我が国経済社会の情報化・サービス化に係る基盤整備
（ブロックチェーン技術を利用したサービスに関する国内外動向調査） 
報告書](http://www.meti.go.jp/press/2016/04/20160428003/20160428003-2.pdf) ←マイニングについて細かく書いてあり、分かりやすい


    
