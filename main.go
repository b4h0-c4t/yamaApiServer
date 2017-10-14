package main

import (
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func main() {
  router := httprouter.New()
  router.GET("/", Logging(Index, "index"))
  router.GET("/yamas", CommonHeaders(YamaIndex, "yama-index"))
  router.GET("/yamas/:yamaId", IDShouldBeInt(YamaShow, "yama-show"))
  router.POST("/yamas", CommonHeaders(YamaCreate, "yama-create"))
  router.DELETE("/yamas/:yamaId", IDShouldBeInt(YamaDelete, "yama-delete"))

  log.Fatal(http.ListenAndServe(":8080", router))
}
