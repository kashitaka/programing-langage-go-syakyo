package eval

import (
	"fmt"
	"math"
)

// 式 exression
type Expr interface {
	Eval(env Env) float64
}

// 変数 例: x
type Var string

// 数値定数 例: 3.14
type literal float64

// 変数と値の対応 "pi" = 3.14
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// 単行演算子
type unary struct {
	op rune // + or -
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// 二項演算
type binary struct {
	op   rune // +, -, *, / のどれか
	x, y Expr
}

func (u binary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return u.x.Eval(env) + u.y.Eval(env)
	case '-':
		return u.x.Eval(env) - u.y.Eval(env)
	case '*':
		return u.x.Eval(env) * u.y.Eval(env)
	case '/':
		return u.x.Eval(env) / u.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// 関数呼び出し
type call struct {
	fn   string // pow, sin, sqrtのどれか
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
