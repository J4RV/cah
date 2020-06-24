package cah

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_ExtractCardFromHand(t *testing.T) {
	p := Player{}
	c1 := WhiteCard{Text: "A"}
	c2 := WhiteCard{Text: "B"}
	c3 := WhiteCard{Text: "C"}
	p.Hand = []WhiteCard{c1, c2, c3}

	_, err := p.ExtractCardFromHand(-1)
	if err == nil {
		t.Error("Expected error but did not found it, negative index")
	}

	_, err = p.ExtractCardFromHand(9)
	if err == nil {
		t.Error("Expected error but did not found it, index over hand size")
	}

	c, err := p.ExtractCardFromHand(1)
	assert.Equal(t, c.Text, "B", "Unexpected text in extracted hand")
	assert.Equalf(t, err, nil, "Unexpected error %v", err)
	assert.Equalf(t, len(p.Hand), 2, "Hand size did not get reduced, hand: %s, len: ", p.Hand)
	expectedResultHand := []WhiteCard{c1, c3}

	for i := range p.Hand {
		assert.Equalf(t, p.Hand[i], expectedResultHand[i], "Unexpected hand card at position %d", i)
	}
}

func TestPlayer_ExtractCardsFromHand(t *testing.T) {
	assert := assert.New(t)
	p := Player{}
	p.Hand = getWhiteCardsFixture(10)
	indexes := []int{0, 9, 5}
	cards, err := p.ExtractCardsFromHand(indexes)

	assert.NoError(err)
	assert.Equal(7, len(p.Hand), "Unexpected hand size")
	assert.Equal(3, len(cards), "Unexpected cards size")

	for i, index := range indexes {
		assert.True(strings.Contains(cards[i].Text, fmt.Sprintf("(%d)", index)),
			"Unexpected card order in cards result", cards[i].Text)
	}

	expectedHand := []int{1, 2, 3, 4, 6, 7, 8}
	for i, index := range expectedHand {
		assert.True(strings.Contains(p.Hand[i].Text, fmt.Sprintf("(%d)", index)),
			"Unexpected card order in hand", p.Hand[i].Text)
	}
}

func getWhiteCardsFixture(amount int) []WhiteCard {
	ret := make([]WhiteCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = WhiteCard{Text: fmt.Sprintf("White card fixture (%d)", i)}
	}
	return ret
}
