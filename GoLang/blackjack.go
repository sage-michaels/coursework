package main
import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

//method for getting user input
func input(prompt string) string {
	var s string = ""
	fmt.Printf("%s", prompt)
	fmt.Scanln(&s)
	return s
}

//create list of cards
func make_deck(x []int) []int {
	var i int = 0
	for i < 52 {
		x[i]=i
		i = i + 1
	}
	return x
}

func shuffle_deck(x []int) []int{
 var i int=51
 for i>0 {
 	rand.Seed(time.Now().UnixNano())
 	var v=rand.Intn(i)
 	x[i],x[v] = x[v],x[i]
 	i=i-1
 }
 return x //this uses the fisher-yates shuffleing algorithm
}
func draw(x []int) []int{
	var n int=len(x)
	return x[0:(n-1)]
}
func add_to_hand(hand []int, n int) []int{
	var i int = 0
	for hand[i]!=0{
		i=i+1
	}
	hand[i]=n
	return hand
}

func sum(a []int) int {
	var s int = 0
	var i int = 0
	var n int = len(a)
	for i < n {
		s = s + a[i]
		i = i + 1
	}
	return s
}
func alt_sum(a []int) int {
	var s int = 0
	var i int = 0
	var n int = len(a)
	for i < n {
		if a[i]==1{
			s = s + 11
			s = s + sum(a[i+1:])
			i = n
		}else{
		 s = s + a[i]
		 i = i + 1
		}
	}
	return s
}
func sums(x []int) int{
	var s_1 int = sum(x)
	var s_2 int = alt_sum(x)
	if s_1==s_2 || s_2>21{
		fmt.Printf(" total is %d\n",s_1)
		return s_1
	}else{
		fmt.Printf(" total is either \"a\"->(%d) or \"b\"->(%d)\n",sum(x),alt_sum(x))
		if s_2 == 21{
			return 21
		}else{
			return s_2
		}
	}
}
func sums_quiet(x []int) int{
	var s_1 int = sum(x)
	var s_2 int = alt_sum(x)
	if s_1==s_2 || s_2>21{
		return s_1
	}else{
		if s_2 == 21{
			return 21
		}else{
			return s_2
		}
	}
}
func blackjack_bust(n int) bool{
	if n == 21{
		return false
	} else if n>21 {
		return false
	} else{
		return true
	}
} 

func read_suit(x int) string{
	var suit string
	if x/13 == 0{
			suit = "spades"
	} else if x/13 == 1{
		 	suit = "clubs"
	} else if x/13 == 2{
		 	suit = "diamonds"
	} else{
		 suit = "hearts"
 }
 return suit
}

func read_card(x int) int{
	var i int= 2
	var n string 
	for i <=10{
		if x%13==i{
			break
		}else{
			i=i+1
		}
	}
	if i<=10{
		n=strconv.Itoa(i)
	}else if x%13==0{
		i=1
		n="ace"
	}else if x%13==1{
		n="jack"
		i=10
	}else if x%13==11{
		n="queen"
		i=10
	}else{
		n="king"
		i=10
	}
	fmt.Printf("%s of %s ~~~\n",n,read_suit(x))
	return i
}
func read_card_quiet(x int) int{
	var i int= 2
	for i <=10{
		if x%13==i{
			break
		} else{
			i=i+1
		}
	}
	if x%13==0{
		i=1
	} else if x%13==1{
		i=10
	} else if x%13==11{
		i=10
	}else{
		i=10
	}
	return i
}

func main() {
//setup game: make and shuffle deck, deal 2 cards to player and 2 cards (one hidden) to dealer
 var deck []int
 deck = make([]int,52)
 deck = make_deck(deck)
 deck = shuffle_deck(deck)
	var p_deck []int= make([]int,10)
	var h_d_deck []int= make([]int,1)
	var s_d_deck []int= make([]int,10)
	var hh_d_deck []int = make([]int, 1)
	fmt.Printf("player drew a ~~~ ")
 p_deck = add_to_hand(p_deck, read_card(deck[len(deck)-1]))
	deck=draw(deck)
	fmt.Printf("dealer drew a hidden card\n")
	hh_d_deck[0] = deck[len(deck)-1]
	h_d_deck =add_to_hand(h_d_deck, read_card_quiet(deck[len(deck)-1]))
	deck=draw(deck)
	fmt.Printf("dealer drew a ~~~ ")
	s_d_deck = add_to_hand(s_d_deck, read_card(deck[len(deck)-1]))
	s_d_deck = add_to_hand(s_d_deck, h_d_deck[0])
	deck=draw(deck)
	fmt.Printf("player drew a ~~~ ")
	p_deck= add_to_hand(p_deck, read_card(deck[len(deck)-1]))
	deck=draw(deck)
	fmt.Printf("Player")
	var p_sum int = sums(p_deck)
	var d_sum int = sums_quiet(s_d_deck)
	var cont bool = blackjack_bust(p_sum) && blackjack_bust(d_sum)
	var turn string
	if cont{
		turn = input("\"hit\" or \"stay\"?\n")
	}else{
		turn = ""
	}

	//while player inputs "hit" and has a card total < 21 continue game
	//end players turn when they input "hit" or have a card total over 20
	for "hit"==turn && cont && 21>p_sum {
			fmt.Printf("player drew a ~~~ ")
			p_deck = add_to_hand(p_deck, read_card(deck[len(deck)-1]))
			deck=draw(deck)
			p_sum = sums(p_deck)
			cont = blackjack_bust(p_sum) || blackjack_bust(d_sum)
			if p_sum>=21{
				turn = "autostay"
			}else{
			    turn = input("\"hit\" or \"stay\"? \n")
			}
	}
	//dealer keeps hitting while their total is under 17, then dealers hidden card is revealed
	for sums_quiet(s_d_deck)<17 && cont {
	     fmt.Printf("dealer drew a ~~~ ")
	 s_d_deck = add_to_hand(s_d_deck, read_card(deck[len(deck)-1]))
	 deck = draw(deck)
	 cont = blackjack_bust(p_sum) || blackjack_bust(d_sum)
	}
	fmt.Printf("Dealer's hidden card was %d \n",read_card_quiet(hh_d_deck[0]))
	p_sum = sums_quiet(p_deck)
	d_sum = sums_quiet(s_d_deck)

	//print end of game senario (win, tie, loss, blackjack)
	if p_sum>21 {
		if d_sum>21{
			fmt.Printf("Tie! You both busted")
		}else {
			fmt.Printf("Dealer won, you busted")
		}	
	}else if p_sum == 21{
		if d_sum == 21{
			fmt.Printf("Tie! You both got blackjack!")
		} else{
			fmt.Printf("You win! You got a blackjack!")
		}
	}else if d_sum>21{
		fmt.Printf("The dealer busted! You Won!")
	}else{
		if d_sum>p_sum{
			fmt.Printf("Dealer won with %d to your %d",d_sum,p_sum)
		}else if d_sum==p_sum{
			fmt.Printf("You both tied at %d",d_sum)
		}else{
			fmt.Printf("You won with %d to the dealer's %d",p_sum,d_sum)
		}
	}
	fmt.Printf("\n ~~~Game Over~~~\n")


	//fmt.Printf("player hand %v\n",p_deck)
	//fmt.Printf("showing dealer hand %v\n",s_d_deck)
	//fmt.Printf("hidden dealer hand %v\n",h_d_deck)
}