package koboldai

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	http "github.com/bogdanfinn/fhttp"

	"github.com/aandrew-me/tgpt/v2/client"
	"github.com/aandrew-me/tgpt/v2/structs"
)

type Response struct {
	Token      string `json:"token"`
}

func NewRequest(input string, params structs.Params, prevMessages string) (*http.Response, error) {
	client, err := client.NewClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	safeInput, _ := json.Marshal(input)

	temperature := "0.5"
	if params.Temperature != ""{
		temperature = params.Temperature
	}

	top_p := "0.5"
	if params.Top_p != ""{
		top_p = params.Top_p
	}

	var data = strings.NewReader(fmt.Sprintf(`{
		"prompt": %v,
		"temperature": %v,
		"top_p": %v,
		"max_length": 300
	  }
	`, string(safeInput), temperature, top_p))

	req, err := http.NewRequest("POST", "https://koboldai-koboldcpp-tiefighter.hf.space/api/extra/generate/stream", data)
	if err != nil {
		fmt.Println("\nSome error has occurred.")
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// Setting all the required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// Return response
	return (client.Do(req))
}

func GetMainText(line string) (mainText string) {
	var obj = "{}"
	if len(line) > 1 && strings.Contains(line, "data:") {
		obj = strings.Split(line, "data: ")[1]
	}

	var d Response
	if err := json.Unmarshal([]byte(obj), &d); err != nil {
		return ""
	}

	if d.Token != "" {
		mainText = d.Token
		return mainText
	}
	return ""
}
