package controllers

import (
	"github.com/icattlecoder/tgw"
	. "github.com/icattlecoder/tgw/example/models"
	"log"
)

type Server struct {
	//
}

func NewServer( /**/) *Server {
	return &Server{}
}

type TestArgs struct {
	Msg string
}

func (s *Server) Hello(args TestArgs, env tgw.ReqEnv) {

	env.RW.Write([]byte(args.Msg))
	err := env.Session.Set("key", args)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) Session(env tgw.ReqEnv) {

	keyValue :=""
	err := env.Session.Get("key",&keyValue)
	if err != nil {
		log.Println(err)
	}
	env.RW.Write([]byte(keyValue))
	
}

func (s *Server) Index(env tgw.ReqEnv) (data map[string]interface{}) {
	data = map[string]interface{}{}
	author := Author{
		Name:  "icattlecoder",
		Email: []string{"icattlecoder@gmail.com", "iwangming@hotmail.com"},
		QQ:    "405283013",
		Blog:  "http://blog.segmentfault.com/icattlecoder",
	}
	data["author"] = author
	data["index"] = true
	env.Session.Set("hello", "Tiny GO Web")
	return
}

func (s *Server) Doc() (data map[string]interface{}) {
	data = map[string]interface{}{}
	data["doc"] = true
	return
}

func (s *Server) Json() (data map[string]interface{}) {
	data = map[string]interface{}{}
	author := Author{
		Name:  "icattlecoder",
		Email: []string{"icattlecoder@gmail.com", "iwangming@hotmail.com"},
		QQ:    "405283013",
		Blog:  "http://blog.segmentfault.com/icattlecoder",
	}
	data["author"] = author
	return
}

func (s *Server) AdminIndex() (data map[string]interface{}) {
	return
}
