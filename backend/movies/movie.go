package movies

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
)

type Movie struct {
	Searches []Search `json:"Search"`
	Response string   `json:"Response"`
}
type Search struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", "http://www.omdbapi.com/?i=tt3896198&apikey=564dc484&s=major", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("cannot obtain movies")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result Movie
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("unmarshal bokka %v", err)
	}
	path := path.Join("../../templates", "index.html")
	exec, err := template.ParseFiles(path)
	if err != nil {
		fmt.Print(err)
	}
	results := result.Searches
	if err = exec.Execute(w, results); err != nil {
		fmt.Print(err)
	}

	//fmt.Println(res)
	//fmt.Println(string(body))
}
