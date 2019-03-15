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
	rand.Seed(time.Now().Unix())
	var opt int
	for {
		fmt.Print("\n1)夸智慧 2)夸仪态 3)夸口才 4)夸品质\n>>> ")
		_, err := fmt.Scan(&opt)
		if err != nil {
			log.Fatalln("I am not playing with you.")
		}

		shock := readWords("data/惊叹.txt")
		source := readWords("data/来源.txt")
		goodGoodStudyDayDayUp := readWords("data/事业祝福语.txt")
		var praise [][]string
		switch opt {
		case 1:
			praise = readWords("data/智慧.txt")
			break
		case 2:
			praise = readWords("data/仪态.txt")
			break
		case 3:
			praise = readWords("data/口才.txt")
			break
		case 4:
			praise = readWords("data/品质.txt")
			break
		}

		fmt.Println(
			shock[rand.Intn(len(shock))][0] +
				"，从" +
				source[rand.Intn(len(source))][0] +
				"可以看得出你是个" +
				praise[rand.Intn(len(praise))][0] +
				"、" +
				praise[rand.Intn(len(praise))][0] +
				"、" +
				praise[rand.Intn(len(praise))][0] +
				"的人，日后的道路一定会" +
				goodGoodStudyDayDayUp[rand.Intn(len(goodGoodStudyDayDayUp))][0] +
				"、" +
				goodGoodStudyDayDayUp[rand.Intn(len(goodGoodStudyDayDayUp))][0] +
				"！")
	}
}

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
