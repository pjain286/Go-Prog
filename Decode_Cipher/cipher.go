package main

import(
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"flag"
)

func main() {

	// Read the ciphers
	fileName := flag.String("file","cipher.txt","Input file to read Ciphers")
	flag.Parse()
	ciphers := readFile(*fileName)

	// Find the Key
	keys := getkey(ciphers)
	
	//Finally decode the cipher using the key
	result := decipher(ciphers[len(ciphers)-1],keys)
	fmt.Println(result)
	
}

/**
Read Ciphers from file
Input - Filename.txt
Output - Ciphers []string
*/
func readFile(s string) []string{
	content,err := ioutil.ReadFile(s)
	if err!= nil{
		fmt.Println("Please check the input file (if file exists) and permissions")
		panic(err)
	}
	ciphers := strings.SplitAfter(string(content),"\n")
	return ciphers
}

/**
Get the final key from all the ciphers
Input - Ciphers []string
Output - keys []int

Logic - Convert the hex for each character to int
and decode the key for character by character
*/
func getkey(ciphers []string) []int{
	N := len(ciphers[len(ciphers)-1])
	var keys []int
	for i := 0; i < N; i+=2 {
		var sample []int
		for j := 0; j < len(ciphers); j++ {
			sample = append(sample,formSample(ciphers[j][i:i+2]))
		}
		key := findKey(sample)
		keys = append(keys,key)
	}
	return keys
}

/**
Decipher the cipher string using the key
Input - string,[]int
Output - string

Logic - Convert the cipher text to int char by char
and XOR with the key to get the original char
*/
func decipher(c string,k []int) string{
	var cipher,decipher []int
	for i := 0; i < len(c); i+=2 {
		cipher = append(cipher,formSample(c[i:i+2]))
	}
	for i := 0; i < len(k); i++ {
		decipher = append(decipher,cipher[i]^k[i])
	}

	var s string
	for i := 0; i < len(decipher); i++ {
		s+=string(decipher[i])
	}
	return s

}

/**
Parse the string hex to form an int using "strconv"
Input - string
Output - int
*/
func formSample(s string) int{
	res,err := strconv.ParseInt(s,16,32)
	if err!=nil{
		panic(err)
	}
	return int(res)
}


/**
Check if x is an alphabet a-z or A-Z
Input - x int
Output - bool
*/
func checkValid(x int) bool{
	
	if x>=65&&x<=90{
		return true
	}
	if x>=97&&x<=122{
		return true
	}
	return false
}

/**
Calculate the key for specific position in the string
Input - s []int
Output - key int

Logic - Use the property that a^space = A and vice-versa
on all the characters in the slice and decode the key
*/
func findKey(s []int) int{
	n := len(s)
	m := make(map[int]int)
	var ans,temp int
	ans = 0
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if true==checkValid(s[i]^s[j]){
				m[s[i]]++
				m[s[j]]++
			}
		}
	}
	for a,b := range m{
		if b>temp {
			temp = b
			ans = a
		}
	}
	if ans != 0{
		return ans^32 	
	}
	return 0
}
