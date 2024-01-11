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

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Use this command to login to the onepixel url shortener.",
	Long: `Use this command to login to the onepixel url shortener.`,

	Run: func(cmd *cobra.Command, args []string) {
		var username, password string
		fmt.Println("Enter your username: ")
		fmt.Scanln(&username)
		fmt.Println("Enter your password: ")
		fmt.Scanln(&password)

		login(username, password)
		fmt.Println("login called")
	},
}

func login(username, password string) {
	credentials := struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}	{
		Email: username,
		Password: password,
	}

	jsonValue, _ := json.Marshal(credentials)
	response, err := http.Post("http://127.0.0.1:3000/api/v1/users/login", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return 
	} 

	var result struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response: ", err)
		return
	}

	if result.Token != "" {
      // Save the token to a file
      err := ioutil.WriteFile("token.txt", []byte(result.Token), 0644)
      if err != nil {
          fmt.Println("Error saving token:", err)
          return
      }
      fmt.Println("Login successful. Token saved.")
  } else {
      fmt.Println("No token received.")
  }
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
