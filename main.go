package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pokemon struct {
    JpName     string
    EnName     string
    ImgUrl     string
    FlavorText string
}

var pokedex = []Pokemon{
    {
        JpName:     "フシギダネ",
        EnName:     "Bulbasaur",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png",
        FlavorText: "うまれたときから　せなかに ふしぎな　タネが　うえてあって からだと　ともに　そだつという。",
    },
    {
        JpName:     "フシギソウ",
        EnName:     "Ivysaur",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png",
        FlavorText: "つぼみが　せなかに　ついていて ようぶんを　きゅうしゅうしていくと おおきな　はなが　さくという。",
    },
    {
        JpName:     "フシギバナ",
        EnName:     "Venusaur",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/3.png",
        FlavorText: "おおきな　はなびらを　ひろげ たいようの　ひかりを　あびていると からだに　げんきが　みなぎっていく。",
    },
    {
        JpName:     "ヒトカゲ",
        EnName:     "Charmander",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/4.png",
        FlavorText: "しっぽの　ほのおは　ヒトカゲの せいめいりょくの　あかし。 げんきだと　さかんに　もえさかる。",
    },
    {
        JpName:     "リザード",
        EnName:     "Charmeleon",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/5.png",
        FlavorText: "しっぽを　ふりまわして　あいてを なぎたおし　するどい　ツメで ズタズタに　ひきさいてしまう。",
    },
    {
        JpName:     "リザードン",
        EnName:     "Charizard",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/6.png",
        FlavorText: "くちから　しゃくねつの　ほのおを はきだすとき　しっぽの　さきは より　あかく　はげしく　もえあがる。",
    },
    {
        JpName:     "ゼニガメ",
        EnName:     "Squirtle",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/7.png",
        FlavorText: "こうらに　とじこもり　みを　まもる。 あいての　すきを　みのがさず みずを　ふきだして　はんげきする。",
    },
    {
        JpName:     "カメール",
        EnName:     "Wartortle",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/8.png",
        FlavorText: "ポカンと　あたまを　たたかれるとき こうらに　ひっこんで　よける。でも ちょっとだけ　しっぽが　でているよ。",
    },
    {
        JpName:     "カメックス",
        EnName:     "Blastoise",
        ImgUrl:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/9.png",
        FlavorText: "からだが　おもたく　のしかかって あいてを　きぜつさせる。 ピンチの　ときは　カラに　かくれる。",
    },
}


func main() {
    e := echo.New()

    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Response().Header().Set("Access-Control-Allow-Origin", "*")
            c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
            c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
            if c.Request().Method == "OPTIONS" {
                return c.NoContent(http.StatusNoContent)
            }
            return next(c)
        }
    })

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to the API server!")
    })

    e.GET("/pokemon", func(c echo.Context) error {
        var startIndex, endIndex int
        var err error

        startIndexStr := c.QueryParam("startIndex")
        if startIndexStr == "" {
            startIndex = 0
        } else {
            startIndex, err = strconv.Atoi(startIndexStr)
            if err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for startIndex"})
            }
        }

        endIndexStr := c.QueryParam("endIndex")
        if endIndexStr == "" {
            endIndex = len(pokedex) - 1
        } else {
            endIndex, err = strconv.Atoi(endIndexStr)
            if err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for endIndex"})
            }
        }

        if startIndex < 0 || endIndex >= len(pokedex) || startIndex > endIndex {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid index range"})
        }

        return c.JSON(http.StatusOK, pokedex[startIndex:endIndex+1])
    })

    e.Logger.Fatal(e.Start(":8080"))
}