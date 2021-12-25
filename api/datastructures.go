package api

import (
	"sort"

	"github.com/alessiosavi/GoGPUtils/helper"
)

type Battle struct {
	MyDeck   []Card `bson:"my_deck" json:"my_deck"`
	Opponent []Card `bson:"opponent" json:"opponent"`
	Win      bool   `bson:"win" json:"win"`
}

type Card struct {
	Level int    `bson:"level" json:"level"`
	Name  string `bson:"name" json:"name"`
}

func (b *Battle) SortDecks() {
	sort.Slice(b.MyDeck, func(i, j int) bool {
		if b.MyDeck[i].Level == b.MyDeck[j].Level {
			return b.MyDeck[i].Name < b.MyDeck[j].Name
		}
		return b.MyDeck[i].Level > b.MyDeck[j].Level
	})
	sort.Slice(b.Opponent, func(i, j int) bool {
		if b.Opponent[i].Level == b.Opponent[j].Level {
			return b.Opponent[i].Name < b.Opponent[j].Name
		}
		return b.Opponent[i].Level > b.Opponent[j].Level
	})
}

type Deck []struct {
	Count int `bson:"count,omitempty" json:"count,omitempty"`
	//IconUrls struct {
	//	Medium string `json:"medium"`
	//} `json:"iconUrls"`
	//ID        int `json:"id"`
	MaxLevel  int `bson:"maxLevel" json:"maxLevel"`
	StarLevel int `bson:"starLevel,omitempty" json:"starLevel,omitempty"`
	Card
}

func (d *Deck) Equal(deck Deck) bool {
	n := 0
	for i := 0; i < len(*d)-1; i++ {
		for j := 0; j < len(deck); j++ {
			if (*d)[i].Name == deck[j].Name && (*d)[i].Level == deck[j].Level {
				n++
				break
			}
		}
	}
	return len(*d) == n
}

func (d *Deck) FixLevel() {
	for i, card := range *d {
		(*d)[i].Level = 14 - card.MaxLevel + card.Level
	}
}

