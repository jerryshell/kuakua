package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 随机数种子
	rand.Seed(time.Now().Unix())
	// 用户选择
	var opt int
	for {
		fmt.Print("\n1)夸智慧 2)夸仪态 3)夸口才 4)夸品质\n>>> ")
		_, err := fmt.Scan(&opt)
		if err != nil {
			log.Fatalln("I am not playing with you.")
		}

		// 读取惊叹词列表
		shockSlice := readWords("data/惊叹.txt")
		// 随机选择惊叹词
		shock := shockSlice[rand.Intn(len(shockSlice))][0]

		// 读取来源词列表
		sourceSlice := readWords("data/来源.txt")
		// 随机选择来源词
		source := sourceSlice[rand.Intn(len(sourceSlice))][0]

		// 读取事业祝福词列表
		careerSlice := readWords("data/事业祝福语.txt")
		// 随机选择事业祝福词
		// 两个事业祝福词索引
		careerIndex1 := rand.Intn(len(careerSlice))
		careerIndex2 := rand.Intn(len(careerSlice))
		// 需要保证 [事业祝福词索引 1] 和 [事业祝福词索引 2] 不相同
		for careerIndex1 == careerIndex2 {
			careerIndex2 = rand.Intn(len(careerSlice))
		}
		// 两个事业祝福语
		career1 := careerSlice[careerIndex1][0]
		career2 := careerSlice[careerIndex2][0]

		// 声明赞美词列表
		var praiseSlice [][]string
		// 根据用户选择读取相应的赞美词列表
		switch opt {
		case 1:
			praiseSlice = readWords("data/智慧.txt")
			break
		case 2:
			praiseSlice = readWords("data/仪态.txt")
			break
		case 3:
			praiseSlice = readWords("data/口才.txt")
			break
		default:
			praiseSlice = readWords("data/品质.txt")
			break
		}
		// 随机选择赞美词
		// 三个赞美词索引
		praiseIndex1 := rand.Intn(len(praiseSlice))
		praiseIndex2 := rand.Intn(len(praiseSlice))
		praiseIndex3 := rand.Intn(len(praiseSlice))
		// 需要保证三个赞美词不相同
		if praiseIndex1 == praiseIndex2 || praiseIndex1 == praiseIndex3 || praiseIndex2 == praiseIndex3 {
			praiseIndex2 = rand.Intn(len(praiseSlice))
			praiseIndex3 = rand.Intn(len(praiseSlice))
		}
		// 三个赞美词
		praise1 := praiseSlice[praiseIndex1][0]
		praise2 := praiseSlice[praiseIndex2][0]
		praise3 := praiseSlice[praiseIndex3][0]

		// 拼接和输出结果
		result := shock + "，从" + source + "可以看得出你是个" + praise1 + "、" + praise2 + "、" + praise3 + "的人，日后的道路一定会" + career1 + "、" + career2 + "！"
		fmt.Println(result)
	}
}

// 根据文件名读取相应的词语文件
func readWords(filename string) (records [][]string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	reader := csv.NewReader(f)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	_ = f.Close()
	return
}
