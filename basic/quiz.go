package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	limit := flag.Int("limit", 5, "クイズの制限時間です。")
	csvFileName := flag.String("csv", "problems.csv", "問題が書かれたcsvファイル名を指定してください。")

	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("CSVファイルを開くのに失敗しました: %s", *csvFileName)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("CSVファイルの読み取りに失敗しました。")
	}

	problems := parseLines(lines)
	fmt.Println(problems)
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	<-timer.C //チャネルから値が返ってくるまでここで止まる

	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("\n%d問中%d問正解です。 \n", len(problems), correct)
			return
		default:
			fmt.Printf("第%d問 %s = ", i+1, p.question)
			var answer string
			fmt.Scanf("%s\n", &answer) //Scanfはスペースを取り除く
			if p.answer == answer {
				fmt.Println("正解です。")
				correct++
			}
			if p.answer != answer {
				fmt.Println("不正解です。")
				fmt.Printf("%s\n", answer)
			}
		}

		fmt.Printf("this is it. %s\n", answer)

	}
	fmt.Printf("\n%d問中%d問正解です。 \n", len(problems), correct)

}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), //Scanfはスペースを取り除くため、csvファイルで設定されている答えからもスペースを取り除いておく
		}
	}
	return problems
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Printf("%s\n", msg)
	os.Exit(1)
}
