package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Member struct {
	ID   int
	Name string
}

type ApiResponse struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	sheet_id := os.Getenv("SHEET_ID")
	sheet_name := os.Getenv("SHEET_NAME")
	api_key := os.Getenv("API_KEY")

	fmt.Println(sheet_id)
	fmt.Println(api_key)

	members, err := getMembers(sheet_id, sheet_name, api_key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 現在の週を取得
	_, currentWeek := time.Now().ISOWeek()

	// メンバーの週替わりの掃除当番を出力
	printCleaningDuty(members, currentWeek)

}

func getMembers(sheet_id string, sheet_name string, api_key string) ([]Member, error) {
	url := fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%s/values/%s?key=%s", sheet_id, sheet_name, api_key)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	members := convertToMembers(apiResponse.Values)

	return members, nil
}

func convertToMembers(values [][]string) []Member {
	members := make([]Member, 0, len(values)-1)

	for i := 1; i < len(values); i++ {
		id, _ := strconv.Atoi(values[i][0])
		name := values[i][1]

		members = append(members, Member{
			ID:   id,
			Name: name,
		})
	}

	return members
}

func printCleaningDuty(members []Member, currentWeek int) {
	numMembers := len(members)
	fmt.Println("週替わりの掃除当番:")
	for i, member := range members {
		fmt.Printf("Week %d: %s\n", i+1, member.Name)
	}

	// 現在の週の当番
	fmt.Println("\nこの週の当番:")
	currentWeekDuty := members[(currentWeek-1)%numMembers]
	fmt.Printf("Week %d: %s\n", currentWeek, currentWeekDuty.Name)

	// 次の週の当番
	fmt.Println("\n次の週の当番:")
	nextWeekDuty := members[currentWeek%numMembers]
	fmt.Printf("Week %d: %s\n", currentWeek+1, nextWeekDuty.Name)
}
