package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type fighter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Backbone string `json:"backbone"`
}

var albums = []album{
	{ID: "1", Title: "ChuchuTrain", Artist: "EXILE", Price: 5000},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 4000},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3000},
}

// 個人的な好きな総合格闘家一覧
var fighters = []fighter {
	{ ID: "1", Name: "堀口恭司", Age: 30, Backbone: "Traditional Karate" },
	{ ID: "2", Name: "朝倉未来", Age: 29, Backbone: "Fighting"},
	{ ID: "3", Name: "朝倉海", Age: 28, Backbone: "Karate" },
	{ ID: "4", Name: "井上直樹", Age: 24, Backbone: "Karate" },
	{ ID: "5", Name: "金太郎", Age: 28, Backbone: "Kickboxing" },
	{ ID: "6", Name: "斉藤裕", Age: 33, Backbone: "Brazilian Jiu-Jitsu" },
}

func main() {
	router := gin.Default()
	// Defaultを使用してGinルーターを初期化
	router.GET("/albums", getAlbums)
	// GET関数を使用してGET HTTPメソッドと/albumsパスをハンドラ関数に関連付ける。getAlbums関数を渡す。
	router.GET("/albums/:id", getAlbumByID)
	// /albums/:idのパスとgetAlbumByID関数を関連付ける。getAlbumByID関数を返す。
	router.GET("/fighters", getFighters)
	router.POST("/albums", postAlbums)
	// /albumsパスのPOSTメソッドとpostAlbums関数を関連付ける。
	router.Run("localhost:8080")
	// Run関数を使って、ルータをhttp.Serverに接続し、サーバを起動する。
}

func getAlbums(c *gin.Context) {
	// 引数はクライアントに送信したいHTTPステータスコード。200OKを示すnet/httpパッケージの定数StatusOKを渡している。
	c.IndentedJSON(http.StatusOK, albums)
	// c.IndentedJSONで構造体をJSONにシリアライズし、レスポンスに追加する。
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		// Context.BindJSON を使用して、リクエストボディをnewAlbumにバインドする。
		return 
	}

	albums = append(albums, newAlbum)
	// 第一引数ではスライス、第二引数では追加する値をとっている。
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// クライアントから送られてきたidパラメータと一致するIDを持つアルバムを探す関数。
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		// スライスのalbumsをループして、IDフィールドの値がidパラメータの値と一致するものを探す。
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			// アルバム構造体をJSONにシリアライズし、200OKのHTTPコードでレスポンスとして返します。
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	// アルバムが見つからない場合にHTTP404エラーを返している。
}

func getFighters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fighters)
}

