package bank

var deposits = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高を送信する
var withdraw = make(chan int)
var withdrawRes = make(chan bool)

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	withdraw <- amount
	return <-withdrawRes
}

func teller() {
	var balance int // 残高の管理変数
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraw:
			if balance < amount {
				withdrawRes <- false
			} else {
				balance -= amount
				withdrawRes <- true
			}
		case balances <- balance:
			// select case の 送信では channelに空きがある場合に送信が実行される
			// このselectの動きとしては
			// 1. depositsのチャネルをチェック。あれば取得。なければ次
			// 2. balances に空きがあれば残高を送信。空きがなければ何もしない
		}
	}
}

func init() {
	go teller()
}
