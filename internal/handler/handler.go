package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/as-go/first/internal/models"
	"github.com/gorilla/mux"
)

type handler struct {
	userService userService
	router      *mux.Router
}

type Config struct {
	UserService userService
}

func New(cfg Config) *handler {
	h := &handler{
		router:      mux.NewRouter(),
		userService: cfg.UserService,
	}

	return h
}

func (h *handler) Handler() http.Handler {
	h.router.HandleFunc("/user", h.createUser).Methods(http.MethodPost)
	h.router.HandleFunc("/user/{id}", h.getUser).Methods(http.MethodGet)
	h.router.HandleFunc("/user/{id}", h.updateUser).Methods(http.MethodPatch)
	h.router.HandleFunc("/user/{id}", h.deleteUser).Methods(http.MethodDelete)

	return h.router
}

func (h *handler) createUser(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	var newUser models.CrateUser
	if err := json.Unmarshal(body, &newUser); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userService.Create(ctx, newUser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnUser(res, user)
}

func (h *handler) getUser(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := getId(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userService.FindByID(ctx, id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	returnUser(res, user)
}

func (h *handler) updateUser(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, err := getId(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	var updUser models.UpdateUser
	if err := json.Unmarshal(body, &updUser); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userService.Update(ctx, id, updUser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnUser(res, user)
}

func (h *handler) deleteUser(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, err := getId(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.userService.Delete(ctx, id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

func returnUser(res http.ResponseWriter, user models.User) {
	response, err := json.Marshal(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.Write(response)
	res.WriteHeader(http.StatusCreated)
}

func getId(req *http.Request) (int, error) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		return 0, errors.New("id is required")
	}

	return strconv.Atoi(id)
}
