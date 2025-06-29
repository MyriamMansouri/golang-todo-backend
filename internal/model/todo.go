package main

type todo struct {
    ID          string  `json:"id"`
    Item        string  `json:"item"`
    Completed   bool    `json:"completed"`
}