package main

import (
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "strconv"

  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Welcmoe!")
}

func YamaIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(yamas); err != nil {
      panic(err)
  }
}

func YamaShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  id, _ := strconv.Atoi(ps.ByName("yamaId"))
  t := RepoFindYama(id)
  if t.ID == 0 && t.Proposer == "" {
      w.WriteHeader(http.StatusNotFound)
      return
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(t); err != nil {
      panic(err)
  }
}

func YamaCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  var yama Yama

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // 1MiB
  if err != nil {
      panic(err)
  }
  defer r.Body.Close()

  if err := json.Unmarshal(body, &yama); err != nil {
      w.WriteHeader(500)
      if err := json.NewEncoder(w).Encode(err); err != nil {
          panic(err)
      }
      return
  }

  t := RepoCreateYama(yama)
  location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
  w.Header().Set("Location", location)
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(t); err != nil {
      panic(err)
  }
}

func YamaDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  id, _ := strconv.Atoi(ps.ByName("yamaId"))
  if err := RepoDestroyYama(id); err != nil {
      w.WriteHeader(http.StatusNotFound)
      if err := json.NewEncoder(w).Encode(err); err != nil {
          panic(err)
      }
      return
  }

  w.Header().Del("Content-Type")
  w.WriteHeader(204) // 204 No Content
}
