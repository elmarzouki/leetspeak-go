package leetspeak

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func getAlphabets() map[string][]string {

	baseDir, _ := os.Getwd()
	//var alphabetics *Alphabet
	alphabetics := make(map[string][]string)
	// Open our jsonFile
	jsonFile, err := os.Open(baseDir + "/alphabet.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	err = json.Unmarshal(byteValue, &alphabetics)

	return alphabetics
}

func getRndLeet(substitutions []string) string {
	rnd := rand.Intn(len(substitutions))
	rndLeet := substitutions[rnd]
	return rndLeet
}

func ConvertToLeet(text string) string {

	alphabetics := getAlphabets()
	var leet string = ""

	for _, char := range strings.ToLower(text) {
		substitutions := alphabetics[string(char)]
		rndLeet := getRndLeet(substitutions)
		leet += rndLeet
	}

	return leet

}
