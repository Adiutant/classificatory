package payload

import (
	"fmt"
	"net"
)

type Connector interface {
	SendData(string) error
	ReceiveData() (string, error)
}
type ClassificatoryCommand int

const (
	CheckHealth ClassificatoryCommand = 1 << 1
	Add                               = 1 << 2
	Remove                            = 1 << 3
	List                              = 1 << 4
	Text                              = 1 << 5
)

type ClassificationConnector struct {
	commands   map[ClassificatoryCommand]string
	connection net.Conn
}

func NewClassificationConnector(conn net.Conn) (*ClassificationConnector, error) {
	classificationConnector := ClassificationConnector{
		commands:   make(map[ClassificatoryCommand]string),
		connection: conn,
	}
	return &classificationConnector, nil
}
func (c *ClassificationConnector) ApplyCommand(command ClassificatoryCommand, payload string) {
	switch command {
	case CheckHealth:
		err := c.SendData("CheckHealth")
		if err != nil {
			fmt.Println(err)
			return
		}
	case Add:
		data := "Add"
		data += "\n"
		data += payload
		err := c.SendData(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	case Remove:
		data := "Remove"
		data += "\n"
		data += payload
		err := c.SendData(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	case List:
		err := c.SendData("List")
		if err != nil {
			fmt.Println(err)
			return
		}
	case Text:
		data := "Text"
		data += "\n"
		data += payload
		err := c.SendData(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (c *ClassificationConnector) SendData(str string) error {
	_, err := net.Conn.Write(c.connection, []byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClassificationConnector) ReceiveData() (string, error) {
	var buffer []byte
	_, err := net.Conn.Read(c.connection, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
