package main

import (
	"fmt"
	"strings"
	"unicode"
)

func getHeroImageUrl(heroName string) string {
	heroImageUrl := fmt.Sprintf(baseImageUrl, strings.ReplaceAll(strings.TrimPrefix(heroName, "npc_dota_hero_"), "_", "-"))
	fmt.Println("IMAGE URL:", heroName, heroImageUrl)
	return heroImageUrl
}

func fixName(name string) string {
	var result string = name
	result = strings.TrimPrefix(result, "npc_dota_hero_")
	result = strings.ReplaceAll(result, "_", " ")
	return titleCase(result)
}

func getItemsAsString() string {
	var items []string
	for _, item := range activity.Items {
		if item.Name != "empty" && item.Name != "item_tpscroll" {
			items = append(items, fixName(item.Name))
		}
	}
	return strings.Join(items, ", ")
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}
