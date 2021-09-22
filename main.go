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

var albums = []album{
	{ID: "1", Title: "ChuchuTrain", Artist: "EXILE", Price: 5000},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 4000},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3000},
}

func main() {
	router := gin.Default()
	// Defaultを使用してGinルーターを初期化
	router.GET("/albums", getAlbums)
	// GET関数を使用してGET HTTPメソッドと/albumsパスをハンドラ関数に関連付ける。getAlbums関数を渡す。
	router.Run("localhost:8080")
	// Run関数を使って、ルータをhttp.Serverに接続し、サーバを起動する。
}