package controllers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/giancarloha/go-rest-api/database"
	"github.com/giancarloha/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Home Page")
	} else {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		slog.Error("Cliente esta tentando utilizar metodo não aceito")
	}
}

func TodosMangas(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var m []models.Manga
		database.DB.Find(&m)
		json.NewEncoder(w).Encode(m)
	} else {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		slog.Error("Cliente esta tentando utilizar metodo não aceito")
	}

}

func RetornaUmManga(w http.ResponseWriter, r *http.Request) {
	urlid := r.URL.Query().Get("id")
	fmt.Println(urlid)
	vars := mux.Vars(r)
	id := vars["id"]
	var manga models.Manga

	database.DB.First(&manga, id)
	json.NewEncoder(w).Encode(manga)
}

func AddManga(w http.ResponseWriter, r *http.Request) {
	var novoManga models.Manga

	err := json.NewDecoder(r.Body).Decode(&novoManga)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON: "+err.Error(), http.StatusBadRequest)
		slog.Error("Erro ao decodificar o JSON")
		return
	}

	if strings.TrimSpace(novoManga.Nome) == "" {
		http.Error(w, "O campo Nome é obrigatório", http.StatusBadRequest)
		slog.Error("O campo Nome é obrigatório")
		return
	}
	if novoManga.Preco == 0 {
		http.Error(w, "O campo Preço é obrigatório e não pode ser zero", http.StatusBadRequest)
		slog.Error("O campo Preço é obrigatório e não pode ser zero")
		return
	}
	if novoManga.Vnumero == 0 {
		http.Error(w, "O campo numero do volume é obrigatório e não pode ser zero", http.StatusBadRequest)
		slog.Error("O campo numero do volume é obrigatório e não pode ser zero")
		return
	}
	if novoManga.Mespub == 0 {
		http.Error(w, "O campo Mes de publicação é obrigatório e não pode ser zero", http.StatusBadRequest)
		slog.Error("O campo Mes de publicação é obrigatório e não pode ser zero")
		return
	}
	if novoManga.Anopub == 0 {
		http.Error(w, "O campo Ano de publicação é obrigatório e não pode ser zero", http.StatusBadRequest)
		slog.Error("O campo Mes de publicação é obrigatório e não pode ser zero")
		return
	}
	if strings.TrimSpace(novoManga.Img) == "" {
		http.Error(w, "O campo link da Imagem é obrigatório", http.StatusBadRequest)
		slog.Error("O campo Mes de publicação é obrigatório e não pode ser zero")
		return
	}

	database.DB.Create(&novoManga)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoManga)
}

func DeleteManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var manga models.Manga
	database.DB.Delete(&manga, id)
	json.NewEncoder(w).Encode(manga)
}

func EditManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var manga models.Manga
	database.DB.First(&manga, id)
	json.NewDecoder(r.Body).Decode(&manga)
	database.DB.Save(&manga)
	json.NewEncoder(w).Encode(manga)
}
