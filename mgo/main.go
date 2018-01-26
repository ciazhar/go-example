package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Person struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name"`
	Phone string        `json:"phone"`
}

type Server struct {
	session *mgo.Session
}

func NewServer() (*Server, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Server{session: session}, nil
}

func (s *Server) Close() {
	s.session.Close()
}

func (s *Server) personList(w http.ResponseWriter, r *http.Request) {
	session := s.session.Copy()
	defer session.Close()

	result := []Person{}
	c := session.DB("spring").C("user")
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	j, _ := json.Marshal(result)
	w.Write(j)
}

func main() {
	srv, err := NewServer()
	if err != nil {
		panic(err)
	}
	defer srv.Close()

	http.HandleFunc("/people", srv.personList)
	http.ListenAndServe(":8080", nil)
}