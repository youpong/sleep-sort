# README.md

オリジナルは sat さんによる sleep sort(※1) の動画から起こした sleep-sort.go。

## 修正点
1. 整列の要素に負数が含まれた際にエラーを返却する
2. テスト sleep-sort_test.go を追加

要素数が多数となるテスト TestTail2Head() を含む。私の環境(※2)では 
(128+12)k あたりを超える要素をソートする場合にエラーがしばしば発生した。

## How to test

```bash
$ go test -parallel 4
```
sleep が処理時間に及ぼす影響が大きいテストプログラムなので、 CPU の個数を超え
て並行処理した方がパフォーマンスが出る。
テストの並行化については、資料(※3)を参考にした。

## Footnote

※1) 07-Sep-22 satlinuxtube(武内 覚),
THE FIRST CODE 寝てたら終わるスリープソート(https://youtu.be/PSeIEBPnq-E)

※2) Intel Core 2 Duo T9400 @ 2.53GHz

※3) yoshikishibata, [Go言語でのテストの並列化 〜t.Parallel()メソッドを理解する〜](https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/), mercari engineering
