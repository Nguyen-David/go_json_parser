package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "github.com/gorilla/mux"
    "net/http"
)

// User struct which contains a name and email
type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}


func index(w http.ResponseWriter, r *http.Request) {
    // Open our jsonFile
    jsonFile, err := os.Open("users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var users []User

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &users)

    res, err := json.Marshal(users)
    if err != nil {
      panic(err)
    }

    // output as json file
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}


func handleFunc() {
  rtr := mux.NewRouter()
  rtr.HandleFunc("/users", index).Methods("GET")

  http.Handle("/", rtr)

  http.ListenAndServe(":8082", nil)
}

func main()  {
  handleFunc()
}