type Player struct {
	Achievements []struct {
		CompletionInfo interface{} `json:"completionInfo,omitempty"  bson:"completion_info"`
		Info           string      `json:"info,omitempty"  bson:"info"`
		Name           string      `json:"name,omitempty"  bson:"name"`
		Stars          int         `json:"stars,omitempty"  bson:"stars"`
		Target         int         `json:"target,omitempty"  bson:"target"`
		Value          int         `json:"value,omitempty"  bson:"value"`
	} `json:"achievements,omitempty"  bson:"achievements"`
	Arena struct {
		ID   int    `json:"id,omitempty"  bson:"id"`
		Name string `json:"name,omitempty"  bson:"name"`
	} `json:"arena,omitempty"  bson:"arena"`
	Badges []struct {
		Name     string `json:"name,omitempty"  bson:"name"`
		Progress int    `json:"progress,omitempty"  bson:"progress"`
		Level    int    `json:"level,omitempty"  bson:"level"`
		MaxLevel int    `json:"maxLevel,omitempty"  bson:"max_level"`
		Target   int    `json:"target,omitempty"  bson:"target"`
	} `json:"badges,omitempty"  bson:"badges"`
	BattleCount       int  `json:"battleCount,omitempty"  bson:"battle_count"`
	BestTrophies      int  `json:"bestTrophies,omitempty"  bson:"best_trophies"`
	Cards             Deck `json:"cards,omitempty"  bson:"cards"`
	ChallengeCardsWon int  `json:"challengeCardsWon,omitempty"  bson:"challenge_cards_won"`
	ChallengeMaxWins  int  `json:"challengeMaxWins,omitempty"  bson:"challenge_max_wins"`
	Clan              struct {
		BadgeID int    `json:"badgeId,omitempty"  bson:"badge_id"`
		Name    string `json:"name,omitempty"  bson:"name"`
		Tag     string `json:"tag,omitempty"  bson:"tag"`
	} `json:"clan,omitempty"  bson:"clan"`
	ClanCardsCollected   int  `json:"clanCardsCollected,omitempty"  bson:"clan_cards_collected"`
	CurrentDeck          Deck `json:"currentDeck,omitempty"  bson:"current_deck"`
	CurrentFavouriteCard struct {
		//IconUrls struct {
		//	Medium string `json:"medium"`
		//} `json:"iconUrls"`
		ID       int    `json:"id,omitempty"  bson:"id"`
		MaxLevel int    `json:"maxLevel,omitempty"  bson:"max_level"`
		Name     string `json:"name,omitempty"  bson:"name"`
	} `json:"currentFavouriteCard,omitempty"  bson:"current_favourite_card"`
	Donations         int `json:"donations,omitempty"  bson:"donations"`
	DonationsReceived int `json:"donationsReceived,omitempty"  bson:"donations_received"`
	ExpLevel          int `json:"expLevel,omitempty"  bson:"exp_level"`
	ExpPoints         int `json:"expPoints,omitempty"  bson:"exp_points"`
	LeagueStatistics  struct {
		BestSeason struct {
			ID       string `json:"id,omitempty"  bson:"id"`
			Trophies int    `json:"trophies,omitempty"  bson:"trophies"`
		} `json:"bestSeason,omitempty"  bson:"best_season"`
		CurrentSeason struct {
			BestTrophies int `json:"bestTrophies,omitempty"  bson:"best_trophies"`
			Trophies     int `json:"trophies,omitempty"  bson:"trophies"`
		} `json:"currentSeason,omitempty"  bson:"current_season"`
		PreviousSeason struct {
			BestTrophies int    `json:"bestTrophies,omitempty"  bson:"best_trophies"`
			ID           string `json:"id,omitempty"  bson:"id"`
			Trophies     int    `json:"trophies,omitempty"  bson:"trophies"`
		} `json:"previousSeason,omitempty"  bson:"previous_season"`
	} `json:"leagueStatistics,omitempty"  bson:"league_statistics"`
	Losses                int    `json:"losses,omitempty"  bson:"losses"`
	Name                  string `json:"name,omitempty"  bson:"name"`
	Role                  string `json:"role,omitempty"  bson:"role"`
	StarPoints            int    `json:"starPoints,omitempty"  bson:"star_points"`
	Tag                   string `json:"tag,omitempty"  bson:"tag"`
	ThreeCrownWins        int    `json:"threeCrownWins,omitempty"  bson:"three_crown_wins"`
	TotalDonations        int    `json:"totalDonations,omitempty"  bson:"total_donations"`
	TournamentBattleCount int    `json:"tournamentBattleCount,omitempty"  bson:"tournament_battle_count"`
	TournamentCardsWon    int    `json:"tournamentCardsWon,omitempty"  bson:"tournament_cards_won"`
	Trophies              int    `json:"trophies,omitempty"  bson:"trophies"`
	WarDayWins            int    `json:"warDayWins,omitempty"  bson:"war_day_wins"`
	Wins                  int    `json:"wins,omitempty"  bson:"wins"`
}

func (p *Player) FixLevel() {
	p.CurrentDeck.FixLevel()
	p.Cards.FixLevel()
}

type Chests struct {
	Items []struct {
		Index int    `json:"index,omitempty"  bson:"index"`
		Name  string `json:"name,omitempty"  bson:"name"`
	} `json:"items,omitempty"  bson:"items"`
}

