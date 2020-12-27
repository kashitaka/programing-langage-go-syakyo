package bank

var (
	sema    = make(chan struct{}, 1) // balance にアクセスする際はこのtokenを取得すること
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // トークンの獲得
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{} // トークンの獲得
	b := balance
	<-sema
	return b
}
