package main

import "fmt"

/*コンピュータにはメモリと呼ばれる作業場所のようなものが存在します。
変数はそのメモリに記録されており、その場所をアドレス(16進数)と呼びます。
アドレスは値です。整数や文字列のように操作したり変数に代入することができます。Goではアドレスを(ポインタ)と呼んで扱っています。
またポインタが代入された変数をポインタ型変数といいます。
ポインタ型変数を定義するには、変数のデータ型に「*(アスタリスク)」をつけて宣言します。*/
func main() {
	totalScore := 0
	// 引数にtotalScoreのポインタを渡してください アドレスの取得(&変数名)
	ask(1, "dog",&totalScore)
	ask(2, "cat",&totalScore)
	ask(3, "fish",&totalScore)

	fmt.Println("スコア", totalScore)
}

// 渡されるtotalScoreのポインタを受け取るように変更してください
// scorePtr *int はポインタ型変数
func ask(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("正解!")
		// ポインタを使って加算してください
		// ポインタを使用してmainの変数の値を更新することができる。
		*scorePtr += 10
	} else {
		fmt.Println("不正解!")
	}
}