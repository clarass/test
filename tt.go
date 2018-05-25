package main
import (
  "log"
)

func main(){
  x := 3
  defer log.Println(x)
  defer func(){
    log.Println(x)
  }()
  //defer log.Println("ss",recover())
  //defer log.Println(recover())
  x = 5
  s := "ab"+
"cd"
  println(s)

}
