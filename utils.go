package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func getHeroImageUrl(heroName string) string {
	return fmt.Sprintf(baseImageUrl, heroName)
}

func fixName(name string) string {
	var result string = name
	result = strings.TrimPrefix(result, "npc_dota_hero_")
	result = strings.ReplaceAll(result, "_", " ")
	return titleCase(result)
}

func getItemsAsString(items Items) string {
	var itemsNamesList []string
	for _, item := range items {
		if item.Name != "empty" && item.Name != "item_tpscroll" {
			itemsNamesList = append(itemsNamesList, strings.ReplaceAll(fixName(item.Name), "Item ", ""))
		}
	}
	return strings.Join(itemsNamesList, ", ")
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}

func HasField(v interface{}, fieldName string) bool {
	val := reflect.ValueOf(v)
	typ := val.Type()
	if typ.Kind() != reflect.Struct {
		return false
	}
	field := val.FieldByName(fieldName)
	return field.IsValid()
}
