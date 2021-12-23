package main

import (
	"GoClashRoyale/api"
	"encoding/json"
	"github.com/alessiosavi/GoGPUtils/helper"
	"io/ioutil"
	"log"
)

type Conf struct {
	BaseURL   string `json:"base_url"`
	PlayerTag string `json:"player_tag"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
}

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds)

	// Read the configuration and load into a typed struct
	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}
	var conf Conf
	if err = json.Unmarshal(file, &conf); err != nil {
		panic(err)
	}

	// Generate a new API if necessary or retrieve the last one created
	token, err := api.NewKey(conf.Mail, conf.Password)
	if err != nil {
		panic(err)
	}

	// Set the base URL
	api.BASE_URL = conf.BaseURL

	// Retrieve the player information
	getPlayer, err := api.GetPlayer(conf.PlayerTag, token)
	if err != nil {
		panic(err)
	}
	log.Println("============= Player Info =============")
	log.Println(helper.MarshalIndent(getPlayer))

	// Retrieve the next chest that you will receive
	chest, err := api.GetUpcomingChest(conf.PlayerTag, token)
	if err != nil {
		panic(err)
	}
	log.Println("============= Upcoming Chest =============")
	log.Println(helper.MarshalIndent(chest))

	// Retrieve the latest battle result
	battleLog, err := api.GetBattleLog(conf.PlayerTag, token)
	if err != nil {
		panic(err)
	}
	log.Println("============= Battle Log =============")
	log.Println(helper.MarshalIndent(battleLog))

	// Retrieve the deck that win against you
	lost := api.GetLost(conf.PlayerTag, token)
	log.Println("============= Decks Lost =============")
	log.Println(helper.MarshalIndent(lost))
}
