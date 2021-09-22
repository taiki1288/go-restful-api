package main 

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price float64  `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "ChuchuTrain", Artist: "EXILE", Price: 5000},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 4000},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3000},
}