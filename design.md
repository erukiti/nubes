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

* system("docker run -d hoge") みたいに CLI 経由でたたくのはシェルエスケープが現状では面倒
* docker APIをたたく github.com/fsouza/go-dockerclient は cgo/mruby とのつなぎ込みがダルい部分はあるが、細かい操作が可能

daemon化
--------

* go標準のdaemon化方法はない
* syscallを使ったOS依存ありなやり方はある
* あるいはos.ProcessStart を使った擬似的なやり方もある

serf/consul/docker などは、daemon化は行ってない。
nubes でも面倒なので daemon 化はしない方向で…

cron
----

* github.com/gorhill/cronexpr はライセンス的に不採用
* github.com/robfig/cron は一部を利用する感じで

