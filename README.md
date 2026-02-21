## API Key From Football-Data.org

`API_KEY=api-key-here`

## Simple Request

```
func main() {
	client := footballdataapi.NewClient()
	resp, err := client.TeamMatches.Fetch(759)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
```
