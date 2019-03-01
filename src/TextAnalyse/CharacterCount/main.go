package main

import (
	"fmt"
	"os"
	"time"
	"sort"
	"TextAnalyse/CharacterCount"
)

type Pair struct{
	key rune
	value int
}

type PairList []Pair

func main(){

	parameters := os.Args[1:]
	if len(parameters) == 0 {
		fmt.Println("请输入文件名")
	} else if len(parameters) > 1 {
		fmt.Println("只接受一个文件")
	} else {
		start := time.Now()
		//统计字符
		dictionary_One,totalCount_One := CharacterCount.CountByLine(parameters[0])

		CountByLine := time.Now()

		dictionary_two,totalCount_two := CharacterCount.CountByAll(parameters[0])
		
		CountByAll := time.Now()

		fmt.Printf("totalCount:%d\t[CountByLine]Cost time %v\n",totalCount_One,CountByLine.Sub(start))
		sortMapByValueDesc(dictionary_One)

		fmt.Printf("totalCount:%d\t[CountByAll]Cost time %v\n",totalCount_two,CountByAll.Sub(CountByLine))
		sortMapByValueDesc(dictionary_two)
	}
}

func sortMapByValueDesc(dictionary map[rune]int) {
	p := make(PairList,len(dictionary))
	i := 0
	for k,v := range(dictionary){
		p[i] = Pair{k,v}
		i++
	}
	sort.SliceStable(p, func(i, j int) bool{ return p[i].value > p[j].value })

	for i,t := range(p[:10]){
		fmt.Printf("%d\t%c(%#[2]x):%d\n",i,t.key,t.value)
	}
}
