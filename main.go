package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var me = ""
var authUrl = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
var tokenUrl = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
var redirectUri = "http://localhost:5000/login/authorized"
var clientId = "Register your app at apps.dev.microsoft.com"
var clientSecret = "Register your app at apps.dev.microsoft.com"
var scope = "User.Read"

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login/authorized", authorizedHandler)
	http.HandleFunc("/me", meHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	renderTemplate("hello.html", responseWriter, nil)
}

func loginHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		responseType := "code"
		state := "1234"

		http.Redirect(responseWriter, request, fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s", authUrl, clientId, redirectUri, responseType, scope, state), 301)
	}
}

func authorizedHandler(responseWriter http.ResponseWriter, request *http.Request) {
	accessToken := getAccessToken(request.FormValue("code"))
	me = getMe(accessToken)

	http.Redirect(responseWriter, request, "/me", 301)
}

func meHandler(responseWriter http.ResponseWriter, request *http.Request) {
	renderTemplate("me.html", responseWriter, map[string]string{"me": me})
}

func renderTemplate(fileName string, responseWriter http.ResponseWriter, data map[string]string) {
	fileTemplate, _ := template.ParseFiles("templates/" + fileName)
	fileTemplate.Execute(responseWriter, data)
}

func getAccessToken(code string) string {
	grantType := "authorization_code"

	form := url.Values{}
	form.Add("client_id", clientId)
	form.Add("client_secret", clientSecret)
	form.Add("code", code)
	form.Add("redirect_uri", redirectUri)
	form.Add("grant_type", grantType)
	form.Add("scope", scope)

	request, err := http.NewRequest("POST", tokenUrl, strings.NewReader(form.Encode()))

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatal(err)
	}

	accessToken := result["access_token"]

	return accessToken.(string)
}

func getMe(accessToken string) string {
	request, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return string(body)
}
