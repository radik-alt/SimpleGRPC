package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	baseUrl       = "localhost:8081"
	createPostfix = "/notes"
	getPostfix    = "/notes/%id"
)

type NoteInfo struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

type Note struct {
	ID        int64     `json:"id"`
	Info      NoteInfo  `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SyncMap struct {
	elem map[int64]*Note
	mu   sync.RWMutex
}

var notes = &SyncMap{
	elem: make(map[int64]*Note),
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	info := &NoteInfo{}

	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	note := &Note{
		ID:        rand.Int63(),
		Info:      *info,
		CreatedAt: now,
		UpdatedAt: now,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	notes.mu.Lock()
	defer notes.mu.Unlock()

	notes.elem[note.ID] = note
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "id")
	id, err := parseNoteId(noteID)

	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	notes.mu.RLock()
	defer notes.mu.RUnlock()

	note, ok := notes.elem[id]
	if !ok {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func parseNoteId(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func main() {
	router := chi.NewRouter()

	router.Post(createPostfix, createNoteHandler)
	router.Get(getPostfix, getNoteHandler)

	err := http.ListenAndServe(baseUrl, router)
	if err != nil {
		log.Fatal(err)
	}
}
