package web

import "embed"

//go:embed *
var webAssets embed.FS

func GetEmbed() *embed.FS {
	return &webAssets
}
