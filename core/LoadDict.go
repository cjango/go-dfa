package core

import (
	"io/ioutil"
	"os"
	"strings"
)

type KeywordsMap struct {
	Node  map[string]interface{}
	isEnd bool
}

var s *KeywordsMap

/*
初始化敏感词词典结构体
*/
func initKeywordsMap() *KeywordsMap {
	return &KeywordsMap{
		Node:  make(map[string]interface{}),
		isEnd: false,
	}
}

func LoadKeywords(types []string) {
	for _, dict := range types {
		file, err := os.Open("dict/" + dict)
		if err != nil {
			_ = file.Close()
			panic(err)
		}
		str, err := ioutil.ReadAll(file)
		dictionary := strings.Fields(string(str))
		s = initKeywordsMap()
		for _, word := range dictionary {
			sMapTmp := s
			w := []rune(word)
			// 确定敏感词的长度
			wordsLength := len(w)
			// 遍历词中的字符
			for i := 0; i < wordsLength; i++ {
				// 找到第N个字符
				t := string(w[i])
				isEnd := false
				//如果是敏感词的最后一个字，则确定状态
				if i == (wordsLength - 1) {
					isEnd = true
				}
				func(tx string) {
					if _, ok := sMapTmp.Node[tx]; !ok {
						//如果该字在该层级索引中找不到，则创建新的层级
						sMapTemp := new(KeywordsMap)
						sMapTemp.Node = make(map[string]interface{})
						sMapTemp.isEnd = isEnd
						sMapTmp.Node[tx] = sMapTemp
					}
					sMapTmp = sMapTmp.Node[tx].(*KeywordsMap) //进入下一层级
					sMapTmp.isEnd = isEnd
				}(t)
			}
		}
	}
}
