# fib_api

## 使用技術
- 言語:Go
- フレームワーク:echo
- デプロイ:Heroku

## 構成


### domain
*** domain/fibIdx.go ***
役割：ドメインの定義
今回のドメイン：*** フィボナッチ数列のインデックス ***
ドメインが持つ値 : FibIdx(uint64)

*** domain/Repository/fibIdx_repository.go ***
役割：ドメインが持つ関数の定義
ドメインが持つ関数：
CalcFibNum
引数：fibIdx(*domain.fibIdx)
返り値：fibIdxに対応したfibonacci数(uint64)

### usecase
*** usecase/fibIdx.go ***
役割：ドメインが持つビジネスロジックの定義
関数:
CalcFibNum

バリデーション：
- 値の形がUintに変換できるか？
  - 負の値？
  - 数値以外の値？
  - 少数？
  - 文字列？
  - Uintでも表現できないレベルの大きさ？(どうなる？)
- 正しい結果を返せる値か？(引数が大きすぎて結果をUintで表現できないレベルか？)←これは後で調べてconstにいれる
キャッシュで引っ張ってくるのあり？←なし

アルゴリズム(計算量：O(N))：
1. 一つ前の結果をlast、2つ前の結果をlast2として保存
2. result = last + last2で項を計算
3. lastにresultを、last2にlastを代入

### interfaces
ハンドラを定義
想定するレスポンス

   
| ステータスコード | レスポンス(Json) | 想定しているケース|
| ---- | ---- |  ---- |
| 200 | {"result" : 対応するフィボナッチ数列の値} | 成功 |
| 400 | {"message" : "invalid value"} |  値の形式が適切ではない(負の数や文字列など) |
| 400 | {"message" : "too large value"} |  値が大きすぎる |

