package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/bap/service"
	"example.com/bap/slack"
	"example.com/bap/util/data"

	"github.com/gorilla/mux"
)

type Handler struct {
	Srv      service.Service
	SlackSrv slack.Slack
}

var (
	NEWDIALOG = data.DialogRequest{
		TriggerId: "",
		Form: data.Dialog{
			Callback:     "new-blog",
			Title:        "New Blog",
			Label:        "OK",
			NotifyCancel: false,
			Elements: []data.Element{
				{Type: "text", Label: "Title", Name: "title"},
				{Type: "text", Label: "Tag", Name: "tag", Hint: "comma separated"},
				{Type: "text", Label: "Article", Name: "article", Subtype: "url"},
				{Type: "select", Label: "Open", Name: "open", Options: []data.Option{
					{Label: "true", Value: "true"},
					{Label: "false", Value: "false"},
				}},
				{Type: "text", Label: "Date", Name: "date", Hint: "2022/01/01"},
			},
		},
	}
	UPDATEDIALOG = data.DialogRequest{
		TriggerId: "",
		Form: data.Dialog{
			Callback:     "update-blog",
			Title:        "Update Blog",
			Label:        "OK",
			NotifyCancel: false,
			Elements: []data.Element{
				{Type: "text", Label: "Id", Name: "id"},
				{Type: "text", Label: "Title", Name: "title", Optional: true},
				{Type: "text", Label: "Tag", Name: "tag", Optional: true, Hint: "comma separated"},
				{Type: "text", Label: "Article", Name: "article", Optional: true, Subtype: "url"},
				{Type: "select", Label: "Open", Name: "open", Options: []data.Option{
					{Label: "true", Value: "true"},
					{Label: "false", Value: "false"},
				}},
				{Type: "text", Label: "Date", Name: "date", Optional: true, Hint: "2022/01/01"},
			},
		},
	}
)

func (handler *Handler) viewBlog(w http.ResponseWriter, r *http.Request) {
	blogs, err := handler.Srv.Blogs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	body, err := json.Marshal(blogs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (handler *Handler) newTrigger(w http.ResponseWriter, r *http.Request) {
	triggerId, err := handler.SlackSrv.TriggerId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var dialog = NEWDIALOG
	dialog.TriggerId = triggerId
	dialogJson, err := json.Marshal(dialog)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = handler.SlackSrv.DialogApi(bytes.NewBuffer(dialogJson))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) updateTrigger(w http.ResponseWriter, r *http.Request) {
	triggerId, err := handler.SlackSrv.TriggerId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var dialog = UPDATEDIALOG
	dialog.TriggerId = triggerId
	dialogJson, err := json.Marshal(dialog)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = handler.SlackSrv.DialogApi(bytes.NewBuffer(dialogJson))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) submit(w http.ResponseWriter, r *http.Request) {
	pl, err := handler.SlackSrv.Payload(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	var payload data.Payload
	if err = json.Unmarshal([]byte(pl), &payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Routing by callback id
	if payload.CallbackId == "new-blog" {
		var pre = &data.Blog{}
		blog := handler.SlackSrv.MergePayload(pre, &payload)
		err = handler.Srv.NewBlog(blog)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	} else if payload.CallbackId == "update-blog" {
		pre, err := handler.Srv.BlogMeta(payload.Body.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		blog := handler.SlackSrv.MergePayload(pre, &payload)
		err = handler.Srv.UpdateBlog(blog, payload.Body.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Invalid callback id")
		return
	}

	// Post response url as ok sign
	err = handler.SlackSrv.SendResponse(payload.ResponseURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) updateSubmit(payload *data.Payload) error {
	pre, err := handler.Srv.Blog(payload.Body.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	blog := handler.SlackSrv.MergePayload(pre, payload)

	return handler.Srv.UpdateBlog(blog, payload.Body.Id)
}

func (handler *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, _ := handler.SlackSrv.VerifyWebHook(r)
		if result {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			log.Println("Auth error")
			return
		}
	})
}

func main() {
	srv := service.NewService()
	slackSrv := slack.NewSlack()
	handler := Handler{Srv: srv, SlackSrv: slackSrv}

	err := handler.Srv.ConstructDB("database")
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Srv.DestructDB()
	if err = handler.Srv.ConstructDrive(); err != nil {
		log.Fatal(err)
	}
	handler.SlackSrv.Construct()

	router := mux.NewRouter()
	router.HandleFunc("/view", handler.viewBlog).Methods(http.MethodPost)
	router.HandleFunc("/new", handler.newTrigger).Methods(http.MethodPost)
	router.HandleFunc("/update", handler.updateTrigger).Methods(http.MethodPost)
	router.HandleFunc("/submit", handler.submit).Methods(http.MethodPost)

	router.Use(handler.authMiddleware)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Println("Start Server")
	if os.Getenv("GO_ENV") == "production" {
		log.Fatal(s.ListenAndServeTLS(
			"/etc/letsencrypt/live/www.example.com/fullchain.pem",
			"/etc/letsencrypt/live/www.example.com/privkey.pem",
		))
	} else {
		log.Fatal(s.ListenAndServe())
	}
}
