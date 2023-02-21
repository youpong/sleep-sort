# README.md

sat さんによる sleep sort(※1) の動画から起こした sleep-sort.go。

さらに、テスト sleep-sort_test.go を独自に追加している。
ここでは要素数が多くなった場合のテスト TestTail2Head() を含む。私の環境(*2)では 
(128+12)k あたりを超える要素をソートする場合にエラーがしばしば発生した。

## How to test

```bash
$ go test -parallel 4
```
sleep が処理時間に及ぼす影響が大きいテストプログラムなので、 CPU の個数を超え
て並行処理した方がパフォーマンスが出る。

## Footnote

※1)07-Sep-22 satlinuxtube(武内 覚),
THE FIRST CODE 寝てたら終わるスリープソート(https://youtu.be/PSeIEBPnq-E)

*2) Intel Core 2 Duo T9400@2.53GHz
