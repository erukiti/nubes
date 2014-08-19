command line parser
-------------------

* 今のところ決定打がないっぽい。
* https://github.com/codegangsta/cli よく使い方わからなかった
* https://github.com/jessevdk/go-flags なんか複雑すぎ
* flag 微妙だけど、最低限の機能そろってるからこれでいいや

って事で、標準の flag を使うことにした。

Task interpreter
----------------

* 何らかのDSLが使えるのが望ましいが、Go標準にはそんな機能は無い
* cgo/mruby は、build が一手間増える・クロスコンパイルが面倒になる問題はある
