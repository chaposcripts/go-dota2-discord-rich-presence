package main

type Provider struct {
	Name      string `json:"name"`
	AppID     int    `json:"appid"`
	Version   int    `json:"version"`
	Timestamp int    `json:"timestamp"`
}

type Map struct {
	Name                 string `json:"name"`
	MatchID              string `json:"matchid"`
	GameTime             int    `json:"game_time"`
	ClockTime            int    `json:"clock_time"`
	Daytime              bool   `json:"daytime"`
	NightstalkerNight    bool   `json:"nightstalker_night"`
	RadiantScore         int    `json:"radiant_score"`
	DireScore            int    `json:"dire_score"`
	GameState            string `json:"game_state"`
	Paused               bool   `json:"paused"`
	WinTeam              string `json:"win_team"`
	CustomGameName       string `json:"customgamename"`
	WardPurchaseCooldown int    `json:"ward_purchase_cooldown"`
}

type Player struct {
	SteamID            string         `json:"steamid"`
	AccountID          string         `json:"accountid"`
	Name               string         `json:"name"`
	Activity           string         `json:"activity"`
	TeamName           string         `json:"team_name"`
	Kills              int            `json:"kills"`
	Deaths             int            `json:"deaths"`
	Assists            int            `json:"assists"`
	LastHits           int            `json:"last_hits"`
	Denies             int            `json:"denies"`
	KillStreak         int            `json:"kill_streak"`
	CommandsIssued     int            `json:"commands_issued"`
	PlayerSlot         int            `json:"player_slot"`
	TeamSlot           int            `json:"team_slot"`
	Gold               int            `json:"gold"`
	GoldReliable       int            `json:"gold_reliable"`
	GoldUnreliable     int            `json:"gold_unreliable"`
	GoldFromHeroKills  int            `json:"gold_from_hero_kills"`
	GoldFromCreepKills int            `json:"gold_from_creep_kills"`
	GoldFromIncome     int            `json:"gold_from_income"`
	GoldFromShared     int            `json:"gold_from_shared"`
	Gpm                int            `json:"gpm"`
	Xpm                int            `json:"xpm"`
	KillList           map[string]int `json:"kill_list"`
}

type Hero struct {
	Name            string `json:"name"`
	Facet           int    `json:"facet"`
	XPos            int    `json:"xpos"`
	YPos            int    `json:"ypos"`
	ID              int    `json:"id"`
	Level           int    `json:"level"`
	XX              int    `json:"xp"`
	RespawnSeconds  int    `json:"respawn_seconds"`
	BuybackCost     int    `json:"buyback_cost"`
	BuybackCooldown int    `json:"buyback_cooldown"`
	Health          int    `json:"health"`
	MaxHealth       int    `json:"max_health"`
	HealthPercent   int    `json:"health_percent"`
	Mana            int    `json:"mana"`
	MaxMana         int    `json:"max_mana"`
	ManaPercent     int    `json:"mana_percent"`
	AttributesLevel int    `json:"attributes_level"`
	Silenced        bool   `json:"silenced"`
	Stunned         bool   `json:"stunned"`
	Disarmed        bool   `json:"disarmed"`
	MagicImmune     bool   `json:"magicimmune"`
	Alive           bool   `json:"alive"`
	Hexed           bool   `json:"hexed"`
	Muted           bool   `json:"muted"`
	Break           bool   `json:"break"`
	AghanimsScepter bool   `json:"aghanims_scepter"`
	AghanimsShard   bool   `json:"aghanims_shard"`
	Smoked          bool   `json:"smoked"`
	HasDebuff       bool   `json:"has_debuff"`
	Talent1         bool   `json:"talent_1"`
	Talent2         bool   `json:"talent_2"`
	Talent3         bool   `json:"talent_3"`
	Talent4         bool   `json:"talent_4"`
	Talent5         bool   `json:"talent_5"`
	Talent6         bool   `json:"talent_6"`
	Talent7         bool   `json:"talent_7"`
	Talent8         bool   `json:"talent_8"`
}

type Ability struct {
	Name          string `json:"name"`
	Level         int    `json:"level"`
	CanCast       bool   `json:"can_cast"`
	Passive       bool   `json:"passive"`
	AbilityActive bool   `json:"ability_active"`
	Cooldown      int    `json:"cooldown"`
	Ultimate      bool   `json:"ultimate"`
}

type Abilities map[string]Ability

type Item struct {
	Name string `json:"name"`
	// Purchaser int `json:"purchaser"`
	// ItemLevel int `json:"item_level"`
	// CanCast bool `json:"can_cast"`
	// Cooldown int `json:"cooldown"`
	// Passive bool `json:"passive"`
}
type Items map[string]Item

type DotaGsiRequest struct {
	Provider  Provider  `json:"provider"`
	Map       Map       `json:"map,omitempty"`
	Player    Player    `json:"player,omitempty"`
	Hero      Hero      `json:"hero,omitempty"`
	Abilities Abilities `json:"abilities,omitempty"`
	Items     Items     `json:"items,omitempty"`
}

type DotaGameState string

const (
	DotaGameStateDisconnected         DotaGameState = "DOTA_GAMERULES_STATE_DISCONNECTED"
	DotaGameStateGameInProgress       DotaGameState = "DOTA_GAMERULES_STATE_GAME_IN_PROGRESS"
	DotaGameStateStrategyTime         DotaGameState = "DOTA_GAMERULES_STATE_STRATEGY_TIME"
	DotaGameStateWaitForPlayersToLoad DotaGameState = "DOTA_GAMERULES_STATE_WAIT_FOR_PLAYERS_TO_LOAD"
	DotaGameStateInit                 DotaGameState = "DOTA_GAMERULES_STATE_INIT"
	DotaGameStateWaitForMapToLoad     DotaGameState = "DOTA_GAMERULES_STATE_WAIT_FOR_MAP_TO_LOAD"
	DotaGameStateHeroSelection        DotaGameState = "DOTA_GAMERULES_STATE_HERO_SELECTION"
	DotaGameStatePostGame             DotaGameState = "DOTA_GAMERULES_STATE_POST_GAME"
	DotaGameStatePreGame              DotaGameState = "DOTA_GAMERULES_STATE_PRE_GAME"
	DotaGameStateTeamShowcase         DotaGameState = "DOTA_GAMERULES_STATE_TEAM_SHOWCASE"
)

var DotaGameStateLabel map[DotaGameState]string = map[DotaGameState]string{
	DotaGameStateDisconnected:         "In menu",
	DotaGameStateGameInProgress:       "In match",
	DotaGameStateStrategyTime:         "Strategy discussion",
	DotaGameStateWaitForPlayersToLoad: "Waiting for players",
	DotaGameStateInit:                 "Initializing",
	DotaGameStateWaitForMapToLoad:     "Loading to map",
	DotaGameStateHeroSelection:        "Hero selection",
	DotaGameStatePostGame:             "Ending match",
	DotaGameStatePreGame:              "Pre-Game",
}

var DotaGameStateImage map[DotaGameState]string = map[DotaGameState]string{
	DotaGameStateDisconnected:         ImageDotaLogo,
	DotaGameStateGameInProgress:       "",
	DotaGameStateStrategyTime:         ImageDotaMap,
	DotaGameStateWaitForPlayersToLoad: ImageDotaLoading,
	DotaGameStateInit:                 ImageDotaLoading,
	DotaGameStateWaitForMapToLoad:     ImageDotaLoading,
	DotaGameStateHeroSelection:        ImageDotaLoading,
	DotaGameStatePostGame:             ImageDotaLoading,
}
