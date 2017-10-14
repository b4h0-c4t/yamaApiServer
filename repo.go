package main

import "fmt"

var (
  yamas     Yamas
  currentID int
)

func init() {
  RepoCreateYama(Yama{Proposer: "@takachan-mirai"})
  RepoCreateYama(Yama{Proposer: "@Tkon_sec"})
}

func RepoFindYama(id int) Yama {
  for _, t := range yamas {
      if t.ID == id {
          return t
      }
  }
  return Yama{}
}

func RepoCreateYama(t Yama) Yama {
  currentID += 1
  t.ID = currentID
  yamas = append(yamas, t)
  return t
}

func RepoDestroyYama(id int) error {
  for i, t := range yamas {
      if t.ID == id {
          yamas = append(yamas[:i], yamas[i+1:]...)
          return nil
      }
  }

  return fmt.Errorf("Could not find Quest with id of %d to delete", id)
}
