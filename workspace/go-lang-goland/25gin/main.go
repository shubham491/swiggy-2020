package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"

	//"net/http"
)

type Order struct{
	OrderId int `json:"orderid"`
	CustomerName string `json:"custname"`
	OrderReview string `string:"review"`
}

func HomePage(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello world from swiggy",
	})
}

var orders=[] Order{
	{
		101,"praneeth","good",
	},
	{
		102,"somu","good",
	},
	{
		103,"bhavi","bad",
	},
}

func GetOrders(c *gin.Context){
	c.JSON(200,orders)

}
func GetOrder(c *gin.Context){
	orderId:= c.Param("orderId")
	order:= c.Param("order")
	c.JSON(http.StatusOK, gin.H{
		"OrderId":orderId,
		"Order":order,
	})

}

func PostOrder(c *gin.Context){
	body:=c.Request.Body
	content, err:= ioutil.ReadAll(body)
	if err != nil{
		fmt.Println("Sorry no content found",err.Error())

	} else{
		fmt.Println(content)
		c.JSON(http.StatusCreated, gin.H{
			"message":string(content),
		})
	}

}

func GetSpecificOrderByQuery(c *gin.Context){
	minNumberoforders:= c.Query("minNoOfOrders")
	showOrder:=c.Query("showOrder")
	c.JSON(200,gin.H{
		"minNoOfOrders":minNumberoforders,
		"showOrder":showOrder,
	})
}

func main() {
	//router:=gin.new()//this will not carry middle tier works
	router := gin.Default()

	//ipaddress:8080/ping
	api:=router.Group("/api")
	api.GET("/",HomePage)
	api.GET("/orders",GetOrders)
	api.GET("/order/:orderId/:order",GetOrder)//using colon to denote path params
	api.GET("/orders/query",GetSpecificOrderByQuery)
	api.POST("/order",PostOrder)
	router.Run(":5665") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
