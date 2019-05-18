// Package containing struct definitions for holding configs and params used
// in vision
package models

// Struct for holding aliases - alias name and actual path
type alias struct {
	AliasName string `json:"alias_name"`
	AliasTo string `json:"alias_to"`
}

// Struct for holding the config json declared in config file
type ConfigModel struct {
	Port int64 `json:"port"`
	AllowAll bool `json:"allow_all"`
	BlockFor []string `json:"block_for"`
	AllowFor []string `json:"allow_for"`
	Aliases []alias `json:"aliases"`
}
