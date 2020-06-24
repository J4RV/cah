package cah

import (
	"errors"
	"fmt"
	"sort"
)

type Player struct {
	User             User        `json:"user" db:"user"`
	Hand             []WhiteCard `json:"hand" db:"hand"`
	WhiteCardsInPlay []WhiteCard `json:"whiteCardsInPlay"`
	Points           []BlackCard `json:"points" db:"points"`
}

func NewPlayer(u User) *Player {
	return &Player{
		User:             u,
		Hand:             []WhiteCard{},
		WhiteCardsInPlay: []WhiteCard{},
		Points:           []BlackCard{},
	}
}

func (p *Player) RemoveCardFromHand(i int) error {
	if i < 0 || i >= len(p.Hand) {
		msg := fmt.Sprintf("Index out of bounds. Index: %d, Hand size: %d", i, len(p.Hand))
		return errors.New(msg)
	}
	p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
	return nil
}

func (p *Player) ExtractCardsFromHand(indexes []int) ([]WhiteCard, error) {
	ret := make([]WhiteCard, len(indexes))

	for iter, index := range indexes {
		if index < 0 || index >= len(p.Hand) {
			return nil, fmt.Errorf("Non valid white card index: %d", index)
		}
		c := p.Hand[index]
		ret[iter] = c
	}

	// Order ints from highest to lowest to prevent indexes out of bounds
	// since we will be altering the hand slice by removing cards from it
	iOrdered := make([]int, len(indexes))
	copy(iOrdered, indexes)
	sort.Sort(sort.Reverse(sort.IntSlice(iOrdered)))
	for _, index := range iOrdered {
		err := p.RemoveCardFromHand(index)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (p *Player) ExtractCardFromHand(i int) (WhiteCard, error) {
	ret, err := p.ExtractCardsFromHand([]int{i})
	if err != nil {
		return WhiteCard{}, err
	}
	return ret[0], nil
}
