package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Fruit struct {
    Name             string `json:"name"`
    Description      string `json:"description"`
    NutritionalValue string `json:"nutritionalValue"`
    Origin           string `json:"origin"`
}

var dataList = []Fruit{
    {Name: "Apple", Description: "A sweet, edible fruit produced by an apple tree.", NutritionalValue: "52 kcal per 100g", Origin: "Central Asia"},
    {Name: "Banana", Description: "An elongated, edible fruit produced by several kinds of large herbaceous flowering plants.", NutritionalValue: "89 kcal per 100g", Origin: "Southeast Asia"},
    {Name: "Cherry", Description: "A small, round stone fruit that is typically bright or dark red.", NutritionalValue: "50 kcal per 100g", Origin: "Europe and Asia"},
    {Name: "Date", Description: "The fruit of the date palm, which is a sweet and chewy fruit.", NutritionalValue: "277 kcal per 100g", Origin: "Middle East"},
    {Name: "Elderberry", Description: "A small, dark-purple fruit that grows in clusters and is known for its medicinal properties.", NutritionalValue: "73 kcal per 100g", Origin: "Europe, Africa, and parts of Asia"},
    {Name: "Fig", Description: "A soft fruit with a thin skin that can be eaten ripe or dried.", NutritionalValue: "74 kcal per 100g", Origin: "Western Asia"},
    {Name: "Grape", Description: "A fruit, botanically a berry, of the deciduous woody vines of the flowering plant genus Vitis.", NutritionalValue: "69 kcal per 100g", Origin: "Near East"},
    {Name: "Honeydew", Description: "A fruit that has a smooth, pale outer skin and sweet, green flesh inside.", NutritionalValue: "36 kcal per 100g", Origin: "West Africa"},
    {Name: "Indian Fig", Description: "Also known as prickly pear, a species of cactus that produces an edible fruit.", NutritionalValue: "41 kcal per 100g", Origin: "Mexico"},
    {Name: "Jackfruit", Description: "The largest fruit that grows on a tree, with a distinctive sweet and fruity aroma.", NutritionalValue: "95 kcal per 100g", Origin: "South India"},
}

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to the API server!")
    })

    e.GET("/fruits", func(c echo.Context) error {
        startIndex, err := strconv.Atoi(c.QueryParam("startIndex"))
        if err != nil {
            startIndex = 0
        }
        endIndex, err := strconv.Atoi(c.QueryParam("endIndex"))
        if err != nil {
            endIndex = len(dataList) - 1
        }

        if startIndex < 0 || endIndex >= len(dataList) || startIndex > endIndex {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid index range"})
        }

        return c.JSON(http.StatusOK, dataList[startIndex:endIndex+1])
    })

    e.Logger.Fatal(e.Start(":5100"))
}
