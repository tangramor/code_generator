package main

import (
    "fmt"
    "flag"
    "time"
    "strings"
    "math/rand"
    "github.com/nu7hatch/gouuid"
)

func main() {

    length := flag.Int("len", 18, "生成码长度")
    number := flag.Int("num", 100, "生成码数量")
    flag.Parse()

    chars := "23456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
    arr_chars := strings.Split(chars, "")

    rand.Seed(int64(time.Now().Nanosecond()))

    u4, err := uuid.NewV4()
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    prefix := strings.Replace(strings.ToUpper(strings.Split(u4.String(), "-")[0]), "0", arr_chars[rand.Intn(len(arr_chars))], -1)
    prefix = strings.Replace(prefix, "1", arr_chars[rand.Intn(len(arr_chars))], -1)
    prefix = strings.Replace(prefix, "O", arr_chars[rand.Intn(len(arr_chars))], -1)

    for n := 0; n < *number; n ++ {
        result := ""
        for i := 0; i < (*length - 8); i ++ {
            rand.Seed(int64(time.Now().Nanosecond()))
            s := []string{result, arr_chars[rand.Intn(len(arr_chars))]}
            result = strings.Join(s, "")
        }
        fmt.Printf("%s%s\n", prefix, result)
    }
   
}
