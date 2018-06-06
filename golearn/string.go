package golearn

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func StringExample() {
	utf8Example()
	//stringExample()
}

func utf8Example() {
	const nihongo = "日本語ABC"
	const nihongo2 = "日本語\xda\xce"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d---sizeof:%d\n", runeValue, index, unsafe.Sizeof(nihongo))
	}
	for index, runeValue := range nihongo2 {
		fmt.Printf("%#U starts at byte position %d--sizeof:%d\n", runeValue, index, unsafe.Sizeof(runeValue))
	}

	const nihongo3 = "日本語ABC"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
}

func stringExample() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println("")

	for i := 0; i < len(sample); i++ {
		fmt.Print(sample[i])
	}
	fmt.Println("")

	fmt.Printf("%x\n", sample)

	fmt.Printf("% x\n", sample)

	fmt.Printf("%q\n", sample)

	fmt.Printf("%+q\n", sample)

	// Exercise 1
	b := []byte(sample)
	fmt.Println(b)

	for i := 0; i < len(b); i++ {
		fmt.Printf("%x ", b[i])
	}
	fmt.Println("")

	fmt.Printf("%x\n", b)

	fmt.Printf("% x\n", b)

	fmt.Printf("%q\n", b)

	fmt.Printf("%+q\n", b)

	// Exercise 2
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", b[i])
	}
}
