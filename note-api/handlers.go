package main

import (
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/abvarun226/notes-app/note-svc/proto"
	"github.com/go-chi/chi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is the struct for HTTP handlers.
type Handler struct {
	NoteService pb.NoteClient
}

// ListNotesByUser is the endpoint handler for GET /note/v1/notes/{user}
func (h *Handler) ListNotesByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := chi.URLParam(r, "user")

	rspList, errList := h.NoteService.ListNotes(ctx, &pb.NotesRequest{UserId: userID})
	if errList != nil {
		log.Printf("failed to list notes: %v", errList)
		renderError(w, errList)
		return
	}

	response, errMarshal := json.Marshal(rspList.Notes)
	if errMarshal != nil {
		log.Printf("failed to marshal response: %v", errMarshal)
		renderError(w, errMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	w.Header().Set("Content-Type", "application/json")
}

func renderError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	s := status.Convert(err)
	msg := ErrorResponse{Status: http.StatusInternalServerError, Mesasge: s.Message()}

	if s.Code() == codes.NotFound {
		msg.Status = http.StatusNotFound
	}

	jsonMsg, _ := json.Marshal(msg)

	http.Error(w, string(jsonMsg), msg.Status)
}

// ErrorResponse struct.
type ErrorResponse struct {
	Status  int    `json:"status"`
	Mesasge string `json:"message"`
}
