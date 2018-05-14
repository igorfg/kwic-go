package main

import "time"

//import "strings"
//import "github.com/igorfg/kwic-go/kwic"

import "net/http"
import "encoding/json"

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

type DBLP struct {
	Result struct {
		Query  string `json:"query"`
		Status struct {
			Code string `json:"@code"`
			Text string `json:"text"`
		} `json:"status"`
		Time struct {
			Unit string `json:"@unit"`
			Text string `json:"text"`
		} `json:"time"`
		Completions struct {
			Total    string `json:"@total"`
			Computed string `json:"@computed"`
			Sent     string `json:"@sent"`
			C        []struct {
				Sc   string `json:"@sc"`
				Dc   string `json:"@dc"`
				Oc   string `json:"@oc"`
				ID   string `json:"@id"`
				Text string `json:"text"`
			} `json:"c"`
		} `json:"completions"`
		Hits struct {
			Total    string `json:"@total"`
			Computed string `json:"@computed"`
			Sent     string `json:"@sent"`
			First    string `json:"@first"`
			Hit      []struct {
				Score string `json:"@score"`
				ID    string `json:"@id"`
				Info  struct {
					Authors struct {
						Author []string `json:"author"`
					} `json:"authors"`
					Title  string `json:"title"`
					Venue  string `json:"venue"`
					Volume string `json:"volume"`
					Number string `json:"number"`
					Pages  string `json:"pages"`
					Year   string `json:"year"`
					Type   string `json:"type"`
					Key    string `json:"key"`
					Doi    string `json:"doi"`
					Ee     string `json:"ee"`
					URL    string `json:"url"`
				} `json:"info"`
				URL string `json:"url"`
			} `json:"hit"`
		} `json:"hits"`
	} `json:"result"`
}

func main() {
	foo1 := new(DBLP) // or &Foo{}
	getJson("http://dblp.org/search/publ/api?q=golang&format=json", &foo1)
	for i := 0; i < len(foo1.Result.Hits.Hit); i++ {
		println(foo1.Result.Hits.Hit[i].Info.Title)
	}

}
