package api

import (
	"GoClashRoyale/common"
	"encoding/json"
	"fmt"
	"github.com/alessiosavi/GoGPUtils/helper"
	"sort"
)

var BASE_URL = ""

func GetPlayer(playerTag, token string) (*Player, error) {
	request, err := common.InitRequest(fmt.Sprintf("%s/players/%s", BASE_URL, playerTag), token)
	if err != nil {
		return nil, err
	}
	data, err := common.FireRequest(request)
	if err != nil {
		return nil, err
	}

	var res Player
	err = json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}
	return &res, nil
}
func GetUpcomingChest(playerTag, token string) (*Chests, error) {
	request, err := common.InitRequest(fmt.Sprintf("%s/players/%s/upcomingchests", BASE_URL, playerTag), token)
	if err != nil {
		return nil, err
	}
	data, err := common.FireRequest(request)
	if err != nil {
		return nil, err
	}

	var res Chests
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
func GetBattleLog(playerTag, token string) (*Battlelog, error) {
	request, err := common.InitRequest(fmt.Sprintf("%s/players/%s/battlelog", BASE_URL, playerTag), token)
	if err != nil {
		return nil, err
	}
	data, err := common.FireRequest(request)
	if err != nil {
		return nil, err
	}

	var res Battlelog
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetLost(playerTag, token string) [][]string {
	battleLog, err := GetBattleLog(playerTag, token)
	if err != nil {
		panic(err)
	}

	var decks [][]string
	for _, battle := range *battleLog {
		if battle.Type == "PvP" && !battle.BoatBattleWon {
			if len(battle.Opponent) != 1 {
				panic("The following battle have more than 1 opponent! " + helper.MarshalIndent(battle))
			}
			var deck []string
			for _, card := range battle.Opponent[0].Cards {
				deck = append(deck, card.Name)
			}
			decks = append(decks, deck)
		}
	}
	for i := range decks {
		sort.Strings(decks[i])
	}
	decks = append(decks, decks[0])
	decks = UniqueMatrix(decks)

	return decks
}

func UniqueMatrix(decks [][]string) [][]string {
	var result [][]string
	for i := 0; i < len(decks)-1; i++ {
		isEqual := false
		for j := i + 1; j < len(decks); j++ {
			if Equals(decks[i], decks[j]) {
				isEqual = true
				break
			}
		}
		if !isEqual {
			result = append(result, decks[i])
		}
	}
	return result
}

func Equals(v1, v2 []string) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true

}
