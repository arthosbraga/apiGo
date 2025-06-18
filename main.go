package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Importe o pacote docs gerado pelo swag
	_ "api/docs"
)

// Article representa a estrutura de dados de um artigo.
type Article struct {
	ID      string "1"
	Title   string `json:"title" example:"Título do Artigo"`
	Content string `json:"content" example:"Este é o conteúdo do artigo."`
}

// Chave secreta para assinar o token. Em um app real, use uma variável de ambiente!
var jwtKey = []byte("minha_chave_super_secreta")

// Claims são as informações que você armazena no token.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// AuthMiddleware é o nosso middleware para verificar o token JWT.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cabeçalho de autorização não encontrado"})
			c.Abort() // Impede a execução dos próximos handlers
			return
		}

		// O formato esperado é "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Se não havia o prefixo "Bearer "
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token de autorização inválido"})
			c.Abort()
			return
		}

		claims := &Claims{}

		// Analisa e valida o token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Verifica se o método de assinatura é o esperado
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Adiciona o nome de usuário ao contexto para uso posterior nos handlers
		c.Set("username", claims.Username)

		// Continua para o próximo handler
		c.Next()
	}
}

// @title API de Exemplo com Swagger
// @version 1.0
// @description Esta é uma API de exemplo criada em Go com Gin e documentada com Swagger.
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte da API
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer" seguido de um espaço e o token.

// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	v1.Use(AuthMiddleware())
	{
		articles := v1.Group("/articles")
		{
			articles.GET(":id", GetArticleByID)
		}
	}

	// Rota para a documentação do Swagger
	// O URL será http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}

func AuthRequired(c *gin.Context) gin.HandlerFunc {
	fmt.Println("is logged.")
	return c.Handler()
}

// GetArticleByID localiza um artigo pelo seu ID.
// @Summary      Mostra um artigo
// @Description  Obtém um artigo pelo seu ID
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID do Artigo"
// @Success      200  {object}  Article
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /articles/{id} [get]
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	// Lógica de busca simulada
	if id == "1" {
		article := Article{
			ID:      "1",
			Title:   "Aprendendo Go e Swagger",
			Content: "A integração é mais simples do que parece!",
		}
		c.JSON(http.StatusOK, article)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Artigo não encontrado"})
}
