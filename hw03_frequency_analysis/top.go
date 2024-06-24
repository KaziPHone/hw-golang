package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type KeyValue struct {
	key   string
	value int
}

func Top10(text string) []string {
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	return sortedMapCounterToArr(arrToMapCounter(strings.Split(text, " ")))
}

func arrToMapCounter(textArray []string) map[string]int {
	mText := make(map[string]int)
	for _, v := range textArray {
		if v != "" {
			mText[v]++
		}
	}
	return mText
}

func sortedMapCounterToArr(mText map[string]int) []string {
	kvArray := []KeyValue{}
	for key := range mText {
		kvArray = append(kvArray, KeyValue{
			value: mText[key],
			key:   key,
		})
	}
	sort.Slice(kvArray, func(i, j int) bool {
		return kvArray[i].value >= kvArray[j].value
	})
	resultArr := []string{}
	for i := 0; i < len(kvArray); i++ {
		resultArr = append(resultArr, kvArray[i].key)
		if i >= 9 {
			break
		}
	}
	return resultArr
}
