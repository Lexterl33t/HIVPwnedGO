package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/fatih/color"
  "bufio"
  "os"
  "encoding/json"
)

//wait

type JsonInfo struct {
  ID string
  Source string
  Title string
  Date string
}

func Ascii() string {
  resp, err := http.Get("http://artii.herokuapp.com/make?text=IHPwned")
  if err != nil {
    fmt.Printf("Error: %s", err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Printf("Error: %s", err)
  }
  return string(body)
}

func Copyright() {
  color.Blue("\t[ Developed By Muhammad'RB ]")
}

func EmailPwned(email string) {
  resp, err := http.Get("https://haveibeenpwned.com/api/pasteaccount/%s", email)
  if err != nil {
    fmt.Printf("Error: %s", err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Printf("Error: %s", err)
  }

  var jsoninfo JsonInfo

  json.Unmarshal([]byte(string(body)), &jsoninfo)
  //now the json is parse in structure in variable jsoninfo pointer

  for i := 0; i < binary.Size(jsoninfo); i++ {
    fmt.Printf("ID: %s", jsoninfo.ID)
    fmt.Printf("POST website: %s", jsoninfo.Source)
    fmt.Printf("Title post: %s", jsoninfo.Title)
    fmt.Printf("Date: %s", jsoninfo.Date)
  }


}


func main() {
  email := bufio.NewScanner(os.Stdin)
  color.HiGreen(Ascii())
  Copyright()

  fmt.Printf("[*] Email: ")
  email.Scan()
  EmailPwned(email.Text())


}
