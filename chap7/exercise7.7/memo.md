```
flag.CommandLine.Var(&f, name, usage)
```

を呼ぶと Flag.go の中でFlagSetを作るときに

```
func (f *FlagSet) Var(value Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()}
```

となり、第4引数に `value` の `String()` が返す値を入れている。（ここでは 20℃が入る）

Flagの構造体は下記.

```
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```

なので `DefValue` にセットされて、 `PrintDefault` の際に

```
fmt.Sprintf(" (default %q)", flag.DefValue)
```

で出力