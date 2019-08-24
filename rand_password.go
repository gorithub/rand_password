/**
* rand_password
* ランダムなパスワードを生成する
* 引数にパスワードの文字数
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
    "flag"
    "strconv"
    "os"
)

const MAX_LENGTH = 62

func main(){

    flag.Parse()
    // 取得する文字数
    passLength, err := strconv.Atoi(flag.Args()[0])
    if err != nil || passLength > MAX_LENGTH {
        os.Exit(0)
    }

    // ランダム文字列を取得
    var passLang string
    passLang = getRand(passLength)

    fmt.Println(passLang)
}

/**
* ランダム文字列を取得
*/
func getRand(passLength int) string{
    // スライス
    var langList []string
    // 半角大小英数 62文字
    langList = append(getLangList(48, 57),getLangList(65, 90)...)
    langList = append(langList, getLangList(97, 122)...)

    var randNum int;
    var randLang string;
    // 乱数の範囲(0〜61)
    var candidateNum int = MAX_LENGTH; // 62指定で61までかな？
    for i := 0; i < passLength; i++{
        // 乱数の範囲を減らす
        randRange := candidateNum - i;
        //fmt.Println(randRange)
        rand.Seed(time.Now().UnixNano())
        // 範囲から乱数を取得
        randNum = rand.Intn(randRange)
        // 文字を取得
        lang := langList[randNum]
        // 一度取得した文字を対称から外す
        langList = append(langList[:randNum], langList[randNum+1:]...)
        
        randLang = randLang + lang
    }

    return randLang

}

/**
* 指定した範囲の英数字を取得
*/
func getLangList(asciiStart int, asciiEnd int) []string{
    var passLang []string
    for i := asciiStart; i <= asciiEnd; i++ {
        passLang = append(passLang, string(i))
    }

    return passLang
}
