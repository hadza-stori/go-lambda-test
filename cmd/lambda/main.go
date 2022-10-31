package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/core/entities"
	"testProject/core/usecases"
)

var ginLambda *ginadapter.GinLambda

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	g := gin.New()

	g.GET("/og", func(c *gin.Context) {
		names, err := usecases.GetPokemonList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"results": names})
	})

	g.GET("/pokemon/:name", func(c *gin.Context) {
		pkm := entities.Pokemon{Name: c.Param("name")}
		data, err := usecases.GetPokemon(pkm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pokemon": data})
	})

	ginLambda = ginadapter.New(g)
	lambda.Start(handler)
}
