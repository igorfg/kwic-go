package kwic

// import "fmt"
//     import "os"
// import "log"
// import "net/http"
// import "encoding/json"

// type DBLPStorageManager struct {
// 	lines []string
// }

// func (this *DBLPStorageManager) Init() {
// 	var query string

// 	fmt.Print("Enter the DBLP search criteria (such as the author name): ")
// 	fmt.Scan(&query)
// 	url := "http://dblp.org/search/publ/api?q=" + query + "&format=json"
// 	lines := makeRequest(url)
// }

// func makeRequest(url string) []string {
// 	client := &http.Client{Timeout: 10 * time.Second}
// 	request, err := client.Get(url)

// 	if err != nil {
// 		return err
// 	}
// 	defer request.Body.Close()

// 	result := json.NewDecoder(request.Body).Decode(target)
// }
