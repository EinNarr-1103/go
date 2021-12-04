## go でマイス操作を自動化

### バイナリファイルの実行方法

### マウスを自動操作させる時間を変更したい

0. デフォルトは 5 分

1. main.go を編集

```main.go
func main() {
  fmt.Print("Mouse hack processing... \n")
  // ここの数を変更
  mouseMoveByMinute(5)
}
```

2. 指定したディレクトリにバイナリファイルを出力する
   - `go build -o [出力先のパス] [対象のgoファイルのパス]`
