# bingSpellCheck
A GOLANG client for Bing Spell Check API version 7

One Saturday morning I had nothing better to do. You are welcome.

## Getting Started

1. Download/Install it

```sh
$ go get github.com/gotomgo/bingSpellCheck
```

2. In examples/main.go add your Bing Spell Check API key

```go
const key = "<INSERT YOUR API KEY HERE>"
```

3. Run the example:
```sh
$ go run example/*.go "Is teh teh data good to go?"
```

4. Review the example for basic usage

```go
client := bingSpellCheck.NewClient(key)

spellCheck, err := client.SpellCheck(os.Args[1])
if err != nil {
  fmt.Println(err)
} else {
  spew.Dump(spellCheck)
}

correctedText, err := bingSpellCheck.BuildAutoCorrectedText(os.Args[1], spellCheck)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(correctedText)
}
```
