package main

import "fmt"
/* main関数はプログラムに最初に呼び出される特別な関数 記述順は関係なし */
func main() {
	/* askの関数を書いた分だけ実行される mainに記述しないと実行されない。
	エラー表示される。定義して使わない場合も同様 */
	/* 引数の順番は同じにしないとエラーが起きる */
	totalScore := ask(1, "dog")
	totalScore += ask(2, "cat")
	totalScore += ask(3, "fish")

	fmt.Println("スコア", totalScore)
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	/* fmt.Scan(&変数名)はターミナルに入力 */
	fmt.Scan(&input)

	if question == input {
	    // 問題文の単語と入力値が一致するときの処理を追加してください
	    fmt.Println("正解!")
	    return 10
	} else {
	    // 問題文の単語と入力値が一致しないときの処理を追加してください
	    fmt.Println("不正解!")
	    return 0
	}
}
