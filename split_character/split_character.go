package split_character

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

const twitterMaxLength = 280

func SplitCharacter() {
	for {
		fmt.Println("===================================================")
		fmt.Println("Please enter the content of the tweet (finally enter \"exit\" to end the input):")
		var lines []string
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimRightFunc(input, unicode.IsSpace)
			if input == "exit" {
				break
			}
			lines = append(lines, input)
		}

		content := strings.Join(lines, "\n")
		tweetLen, exceedsLimit := calculateTwitterContentLen(content)
		fmt.Println("===================================================")
		fmt.Println("Current Tweet Length:", tweetLen)
		fmt.Println("Exceed the Twitter length limit:", exceedsLimit)
		if !exceedsLimit {
			continue
		}
		fmt.Println("===================================================")
		fmt.Println("Now back to the cut tweet content")
		// 对超过限制的内容进行分割
		outContents := outSplitStr(content)
		for i := 0; i < len(outContents); i++ {
			fmt.Printf(">>>>>>>「the %dth tweet content」\n", i+1)
			fmt.Println(outContents[i])
		}
		fmt.Println(">>>>>>>「End of tweet cutting」")
	}
}

var (
	prefixFormat = "(%d/%d)"
)

func buildPrefixStr(num, all int, content string) string {
	preStr := fmt.Sprintf(prefixFormat, num, all)
	return preStr + content
}

func prefixStrLen(all int) int {
	strLen := 0
	for i := 1; i <= all; i++ {
		strLen += len(fmt.Sprintf(prefixFormat, i, all))
	}
	return strLen
}

func calculateTwitterContentLen(content string) (int, bool) {
	runeCount := utf8.RuneCountInString(content)
	tweetLen := runeCount + (len(content)-runeCount)/2
	return tweetLen, tweetLen > twitterMaxLength
}

func calculateTwitterRuneContentLen(runeContent []rune) (int, bool) {
	tweetLen := len(runeContent) + (len(string(runeContent))-len(runeContent))/2
	if (len(string(runeContent))-len(runeContent))%2 != 0 {
		fmt.Println(111)
	}

	return tweetLen, tweetLen >= twitterMaxLength
}

func splitFirstMaxTwitterContent(content string) (left string, right string) {
	contentLen := len(content)
	if contentLen <= twitterMaxLength {
		return content, ""
	}

	runeContent := []rune(content)
	runeIdx := sort.Search(len(runeContent), func(i int) bool {
		_, isExceed := calculateTwitterRuneContentLen(runeContent[:i])
		return isExceed
	})
	return string(runeContent[:runeIdx]), string(runeContent[runeIdx:])
}

func outSplitStr(content string) []string {
	preLen, _ := calculateTwitterContentLen(content)
	preAll := (preLen + twitterMaxLength) / twitterMaxLength
	nowLen := preLen + prefixStrLen(preAll)
	nowAll := (nowLen + twitterMaxLength) / twitterMaxLength
	result := make([]string, 0, nowAll)
	for i := 1; i <= nowAll; i++ {
		content = buildPrefixStr(i, nowAll, content)
		left, right := splitFirstMaxTwitterContent(content)
		result = append(result, left)
		content = right
	}
	return result
}
