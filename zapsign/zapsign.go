package zapsign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qbem-repos/zapsign-sdk-go/zapsign/filter"
	"github.com/qbem-repos/zapsign-sdk-go/zapsign/request"
	"github.com/qbem-repos/zapsign-sdk-go/zapsign/response"
)

const baseUrl string = "https://api.zapsign.com.br/api/v1/"

// ZapsignAPIClient api interface
type ZapsignAPI interface {
	CreateDocumentViaUpload(body *request.Docs) (*response.Docs, error)
	CreateDocumentViaModel(body *request.Model) (*response.Docs, error)
	AddAttachment(body *request.Attachment) (*response.Attachment, error)
	GetADocument(filter *filter.DocsFilter) (*response.Docs, error)
	ListDocuments(*filter.DocsFilter) (*response.Docs, error)
	RemoveDocument(*filter.DocsFilter) (*response.Docs, error)
	PlaceSignature(body *request.Docs) (*response.Docs, error)
	AddWebhook(body *request.Webhook) (*response.Docs, error)
	RemoveWebhook(filter *filter.WebhookFilter) (*response.Docs, error)
}

// ZapsignClient is a representation to client for zapsign
type ZapsignClient struct {
	token string // user token for use zapsign api
}

// CreateDocumentViaUpload create document via upload
// link: https://docs.zapsign.com.br/v/english/documentos/criar-documento
func (z *ZapsignClient) CreateDocumentViaUpload(body *request.Docs) (*response.Docs, error) {
	var bytesBody []byte
	var err error

	bytesBody, err = json.Marshal(body)

	if err != nil {
		return nil, err
	}

	res, err := http.Post("", "application/json", bytes.NewBuffer(bytesBody))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusBadRequest {
		return nil, fmt.Errorf(res.Status)
	}

	resBody := new(response.Docs)
	err = json.NewDecoder(res.Body).Decode(&resBody)

	if err != nil {
		return nil, err
	}

	return resBody, nil
}

// CreateDocumentViaModel
func (z *ZapsignClient) CreateDocumentViaModel(body *request.Model) (*response.Docs, error) {
	return nil, nil
}

// AddAttachment
func (z *ZapsignClient) AddAttachment(body *request.Attachment) (*response.Attachment, error) {
	return nil, nil
}

// GetADocument
func (z *ZapsignClient) GetADocument(filter *filter.DocsFilter) (*response.Docs, error) {
	return nil, nil
}

// ListDocuments
func (z *ZapsignClient) ListDocuments(*filter.DocsFilter) (*response.Docs, error) {
	return nil, nil
}

// RemoveDocument
func (z *ZapsignClient) RemoveDocument(*filter.DocsFilter) (*response.Docs, error) {
	return nil, nil
}

// PlaceSignature
func (z *ZapsignClient) PlaceSignature(body *request.Docs) (*response.Docs, error) {
	return nil, nil
}

// AddWebhook
func (z *ZapsignClient) AddWebhook(body *request.Webhook) (*response.Docs, error) {
	return nil, nil
}

// RemoveWebhook
func (z *ZapsignClient) RemoveWebhook(filter *filter.WebhookFilter) (*response.Docs, error) {
	return nil, nil
}

// NewClient create a new Zapsign Client instance
func NewClient(token string) (*ZapsignClient, error) {
	if token == "" {
		return nil, fmt.Errorf("token is not should be empty")
	}
	return &ZapsignClient{token: token}, nil
}

/* Zapsign Wenhook*/

// ZapsignWebhook
// type ZapsignWebhook interface {
// 	UseThisPath(path string)
// 	ListenAndServe(port string) error
// }

// // ZapsignClient is a representation to client for zapsign
// type ZapsignWebhookServer struct{}

// // NewClient create a new Zapsign Client instance
// func NewWenhookServer(token string) (*ZapsignClient, error) {
// 	if token == "" {
// 		return nil, fmt.Errorf("token is not should be empty")
// 	}
// 	return &ZapsignClient{token: token}, nil
// }
