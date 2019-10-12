package messages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"

	"github.com/n7down/go-exercises/Grpc/internal/pb/messages"
)

const (
	messagesPort = "8081"
)

type Messages struct {
	Client messages.HelloServiceClient
}

func NewMessages(c messages.HelloServiceClient) *Messages {
	return &Messages{Client: c}
}

type HelloRequest struct {
	Name string `json: "name" binding:"required"`
}

func (r *HelloRequest) Validate() url.Values {
	errs := url.Values{}
	if r.Name == "" {
		errs.Add("name", "The name field is required!")
	}
	return errs
}

type HelloResponse struct {
	Message string `json: "message"`
}

func (m Messages) HelloHandler(c *gin.Context) {
	var (
		req HelloRequest
		res HelloResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// validate input data
	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := m.Client.SayHello(c, &messages.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = HelloResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}