package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func check(token string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/billing/country-code", nil)
	if err != nil {
		fmt.Println("[!] Error Making Request")
		return
	}
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[!] Error Sending Request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("[+] Valid Token")
		data_valid, e1 := os.OpenFile("Data/valid.txt", os.O_APPEND|os.O_WRONLY, 0666)
		if e1 != nil {
			return
		}
		defer data_valid.Close()

		_, _ = data_valid.WriteString("[+]\n")

		valid, e := os.OpenFile("valid.txt", os.O_APPEND|os.O_WRONLY, 0666)
		if e != nil {
			return
		}
		defer valid.Close()

		_, _ = valid.WriteString(token + "\n")
	} else if resp.StatusCode == 429 {
		fmt.Println("[!] Rate Limited")
	} else {
		b, e := os.OpenFile("Data/invalid.txt", os.O_APPEND|os.O_WRONLY, 0666)
		if e != nil {
			return
		}
		defer b.Close()

		fmt.Println("[-] Invalid Token")
		_, _ = b.WriteString("[-]\n")
	}
}

func main() {
	f2, err := os.OpenFile("Data/invalid.txt", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	f3, err2 := os.OpenFile("Data/valid.txt", os.O_RDWR|os.O_TRUNC, 0666)
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer f3.Close()

	var nigger string

	t, e := os.Open("tokens.txt")
	if e != nil {
		return
	}
	tokens_ := 0
	s := bufio.NewScanner(t)
	for s.Scan() {
		tokens_++
	}

	fmt.Println("Tokens:", tokens_, "| Press Any Key Then Enter To Start > ")
	fmt.Scan(&nigger)

	l, e2 := os.Open("tokens.txt")
	if e2 != nil {
		return
	}

	s_ := bufio.NewScanner(l)
	for s_.Scan() {
		tk := s_.Text()
		check(tk)
	}

	var exit string
	valid := 0
	invalid := 0

	val, lav := os.Open("Data/valid.txt")
	if lav != nil {
		return
	}
	alv := bufio.NewScanner(val)
	for alv.Scan() {
		valid++
	}

	val2, lav2 := os.Open("Data/valid.txt")
	if lav2 != nil {
		return
	}
	alv2 := bufio.NewScanner(val2)
	for alv2.Scan() {
		invalid++
	}

	fmt.Println("\nResults | Valid:", valid, "| Invalid:", invalid, "| Press Any Key Then Enter To Close >")
	fmt.Scan(&exit)

}