type Battlelog []struct {
	Type               string `json:"type,omitempty"  bson:"type"`
	BattleTime         string `json:"battleTime,omitempty"  bson:"battle_time"`
	IsLadderTournament bool   `json:"isLadderTournament,omitempty"  bson:"is_ladder_tournament"`
	Arena              struct {
		ID   int    `json:"id,omitempty"  bson:"id"`
		Name string `json:"name,omitempty"  bson:"name"`
	} `json:"arena,omitempty"  bson:"arena"`
	GameMode struct {
		ID   int    `json:"id,omitempty"  bson:"id"`
		Name string `json:"name,omitempty"  bson:"name"`
	} `json:"gameMode,omitempty"  bson:"game_mode"`
	DeckSelection string `json:"deckSelection,omitempty"  bson:"deck_selection"`
	Team          []struct {
		Tag                     string `json:"tag,omitempty"  bson:"tag"`
		Name                    string `json:"name,omitempty"  bson:"name"`
		StartingTrophies        int    `json:"startingTrophies,omitempty"  bson:"starting_trophies"`
		TrophyChange            int    `json:"trophyChange,omitempty"  bson:"trophy_change"`
		Crowns                  int    `json:"crowns,omitempty"  bson:"crowns"`
		KingTowerHitPoints      int    `json:"kingTowerHitPoints,omitempty"  bson:"king_tower_hit_points"`
		PrincessTowersHitPoints []int  `json:"princessTowersHitPoints,omitempty"  bson:"princess_towers_hit_points"`
		Clan                    struct {
			Tag     string `json:"tag,omitempty"  bson:"tag"`
			Name    string `json:"name,omitempty"  bson:"name"`
			BadgeID int    `json:"badgeId,omitempty"  bson:"badge_id"`
		} `json:"clan,omitempty"  bson:"clan"`
		Cards Deck `json:"cards,omitempty"  bson:"cards"`
	} `json:"team,omitempty"  bson:"team"`
	Opponent []struct {
		Tag                     string `json:"tag,omitempty"  bson:"tag"`
		Name                    string `json:"name,omitempty"  bson:"name"`
		StartingTrophies        int    `json:"startingTrophies,omitempty"  bson:"starting_trophies"`
		TrophyChange            int    `json:"trophyChange,omitempty"  bson:"trophy_change"`
		Crowns                  int    `json:"crowns,omitempty"  bson:"crowns"`
		KingTowerHitPoints      int    `json:"kingTowerHitPoints,omitempty"  bson:"king_tower_hit_points"`
		PrincessTowersHitPoints []int  `json:"princessTowersHitPoints,omitempty"  bson:"princess_towers_hit_points"`
		Clan                    struct {
			Tag     string `json:"tag,omitempty"  bson:"tag"`
			Name    string `json:"name,omitempty"  bson:"name"`
			BadgeID int    `json:"badgeId,omitempty"  bson:"badge_id"`
		} `json:"clan,omitempty"  bson:"clan"`
		Cards Deck `json:"cards,omitempty"  bson:"cards"`
	} `json:"opponent,omitempty"  bson:"opponent"`
	IsHostedMatch       bool   `json:"isHostedMatch,omitempty"  bson:"is_hosted_match"`
	BoatBattleSide      string `json:"boatBattleSide,omitempty"  bson:"boat_battle_side"`
	BoatBattleWon       bool   `json:"boatBattleWon,omitempty"  bson:"boat_battle_won"`
	NewTowersDestroyed  int    `json:"newTowersDestroyed,omitempty"  bson:"new_towers_destroyed"`
	PrevTowersDestroyed int    `json:"prevTowersDestroyed,omitempty"  bson:"prev_towers_destroyed"`
	RemainingTowers     int    `json:"remainingTowers,omitempty"  bson:"remaining_towers"`
}

func (b *Battlelog) GetBattles() []Battle {
	var battles []Battle
	for _, battle := range *b {
		if battle.Type == "PvP" {
			if len(battle.Opponent) != 1 || len(battle.Team) != 1 {
				panic("The following battle have more than 1 opponent! " + helper.MarshalIndent(battle))
			}
			var opponentDeck, myDeck []Card
			for _, card := range battle.Opponent[0].Cards {
				opponentDeck = append(opponentDeck, card.Card)
			}
			for _, card := range battle.Team[0].Cards {
				myDeck = append(myDeck, card.Card)
			}

			var win bool
			if battle.Team[0].TrophyChange > 0 {
				win = true
			} else {
				win = false
			}

			currentBattle := Battle{
				MyDeck:   myDeck,
				Opponent: opponentDeck,
				Win:      win}
			currentBattle.SortDecks()

			battles = append(battles, currentBattle)
		}
	}
	return battles
}

func (b *Battlelog) GetWin() []Battle {
	battles := b.GetBattles()
	var win []Battle
	for _, battle := range battles {
		if battle.Win {
			win = append(win, battle)
		}
	}
	return win
}
func (b *Battlelog) GetLose() []Battle {
	battles := b.GetBattles()
	var lose []Battle
	for _, battle := range battles {
		if !battle.Win {
			lose = append(lose, battle)
		}
	}
	return lose
}

func (b *Battlelog) FixLevel() {
	for i := range *b {
		for j := range (*b)[i].Opponent {
			(*b)[i].Opponent[j].Cards.FixLevel()
		}
		for j := range (*b)[i].Team {
			(*b)[i].Team[j].Cards.FixLevel()
		}
	}
}
