package main

import (
	"time"
)

type Login struct {
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	LoginTime time.Time `json:"loginTime"`
	Provider  string    `json:"provider"`
	Roles     []string  `json:"roles"`
	Token     string    `json:"token"`
}
