package main 

import "fmt"
// 「math/rand」パッケージと「time」パッケージをインポートしてください
import "math/rand"
import "time"
func main() {
    // timeのパッケージで使える完全なコードを生成するコード
    rand.Seed(time.Now().Unix())
    
		// for文を作成してください
		// 初期値、条件、変数の更新{繰り返したい処理}
    for i := 1;i <= 3; i++ {
		/* パッケージ、機能名で記述。機能の頭文字は大文字
			 %dは数値 %sは文字列 
			 Printf(出力する文字列、変数の値)*/
		fmt.Printf("%d回目のおみくじ結果:", i)
    number := rand.Intn(6)//0~5の数値をランダムに生成
        switch number {
            case 0:
                fmt.Println("凶です")
            case 1, 2:
                fmt.Println("吉です")
            case 3, 4:
                fmt.Println("中吉です")
            case 5:
                fmt.Println("大吉です")
        }
    }
}