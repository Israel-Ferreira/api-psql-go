package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-list-go/models"
)

func Create(rw http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println("Erro ao Deserializar o Json")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.InsertTodo(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("Ocorreu o erro ao inserir: %s", err.Error()),
		}
	} else {
		resp = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("Todo inserido com sucesso! ID = %d", id),
		}
	}

	rw.Header().Add("Content-Type", "application/json")

	if err = json.NewEncoder(rw).Encode(resp); err != nil {
		log.Println("Erro ao Serializar o Json")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
