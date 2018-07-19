## About go

Go is not a OOP

### Variables

- Static types.
- some basic types: bool, string, int, float64.
- we use := to difined a new variable, only new.
 card:= "Ace of Spades"
 then => 
 card = "Five of Diamonds"

 ### function declarations

```
 fucn newCard() string {

 }
 ```

 func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

 we have to tell what type of value we are going to return.

 ### Array and slice

 array: fixed length list of things
 
 slice: an array that can grow or shirnk, every element must be the same type.
cards := []string{newCard(), newCard()}

to add an element

cards = append(cards, "Six of Spades")

returns a new slice.

### loops

for i, card := range cards {
	fmt.Println(i, card)
}

if we replace i with _ we are telling go that we are not goint to use the index.



