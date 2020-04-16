package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	kvM := make(map[[26]byte][]string)
	for _, s := range strs {
		k := [26]byte{}
		for _, b := range s {
			k[b-'a']++
		}
		if _, ok := kvM[k]; ok {
			kvM[k] = append(kvM[k], s)
		} else {
			kvM[k] = []string{s}
		}
	}
	result := make([][]string, 0, len(kvM))
	for _, value := range kvM {
		result = append(result, value)
	}
	return result
}
