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
- サーバーの起動
- エンドポイントの設定

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

### usecase/error.go
想定しているエラー型の定義
- 入力値が大きすぎる(ErrInvalidInput)
- 入力値が非負整数(ErrTooLargeInput)

### usecase/const.go
入力データの上限を設定現在(200000)

## interfaces
- 直接リクエストを受け取り、usecase層に渡す
- usecase層の返り値から、適切なレスポンスを返す

### interfaces/handler/fibIdx.go
1. URLから入力文字列をを抽出
2. usecase層に入力文字列を渡す
3. usecase層の返り値に応じてレスポンスを返す

> レスポンス一覧
> - 成功
>   - statuscode:200
>   - body: {"result" : (入力インデックスに応じたフィボナッチ数列の項)}
>
> - 入力値が大きすぎるエラー
>   - statuscode:400
>   - body:{"message" : "n is too large, please use less than (入力値の最大値)"}
> 
> - 入力値が非負整数でないエラー
>   - statuscode:400
>   - body:{"message" : "n must be a non-negative integer"}
>
> - 想定していないエラー
>   - statuscode:500
>   - body:{"message" : "failed to calculate Fibonacci number"}
>


