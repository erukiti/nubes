command line parser
-------------------

* 今のところ決定打がないっぽい。
* https://github.com/codegangsta/cli よく使い方わからなかった
* https://github.com/jessevdk/go-flags なんか複雑すぎ
* 標準の flag 機能は最低限

flag ちょっと微妙なところあるけど最低限の機能あるならいいかという事で採用

Task interpreter
----------------

* 何らかのDSLが使えるのが望ましいが、Go標準にはそんな機能は無い
* cgo/mruby は、build が一手間増える・クロスコンパイルが面倒になる問題はある
* mrubyで、本家版とiij版があるが、iij版の方が機能が豊富

* V8とかPython慣れてない

pure go で書かれた処理系が何かあればいいのだが…。

cgo や mrubyの問題は今後解決していく事も期待して、慣れてるRuby処理系を採用


docker
------

* system("docker run -d hoge") みたいに CLI 経由でたたくのは、戻り値の取得とシェルエスケープが現状では面倒
* docker APIをたたく github.com/fsouza/go-dockerclient は cgo/mruby とのつなぎ込みがダルい部分はあるが、細かい操作が可能

docker API をたたくのは面倒が多いが今後解決していく事も期待して採用
