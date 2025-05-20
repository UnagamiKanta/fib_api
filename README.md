<script type="text/javascript" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/3.2.2/es5/tex-mml-chtml.min.js">
</script>
<script type="text/x-mathjax-config">
 MathJax.Hub.Config({
 tex2jax: {
 inlineMath: [['$', '$'] ],
 displayMath: [ ['$$','$$'], ["\\[","\\]"] ]
 }
 });
</script>
# fib_api

## 使用技術
- 言語:Go
- フレームワーク:echo
- デプロイ:Heroku

## 構成

## main.go
- サーバーを起動
- エンドポイント設定

## domain
ドメインロジックの定義
### domain/fibIdx.go
今回のドメイン：フィボナッチ数列計算
| 関数 | 引数 | 返り値|
| ---- | ---- |  ---- |
| CalcFibNum | FibIdx(big.Int) | big.Int, error |

フィボナッチ数列の一般項$F_n$は行列[[1, 1], [1, 0]]^nの対角成分であることを用いた

### domain/fibIdx_helper.go
ドメインロジックのためのヘルパー関数を記述
| 関数 | 引数 | 返り値|役割|
| ---- | ---- |  ---- | ---- |
| MatrixMul | a([2][2]big.Int), b([2][2]big.Int) | [2][2]big.Int | 2×2行列同士の掛け算を定義|
| MatrixPow | m([2][2]big.Int), n(big.Int) | [2][2]big.Int | 行列のべき乗を計算、計算量O($\log n$)|

MatrixPowは繰り返し二乗法を用いることでO($\log n$)の計算量にしている
[参考にしたサイト](https://qiita.com/ophhdn/items/e6451ec5983939ecbc5b)

### domain/fibIdx_test.go
CalcFibNumのテストコード$n\in [0, 6]$の各$n$まで実装


## usecase
役割
- domainロジックに入力値を渡し、返り値をハンドラーに返す
- 入力値のエラーハンドリング
  
> 想定しているエラー(エラー型)
> - 入力値が大きすぎる(ErrInvalidInput)
> - 入力値が非負整数(ErrTooLargeInput)

### usecase/fibIdx.go
1. 入力文字列データの長さでエラーハンドリング(長過ぎる文字列をbigIntに変換できないため)
2. 入力文字列データをbigIntに型変換(エラーの際はErrInvalidInput)
3. 入力データが大きすぎないかチェック(ErrInvalidInput)
4. domain層のフィボナッチ数列を計算
5. handlerに3の計算結果を返す

### usecase/fibIdx_test.go
usecase/fibIdx.goのテスト
想定するテストケース
- 正常な入力
- 非負整数ではない入力
- 大きすぎる入力
- bigInt型に変換できないほど大きすぎる入力

### usecase/error.go
想定しているエラー型の定義
- 入力値が大きすぎる(ErrInvalidInput)
- 入力値が非負整数(ErrTooLargeInput)

### usecase/const.go
入力データの上限を設定現在(200000)

## interfaces
役割
- リクエスト、レスポンスの受け渡し

### interfaces/handler/fibIdx.go
リクエスト
`/fib?n=10`, `GET`

レスポンス
| ステータス | レスポンス | 詳細|
| ---- | ---- |  ---- |
| 200 | `{"result" : (フィボナッチ数)}` | 正常な入出力 |
| 400 | `{"message" : "n must be a non-negative integer"}` | 非負整数以外の入力 |
| 400 | `{"message" : "n is too large, please use less than (最大値)"}` | 大きすぎる入力 |
| 500 | `{"message" : "failed to calculate Fibonacci number"}` | 予期しないエラー |

### interfaces/handler/fibIdx_test.go
上記レスポンスのステータスコード500以外のものをテスト

