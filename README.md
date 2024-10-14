# CLI APPLICATION

はい、その通りです。Go言語には`github.com/urfave/cli`というよく使われるCLIアプリケーション作成用のパッケージがあります。このパッケージを使用すると、より洗練されたCLIアプリケーションを簡単に作成できます。

`cli`パッケージを使用して、同じ機能を実装してみましょう。

まず、パッケージをインストールします：

```
go get github.com/urfave/cli/v2
```

次に、`cli`パッケージを使用してアプリケーションを作成します：



```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "openintextedit",
		Usage: "Open a file in TextEdit",
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.Exit("Please provide a filename", 1)
			}
			filename := c.Args().Get(0)
			return openInTextEdit(filename)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func openInTextEdit(filename string) error {
	cmd := exec.Command("open", "-a", "TextEdit", filename)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	fmt.Printf("Opened %s in TextEdit\n", filename)
	return nil
}

```

このコードの説明：

1. `github.com/urfave/cli/v2`パッケージをインポートしています。

2. `cli.App`構造体を使用してアプリケーションを定義しています。

3. `Action`フィールドにアプリケーションの主要なロジックを定義しています。

4. コマンドライン引数のチェックを行い、ファイル名が提供されていない場合はエラーメッセージを表示します。

5. `openInTextEdit`関数は以前のバージョンと同様の機能を持ちますが、エラーハンドリングが改善されています。

使用方法：

1. このコードを`main.go`として保存します。

2. 以下のコマンドでアプリケーションをビルドします：
   ```
   go build -o openintextedit
   ```

3. 生成された実行ファイルを使用してgreet.txtを開きます：
   ```
   ./openintextedit greet.txt
   ```

`cli`パッケージを使用することの利点：

1. コマンドライン引数の解析が簡単になります。

2. ヘルプメッセージが自動的に生成されます。`./openintextedit --help`を実行すると、使用方法が表示されます。

3. サブコマンド、フラグ、オプションなどの追加が容易です。

4. アプリケーションの拡張性が向上します。

今後の拡張の可能性：

1. ファイルが存在するかどうかのチェックを追加する。
2. 複数のファイルを開けるようにする。
3. `-e`や`--editor`フラグを追加して、使用するエディタを指定できるようにする。

このアプローチを使用すると、より堅牢で機能的なCLIアプリケーションを作成できます。必要に応じて機能を追加したり、エラーハンドリングを強化したりすることができます。

### video
<video width="320" height="240" controls>
  <source src="./video/cli.mov" type="video/quicktime">
  Your browser does not support the video tag.
</video>