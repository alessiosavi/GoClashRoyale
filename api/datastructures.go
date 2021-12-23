package api

type Player struct {
	Achievements []struct {
		CompletionInfo interface{} `json:"completionInfo"`
		Info           string      `json:"info"`
		Name           string      `json:"name"`
		Stars          int         `json:"stars"`
		Target         int         `json:"target"`
		Value          int         `json:"value"`
	} `json:"achievements"`
	Arena struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"arena"`
	Badges []struct {
		Name     string `json:"name"`
		Progress int    `json:"progress"`
		Level    int    `json:"level,omitempty"`
		MaxLevel int    `json:"maxLevel,omitempty"`
		Target   int    `json:"target,omitempty"`
	} `json:"badges"`
	BattleCount  int `json:"battleCount"`
	BestTrophies int `json:"bestTrophies"`
	Cards        []struct {
		Count    int `json:"count"`
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
		ID        int    `json:"id"`
		Level     int    `json:"level"`
		MaxLevel  int    `json:"maxLevel"`
		Name      string `json:"name"`
		StarLevel int    `json:"starLevel,omitempty"`
	} `json:"cards"`
	ChallengeCardsWon int `json:"challengeCardsWon"`
	ChallengeMaxWins  int `json:"challengeMaxWins"`
	Clan              struct {
		BadgeID int    `json:"badgeId"`
		Name    string `json:"name"`
		Tag     string `json:"tag"`
	} `json:"clan"`
	ClanCardsCollected int `json:"clanCardsCollected"`
	CurrentDeck        []struct {
		Count    int `json:"count"`
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
		ID        int    `json:"id"`
		Level     int    `json:"level"`
		MaxLevel  int    `json:"maxLevel"`
		Name      string `json:"name"`
		StarLevel int    `json:"starLevel,omitempty"`
	} `json:"currentDeck"`
	CurrentFavouriteCard struct {
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
		ID       int    `json:"id"`
		MaxLevel int    `json:"maxLevel"`
		Name     string `json:"name"`
	} `json:"currentFavouriteCard"`
	Donations         int `json:"donations"`
	DonationsReceived int `json:"donationsReceived"`
	ExpLevel          int `json:"expLevel"`
	ExpPoints         int `json:"expPoints"`
	LeagueStatistics  struct {
		BestSeason struct {
			ID       string `json:"id"`
			Trophies int    `json:"trophies"`
		} `json:"bestSeason"`
		CurrentSeason struct {
			BestTrophies int `json:"bestTrophies"`
			Trophies     int `json:"trophies"`
		} `json:"currentSeason"`
		PreviousSeason struct {
			BestTrophies int    `json:"bestTrophies"`
			ID           string `json:"id"`
			Trophies     int    `json:"trophies"`
		} `json:"previousSeason"`
	} `json:"leagueStatistics"`
	Losses                int    `json:"losses"`
	Name                  string `json:"name"`
	Role                  string `json:"role"`
	StarPoints            int    `json:"starPoints"`
	Tag                   string `json:"tag"`
	ThreeCrownWins        int    `json:"threeCrownWins"`
	TotalDonations        int    `json:"totalDonations"`
	TournamentBattleCount int    `json:"tournamentBattleCount"`
	TournamentCardsWon    int    `json:"tournamentCardsWon"`
	Trophies              int    `json:"trophies"`
	WarDayWins            int    `json:"warDayWins"`
	Wins                  int    `json:"wins"`
}

type Chests struct {
	Items []struct {
		Index int    `json:"index"`
		Name  string `json:"name"`
	} `json:"items"`
}

type Battlelog []struct {
	Type               string `json:"type"`
	BattleTime         string `json:"battleTime"`
	IsLadderTournament bool   `json:"isLadderTournament"`
	Arena              struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"arena"`
	GameMode struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"gameMode"`
	DeckSelection string `json:"deckSelection"`
	Team          []struct {
		Tag                string `json:"tag"`
		Name               string `json:"name"`
		StartingTrophies   int    `json:"startingTrophies,omitempty"`
		Crowns             int    `json:"crowns"`
		KingTowerHitPoints int    `json:"kingTowerHitPoints"`
		Clan               struct {
			Tag     string `json:"tag"`
			Name    string `json:"name"`
			BadgeID int    `json:"badgeId"`
		} `json:"clan"`
		Cards []struct {
			Name      string `json:"name"`
			ID        int    `json:"id"`
			Level     int    `json:"level"`
			StarLevel int    `json:"starLevel,omitempty"`
			MaxLevel  int    `json:"maxLevel"`
			IconUrls  struct {
				Medium string `json:"medium"`
			} `json:"iconUrls"`
		} `json:"cards"`
	} `json:"team"`
	Opponent []struct {
		Tag                     string `json:"tag"`
		Name                    string `json:"name"`
		StartingTrophies        int    `json:"startingTrophies,omitempty"`
		Crowns                  int    `json:"crowns"`
		KingTowerHitPoints      int    `json:"kingTowerHitPoints"`
		PrincessTowersHitPoints []int  `json:"princessTowersHitPoints"`
		Clan                    struct {
			Tag     string `json:"tag"`
			Name    string `json:"name"`
			BadgeID int    `json:"badgeId"`
		} `json:"clan"`
		Cards []struct {
			Name     string `json:"name"`
			ID       int    `json:"id"`
			Level    int    `json:"level"`
			MaxLevel int    `json:"maxLevel"`
			IconUrls struct {
				Medium string `json:"medium"`
			} `json:"iconUrls"`
			StarLevel int `json:"starLevel,omitempty"`
		} `json:"cards"`
	} `json:"opponent"`
	IsHostedMatch       bool   `json:"isHostedMatch"`
	BoatBattleSide      string `json:"boatBattleSide,omitempty"`
	BoatBattleWon       bool   `json:"boatBattleWon,omitempty"`
	NewTowersDestroyed  int    `json:"newTowersDestroyed,omitempty"`
	PrevTowersDestroyed int    `json:"prevTowersDestroyed,omitempty"`
	RemainingTowers     int    `json:"remainingTowers,omitempty"`
}
