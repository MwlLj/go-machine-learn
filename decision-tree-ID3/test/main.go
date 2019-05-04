package main

import (
	// "fmt"
	"./buffer"
	"./handler"
	db "./infer_db"
	_ "github.com/mattn/go-sqlite3"
)

type CServer struct {
	dbHandler db.CDbHandler
	buffer    *buffer.CData
	handler   *handler.CHandler
}

func (this *CServer) Start() {
	this.dbHandler.Connect(
		"",
		0,
		"",
		"",
		"infer.db",
		"sqlite3",
	)
	this.dbHandler.Create()
	defer this.dbHandler.Disconnect()
	this.initBuffer()
	this.initHandler()
}

func (this *CServer) initBuffer() {
	this.buffer = buffer.NewData(this.dbHandler)
}

func (this *CServer) initHandler() {
	this.handler = handler.NewHandler(this.buffer)
	this.handler.PrintMethod()
	this.handler.InputHandle()
}

func main() {
	server := CServer{}
	server.Start()
}
