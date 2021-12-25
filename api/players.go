package api

import (
	"GoClashRoyale/common"
	"encoding/json"
	"fmt"
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
	res.FixLevel()
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
	res.FixLevel()
	return &res, nil
}

func UniqueMatrix(decks [][]Card) [][]Card {
	var result [][]Card
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

func Equals(v1, v2 []Card) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i].Name != v2[i].Name || v1[i].Level != v2[i].Level {
			return false
		}
	}
	return true

}
