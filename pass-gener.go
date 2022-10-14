package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	help = flag.Bool("h", false, "need help info.")
	in, length = "", ""
)

func init(){
	flag.Parse()
	in = flag.Arg(0)
	length = flag.Arg(1)
	fmt.Printf("input parameters: input password is %s, lenth is %s \n", in, length)
}

func main() {
	if *help {
		fmt.Println("usage: encrypt_pass [mypassword] [length], eg: encrypt_pass abc123456 8")
		return
	}

	md5str := fmt.Sprintf("%x", md5.Sum([]byte(in)))
	fmt.Printf("\nMD5 sum into hex number:\n%s\n", md5str)

	if length != "" {
		truncate, err := strconv.Atoi(length)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("\nlast passpord after truncation:\n%s\n", md5str[:truncate])

		fmt.Printf("\nstrip alphabet and add prefix & postfix left:\n")
		count := 0
		truncate = truncate - 3
		fmt.Printf("%s%s", strings.ToUpper(string(in[0])), strings.ToLower(string(in[1])))
		for i := 0;  i < len(md5str); i++ {
			d := string(md5str[i])
			if "0" <= d && d <= "9" {
				fmt.Print(d)
				count++
			}
			if count >= truncate {
				fmt.Println(".")
				return
			}
		}
	}
}
