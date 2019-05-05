package models

type alias struct {
	AliasName string
	AliasTo string
}

type ConfigModel struct {
	Port string `json:"port"`
	AllowAll bool `json:"allow_all"`
	BlockFor []string `json:"block_for"`
	AllowFor []string `json:"allow_for"`
	Aliases []alias `json:"aliases"`
}
