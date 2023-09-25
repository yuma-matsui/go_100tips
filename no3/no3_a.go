package no3

import "fmt"

// 最初に実行される
var a = func() int {
	fmt.Println("var")
	return 0
}()

// 2番目に実行される
func init() {
	fmt.Println("init")
}

// 最後に実行される
func main() {
	fmt.Println("main")
}

/* メモ
* init関数には引数も返り値も与えない

* mainパッケージが依存するサブパッケージにinit関数がある場合、
	サブパッケージのinit関数が先に実行される
	subPkgのinit -> mainPkgのinit -> main関数

* mainパッケージから"複数"のサブパッケージに依存していて、
	それぞれのサブパッケージにinit関数がある場合、
	サブパッケージのファイル名のアルファベット順でinit関数が実行される

* 同一ファイル内で宣言されたinit関数は上から順に実行される

* init関数は直接呼び出せない
*/
