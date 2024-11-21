package models

type Manga struct {
	Id      int     `json:"id"`
	Nome    string  `json:"nome"`
	Preco   float32 `json:"preco"`
	Vnumero int32   `json:"vnumero"`
	Mespub  int16   `json:"mespub"`
	Anopub  int32   `json:"anopub"`
	Img     string  `json:"img"`
}

var Mangas []Manga
