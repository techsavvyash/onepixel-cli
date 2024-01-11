/*
Copyright © 2024 Yash Mittal <yash@techsavvyash.dev>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Use this command to create random shortURLs.",
	Long: `Use this command to create random shortURLs.`,
	Run: func(cmd *cobra.Command, args []string) {

		data, _ := cmd.Flags().GetString("data")
		slug, _ := cmd.Flags().GetString("short")
		
		if data == "" {
			fmt.Println("Please provide a url to shorten")
		}else if slug != "" && data != "" {
			createShortURL(data, slug)
			return
		} else {
			createRandomURL(data)
		}

		fmt.Println("url called")
	},
}

func createRandomURL(long_url string) {
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Println("Error reading token file: ", err)
		return
	}

	reqBody := struct {
		LongURL string `json:"long_url"`
	} {
		LongURL: long_url,
	}

	jsonValue, err := json.Marshal(reqBody)

	// make the API call
	apiURL := "http://127.0.0.1:3000/api/v1/urls"
	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", string(token))

	// make the request

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}

	defer response.Body.Close()

	// Processing the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Printf("Response: %s\n", responseBody)
}

func createShortURL(long_url, slug string) {
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Println("Error reading token file: ", err)
		return
	}

	reqBody := struct {
		LongURL string `json:"long_url"`
	} {
		LongURL: long_url,
	}

	jsonValue, err := json.Marshal(reqBody)

	// make the API call
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("http://127.0.0.1:3000/api/v1/urls/%s", slug))
	apiURL := sb.String()
	fmt.Println("apiURL: ", apiURL)
	request, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", string(token))

	// make the request

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}

	defer response.Body.Close()

	// Processing the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Printf("Response: %s\n", responseBody)
}

func init() {
	rootCmd.AddCommand(urlCmd)
	// Here you will define your flags and configuration settings.
	urlCmd.Flags().String("data", "", "Data to send with the request")
	urlCmd.Flags().String("short", "", "custom short slug")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
