package http_server

import (
	"fmt"
	"gin_webserver/db_helper"
	"gin_webserver/payload"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type HttpPayloadServer struct {
	engine    *gin.Engine
	connector payload.Connector
	dbHelper  *db_helper.DBHelper
	logger    *logrus.Logger
}
type Request struct {
	Command string `json:"command,omitempty"`
	Payload string `json:"payload,omitempty"`
}

func NewPayloadServer(payloadConn net.Conn) (*HttpPayloadServer, error) {
	payloadConnector, err := payload.NewClassificationConnector(payloadConn)
	if err != nil {
		return nil, err
	}
	httpPayloadServer := HttpPayloadServer{
		engine:    gin.Default(),
		connector: payloadConnector,
		logger:    logrus.New(),
	}
	return &httpPayloadServer, nil
}
func (s *HttpPayloadServer) SetRoutes() {
	s.engine.POST("/request-payload", func(c *gin.Context) {
		var req Request
		err := c.BindJSON(&req)
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("bad json request"))
		if err != nil {
			return
		}
		s.connector.SendData()
		c.AbortWithError(http.StatusNotImplemented, fmt.Errorf("request-payload not implemented"))
		return
	})

}

func (s *HttpPayloadServer) Run() error {
	err := s.engine.Run(":8080")
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
