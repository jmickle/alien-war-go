package main

//import "game/lib/alien"
import (
    "fmt"
    "bufio"
    "os"
    "log"
)

func main() {
  file, err := os.Open("data/test.txt")

  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  //numAliens := 5
  // a := alien.Alien {
  //     Name: "test",
  //     Alive: true,
  //     Trapped: false,
  //     Steps: 0,
  //     Curcity: "test",
  //   }
  //
  //   fmt.Println(a.Destinations())
  //   fmt.Println(a.Destinations())
  //   fmt.Println(a.Destinations())
}
