package main

import "time"

type Yama struct {
  ID        int       `json:"id"`
  Proposer  string    `json:"proposer"`
  Members   []string  `json:"members"`
  Place     string    `json:"place"`
  Completed bool      `json:"completed"`
  Due       time.Time `json:"due"`
}

type Yamas []Yama
