package datastruct

type Album struct {
	ID     int    `json:"id"`
	Tittle string `json:"tittle"`
	Artist string `json:"artist"`
	Price  int    `json:"price"`
}
