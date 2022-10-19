package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var ginLambda *ginadapter.GinLambda

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	g := gin.New()

	g.GET("/og", func(c *gin.Context) {
		names, err := GetPokemonNamesFromAPI()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pokemon": names})
	})

	g.GET("/pokemon/:name", func(c *gin.Context) {
		data, err := GetPokemonDataFromAPI(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pokemon": data})
	})

	ginLambda = ginadapter.New(g)
	lambda.Start(handler)
}

func ApiRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "go-test")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func GetPokemonNamesFromAPI() ([]string, error) {

	url := "https://pokeapi.co/api/v2/pokemon?limit=151"
	var pokemon []string

	var result map[string]interface{}

	resp, err := ApiRequest(url)

	json.Unmarshal(resp, &result)

	pokemonList := result["results"].([]interface{})

	for _, p := range pokemonList {
		pokemon = append(pokemon, p.(map[string]interface{})["name"].(string))
	}

	return pokemon, err
}

// Language: go
// Get pokemon data from pokeapi and return it as a JSON

func GetPokemonDataFromAPI(c *gin.Context) (map[string]interface{}, error) {

	url := "https://pokeapi.co/api/v2/pokemon/" + c.Param("name")
	var pokemon map[string]interface{}

	resp, err := ApiRequest(url)

	json.Unmarshal(resp, &pokemon)

	return pokemon, err
}
