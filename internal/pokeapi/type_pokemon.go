package pokeapi

type Pokemon struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []Ability          `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []GameIndex        `json:"game_indices"`
	HeldItems              []HeldItem         `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []Move             `json:"moves"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                Sprites            `json:"sprites"`
	Cries                  Cries              `json:"cries"`
	Stats                  []Stat             `json:"stats"`
	Types                  []TypeElement      `json:"types"`
	PastTypes              []PastType         `json:"past_types"`
	PastAbilities          []PastAbility      `json:"past_abilities"`
}

type Ability struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type HeldItem struct {
	Item           NamedAPIResource `json:"item"`
	VersionDetails []VersionDetail  `json:"version_details"`
}

type VersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type Move struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type Sprites struct {
	BackDefault      *string                              `json:"back_default"`
	BackFemale       *string                              `json:"back_female"`
	BackShiny        *string                              `json:"back_shiny"`
	BackShinyFemale  *string                              `json:"back_shiny_female"`
	FrontDefault     *string                              `json:"front_default"`
	FrontFemale      *string                              `json:"front_female"`
	FrontShiny       *string                              `json:"front_shiny"`
	FrontShinyFemale *string                              `json:"front_shiny_female"`
	Other            OtherSprites                         `json:"other"`
	Versions         map[string]map[string]VersionSprites `json:"versions"`
}

type OtherSprites struct {
	DreamWorld      BasicSpriteSet `json:"dream_world"`
	Home            HomeSpriteSet  `json:"home"`
	OfficialArtwork ArtworkSet     `json:"official-artwork"`
	Showdown        ShowdownSet    `json:"showdown"`
}

type BasicSpriteSet struct {
	FrontDefault *string `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type HomeSpriteSet struct {
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type ArtworkSet struct {
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type ShowdownSet struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type VersionSprites struct {
	BackDefault      *string `json:"back_default"`
	BackGray         *string `json:"back_gray"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontGray        *string `json:"front_gray"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type TypeElement struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PastType struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []TypeElement    `json:"types"`
}

type PastAbility struct {
	Generation NamedAPIResource `json:"generation"`
	Abilities  []AbilitySimple  `json:"abilities"`
}

type AbilitySimple struct {
	Ability  *NamedAPIResource `json:"ability"`
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
