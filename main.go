package main 

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price float64  `json:"price"`
}

func getAlbums(c *gin.Context) {
	// 引数はクライアントに送信したいHTTPステータスコード。200OKを示すnet/httpパッケージの定数StatusOKを渡している。
	c.IndentedJSON(http.StatusOK, albums)
	// c.IndentedJSONで構造体をJSONにシリアライズし、レスポンスに追加する。
}

var albums = []album {
	{ID: "1", Title: "ChuchuTrain", Artist: "EXILE", Price: 5000},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 4000},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3000},
}