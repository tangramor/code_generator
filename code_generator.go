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

    length := flag.Int("len", 18, "The length of code")
    number := flag.Int("num", 100, "The total number of codes")
    usepre := flag.Bool("prefix", true, "If use prefix with 8 charaters")
    flag.Parse()

    chars := "23456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
    arr_chars := strings.Split(chars, "")

    rand.Seed(int64(time.Now().Nanosecond()))

    leng := *length

    prefix := ""

    if *usepre {
        u4, err := uuid.NewV4()
        if err != nil {
            fmt.Println("error:", err)
            return
        }
        prefix = strings.Replace(strings.ToUpper(strings.Split(u4.String(), "-")[0]), "0", arr_chars[rand.Intn(len(arr_chars))], -1)
        prefix = strings.Replace(prefix, "1", arr_chars[rand.Intn(len(arr_chars))], -1)
        prefix = strings.Replace(prefix, "O", arr_chars[rand.Intn(len(arr_chars))], -1)

        leng = leng - 8
    }

    for n := 0; n < *number; n ++ {
        result := ""
        for i := 0; i < leng; i ++ {
            rand.Seed(int64(time.Now().Nanosecond()))
            s := []string{result, arr_chars[rand.Intn(len(arr_chars))]}
            result = strings.Join(s, "")
        }
        fmt.Printf("%s%s\n", prefix, result)
    }
   
}
