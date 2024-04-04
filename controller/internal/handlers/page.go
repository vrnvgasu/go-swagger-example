package handlers

import (
	apiclient "controller-service/client"
	"controller-service/client/operations"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Handler struct {
	HttpClient *http.Client
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	transport := httptransport.New(apiclient.DefaultHost+":8090", apiclient.DefaultBasePath, []string{"http"})
	transport.Consumers["application/pdf"] = runtime.ByteStreamConsumer()
	client := apiclient.New(transport, strfmt.Default)

	file, _, err := r.FormFile("upfile1")
	if err != nil {
		return
	}
	defer file.Close()

	postSendParams := operations.NewPostSendParams()
	postSendParams.Upfile1 = runtime.NamedReader("upfile1", file)

	resp, err := client.Operations.PostSend(postSendParams, w)
	if err != nil {
		log.Fatal(err)
	}

	payload := resp.GetPayload()
	_ = payload

	w.Write([]byte("OK"))
	return
}

func (h *Handler) Web(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/form.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}
