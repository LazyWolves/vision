package models

type alias struct {
	AliasName string `json:"alias_name"`
	AliasTo string `json:"alias_to"`
}

type ConfigModel struct {
	Port int64 `json:"port"`
	AllowAll bool `json:"allow_all"`
	BlockFor []string `json:"block_for"`
	AllowFor []string `json:"allow_for"`
	Aliases []alias `json:"aliases"`
}
