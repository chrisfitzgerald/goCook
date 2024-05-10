package main

type User struct {
    Username string
    Password string 
}

var users = []User{
    {Username: "admin", Password: "password"},
}