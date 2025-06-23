package controllers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/giancarloha/go-rest-api/database"
	"github.com/giancarloha/go-rest-api/models"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Replace with your allowed origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Home Page")
	} else {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		slog.Error("Cliente esta tentando utilizar metodo não aceito")
	}
}

func TodosMangas(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodGet {
		var m []models.Manga
		database.DB.Find(&m)
		json.NewEncoder(w).Encode(m)
	} else {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		slog.Error("Cliente esta tentando utilizar metodo não aceito")
	}

}

func TodosMangasMes(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodGet {
		mes := r.PathValue("mes")
		ano := r.PathValue("ano")
		var mangas []models.Manga

		database.DB.Where("anopub = ? AND mespub = ?", ano, mes).Find(&mangas)
		json.NewEncoder(w).Encode(mangas)
	}
}

func MangaOperations(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		RetornaUmManga(w, r, id)
	case http.MethodDelete:
		DeleteManga(w, r, id)
	case http.MethodPut, http.MethodPatch:
		EditManga(w, r, id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func RetornaUmManga(w http.ResponseWriter, r *http.Request, id string) {
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

func DeleteManga(w http.ResponseWriter, r *http.Request, id string) {
	var manga models.Manga
	json.NewDecoder(r.Body).Decode(&manga)
	database.DB.Delete(&manga, id)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(manga)

}

func EditManga(w http.ResponseWriter, r *http.Request, id string) {
	var manga models.Manga
	database.DB.First(&manga, id)
	json.NewDecoder(r.Body).Decode(&manga)
	database.DB.Save(&manga)
	json.NewEncoder(w).Encode(manga)
}
