package kwic

import "fmt"
import "log"
import "net/http"
import "encoding/json"
import "time"

type DBLPRecord struct {
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

type DBLPStorageManager struct {
	lines []string
}

func (this *DBLPStorageManager) Init() {
	var query string

	fmt.Print("Enter the DBLP search criteria (such as the author name): ")
	fmt.Scan(&query)
	url := "http://dblp.org/search/publ/api?q=" + query + "&format=json"
	record := new(DBLPRecord)
	this.lines = makeRequest(url, record)
}

func (this *DBLPStorageManager) Line(index int) string {
	return this.lines[index]
}

func (this *DBLPStorageManager) Length() int {
	return len(this.lines)
}

func makeRequest(url string, record *DBLPRecord) []string {
	client := &http.Client{Timeout: 10 * time.Second}
	request, err := client.Get(url)
	var titles []string

	if err != nil {
		log.Fatal(err)
	}
	defer request.Body.Close()

	json.NewDecoder(request.Body).Decode(record)

	for _, hit := range record.Result.Hits.Hit {
		titles = append(titles, hit.Info.Title)
	}

	return titles
}
