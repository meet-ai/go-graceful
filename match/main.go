package main
import "github.com/alexpantyukhin/go-pattern-match"

func main() {
	n := 1
	_, res := match.Match(n).
		When(1, 1).
		When(2, 1).
		When(match.ANY, 3).
		Result()

	println(res.(int))
}
