package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Characters struct{
	Name string `json:"name"`
	MaxPower int `json:"max_power"`
}

type Avengers struct{
	Name string `json:"name"`
	Character []Characters
}

  
type netList struct{
  Name string `json:"name"`
  MaxPower int `json:"max_power"`
  Count int
}

func sorting(finalList []netList){
  sort.Slice(finalList, func(i, j int) bool {
    if finalList[i].Count > finalList[j].Count {
        return true
    }
    if finalList[i].Count < finalList[j].Count {
        return false
    }
    return finalList[i].MaxPower > finalList[j].MaxPower
})
}

func main() {

  var finalList []netList

  urls := []string{
    "http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b",
    "http://www.mocky.io/v2/5ecfd630320000f1aee3d64d",
    "http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e",
  }
  for i:=0; i<len(urls);i++ {
    response, err := http.Get(urls[i])
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Avengers
    json.Unmarshal(responseData, &responseObject)

    for i := 0; i < len(responseObject.Character); i++ {
      results := []netList{{Name: responseObject.Character[i].Name, MaxPower: responseObject.Character[i].MaxPower}}
      finalList=append(finalList, results...)
    }
  }

  fmt.Println(finalList)

  //Searching in the list
  for j:=0;j<5;j++{
  var flag=0
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter name of the caracter: ")
  name, _ := consoleReader.ReadString('\n')

  for i:=0;i<len(finalList);i++{
      if strings.TrimRight(name, "\r\n")==strings.ToLower(finalList[i].Name){
          fmt.Println("Max Power: ")
          finalList[i].Count+=1
          fmt.Println(finalList[i].MaxPower)
          flag=1
          break;
      }
  }
  if flag==0{
      fmt.Println("Character is not there in the list!!")
  }
}
fmt.Println(finalList)
sorting(finalList)

//Removing extra elements
if(len(finalList)>10){
    for len(finalList)!=10 {
        finalList=finalList[:len(finalList)-1]
    }
}
fmt.Println(finalList)

//Adding in the list
  var powerValue int
  var nameValue string
  fmt.Println("Enter new character name and it's max power: ")
  fmt.Scanln(&nameValue)
  fmt.Scanln(&powerValue)
  finalList=finalList[:len(finalList)-1]
  results := []netList{{Name: nameValue, MaxPower: powerValue}}
  finalList = append(finalList,results...)
  sorting(finalList)
  fmt.Println(finalList)
}

