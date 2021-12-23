package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Stats struct {
	Data Data `json:"data"`
}

type Data struct {
	PlatformInfo PlatformInfoDetail `json:"platformInfo"`
	Segments     []SegmentsDetail   `json:"segments"`
}

type PlatformInfoDetail struct {
	Platform_name string `json:"platformSlug"`
	User_id       string `json:"platformUserId"`
}

type SegmentsDetail struct {
	Stats StatsDetail `json:"stats"`
}

type StatsDetail struct {
	Level       Value `json:"level"`
	Rank_score  Value `json:"rankScore"`
	Arena_Score Value `json:"arenaRankScore"`
}

type Value struct {
	Val float32 `json:"value"`
}

func GetApexStats(api_endpoint string, api_key string, platform string, uid string) (*Stats, error) {
	request, err := http.NewRequest("GET", api_endpoint+"/standard/profile/"+platform+"/"+uid, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Set GET params
	params := request.URL.Query()
	params.Add("TRN-Api-Key", api_key)
	request.URL.RawQuery = params.Encode()

	// Set timeouts to 5s
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	// Send request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer response.Body.Close()

	// Read body data
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Get specific stats
	userData := Stats{}
	jsonErr := json.Unmarshal(body, &userData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return nil, err
	}
	userData.Data.Segments = userData.Data.Segments[:1]

	fmt.Printf("%+v\n", userData.Data)
	return &userData, nil
}
