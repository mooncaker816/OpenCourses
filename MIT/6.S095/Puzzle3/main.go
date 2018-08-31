package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/mooncaker816/gophercises/poker/deck"
)

func main() {
	// deck := deck.New(deck.Shuffle)
	cards := randomDistinctN(5)
	// fmt.Println(cards)
	p1, p2, remains := getHiddenPair(cards)
	h, f, delta := decideHFAndDelta(p1, p2)
	remains = encodeDelta(delta, remains)
	var final []deck.Card
	final = append(final, f)
	final = append(final, remains...)
	fmt.Println("请根据以下4张牌猜测隐藏的第五张牌：")
	fmt.Println(final)
	var str string
	fmt.Scanf("%s\n", &str)
	guess, err := parseCard([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}
	if guess == h {
		fmt.Println("猜中了😃")
	} else {
		fmt.Println("没有猜中☹️")
	}
}

func randomDistinctN(n int) []deck.Card {
	var cards []deck.Card
	seen := make(map[int]struct{})
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	for n > 0 {
		num := rand.Intn(52) + 1
		if _, ok := seen[num]; ok {
			continue
		}
		seen[num] = struct{}{}
		cards = append(cards, abs1ToCard(num))
		n--
	}
	return cards
}

func abs1ToCard(n int) deck.Card {
	if n < 1 {
		panic("abs1 of card should >= 1")
	}
	if n == 53 {
		return deck.LittleJoker
	}
	if n == 54 {
		return deck.BigJoker
	}
	var card deck.Card
	card.Rank = deck.Rank((n-1)%13 + 1)
	card.Suit = deck.Suit((n - 1) / 13)
	return card
}

func getHiddenPair(cards []deck.Card) (f, s deck.Card, remains []deck.Card) {
	m := make(map[deck.Suit]int)
	for i, card := range cards {
		if idx, ok := m[card.Suit]; ok {
			remains = append(remains, cards[:idx]...)
			remains = append(remains, cards[idx+1:i]...)
			remains = append(remains, cards[i+1:]...)
			return cards[idx], card, remains
		}
		m[card.Suit] = i
	}
	return
}

func decideHFAndDelta(p1, p2 deck.Card) (h, f deck.Card, delta int) {
	delta = mod(int(p1.Rank)-int(p2.Rank), 13)
	if delta <= 6 {
		h = p1
		f = p2
	} else {
		h = p2
		f = p1
		delta = (mod(int(p2.Rank)-int(p1.Rank), 13))
	}
	return
}

func encodeDelta(delta int, remains []deck.Card) []deck.Card {
	if len(remains) != 3 {
		panic("encoding cards' number is not 3!")
	}
	sort.Slice(remains, func(i, j int) bool {
		if remains[i].Rank != remains[j].Rank {
			return remains[i].Rank < remains[j].Rank
		}
		return remains[i].Suit < remains[j].Suit
	})
	var encode []deck.Card
	switch delta {
	case 1:
		encode = append(encode, remains[:]...)
	case 2:
		encode = append(encode, remains[0], remains[2], remains[1])
	case 3:
		encode = append(encode, remains[1], remains[0], remains[2])
	case 4:
		encode = append(encode, remains[1], remains[2], remains[0])
	case 5:
		encode = append(encode, remains[2], remains[0], remains[1])
	case 6:
		encode = append(encode, remains[2], remains[1], remains[0])
	}
	return encode
}

func mod(m, n int) int {
	r := m % n
	if r < 0 {
		r += n
	}
	return r
}

func parseCard(src []byte) (deck.Card, error) {
	var card deck.Card
	if string(src) == "BigJoker" {
		return deck.BigJoker, nil
	}
	if string(src) == "LittleJoker" {
		return deck.LittleJoker, nil
	}
	_, size := utf8.DecodeRune(src)
	switch string(src[:size]) {
	case "♠":
		card.Suit = deck.Spade
	case "♥":
		card.Suit = deck.Heart
	case "♣":
		card.Suit = deck.Club
	case "♦":
		card.Suit = deck.Diamond
	default:
		return card, fmt.Errorf("suit not supported: %s", src[:size])
	}
	src = src[size:]
	if len(src) == 1 {
		switch src[0] {
		case 'J':
			card.Rank = deck.Jack
			return card, nil
		case 'Q':
			card.Rank = deck.Queen
			return card, nil
		case 'K':
			card.Rank = deck.King
			return card, nil
		case 'A':
			card.Rank = deck.Ace
			return card, nil
		}
	}
	num, err := strconv.Atoi(string(src))
	if err != nil {
		return card, fmt.Errorf("can not convert card rank %s:%v", string(src), err)
	}
	if num <= 0 || num > 13 {
		return card, fmt.Errorf("rank not supported: %d", num)
	}
	card.Rank = deck.Rank(num)
	return card, nil
}
