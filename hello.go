package somemodule

import (
	quote "rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quote.Hello()
}

func HelloWithVersion() string {
	return Hello() + " v1.2.0"
}

func Proverb() string {
	return quoteV3.Concurrency()
}
