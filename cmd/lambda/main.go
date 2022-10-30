package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"time"
)

var ginLambda *ginadapter.GinLambda

type Pokemon struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
}

type PokemonList struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []PokemonListItem `json:"results"`
}

type PokemonListItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

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
		c.JSON(http.StatusOK, gin.H{"results": names})
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

func ApiRequest(url string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "go-test")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	return resp.Body, err
}

func GetPokemonNamesFromAPI() (PokemonList, error) {

	url := "https://pokeapi.co/api/v2/pokemon?limit=151"
	var pokemonList PokemonList

	resp, err := ApiRequest(url)

	err = json.NewDecoder(resp).Decode(&pokemonList)

	log.Println(pokemonList)

	return pokemonList, err
}

func GetPokemonDataFromAPI(c *gin.Context) (Pokemon, error) {

	url := "https://pokeapi.co/api/v2/pokemon/" + c.Param("name")
	var pokemon Pokemon

	resp, err := ApiRequest(url)
	err = json.NewDecoder(resp).Decode(&pokemon)

	log.Println(pokemon)

	return pokemon, err
}
