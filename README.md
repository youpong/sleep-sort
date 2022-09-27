README.md

sat さんによる sleep sort(※) についての要素数が多くなった場合のテスト。私の環境では 256k の
要素をソートする場合にエラーが発生した。
TestTail2Head() あるいは TestHead2Tail() の冒頭にあるスライス sample の要素数を調整して
試してください。

How to test

```
$ go test -v
```

※)07-Sep-22 satlinuxtube(武内 覚)
THE FIRST CODE 寝てたら終わるスリープソート(https://youtu.be/PSeIEBPnq-E)
