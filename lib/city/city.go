package city

//import "game/lib/alien"
import "fmt"

type City struct {
  Name string
//  Aliens [100]Alien
  Destroyed bool
  North, South, East, West string
}

func (c *City) Add_route(direction, city string) {
  if(direction == "N") {
    c.North = city
  } else if(direction == "S") {
    c.South = city
  } else if(direction == "W") {
    c.West = city
  } else if(direction == "E") {
    c.East = city
  }
}

func (c *City) Destroy() {
  c.Destroyed = true
}

func (c *City) CanTravel(direction string) bool {
  var neighbor string
  if(direction == "N") {
    neighbor = c.North
  } else if(direction == "S") {
    neighbor = c.South
  } else if(direction == "W") {
    neighbor = c.West
  } else if(direction == "E") {
    neighbor = c.East
  }

  if(neighbor != "" && !(c.Destroyed)){
    return true
  } else {
    return false
  }
}

func (c *City) Get_neighbor(direction string) string {
  var neighbor string
  if(direction == "N") {
    neighbor = c.North
  } else if(direction == "S") {
    neighbor = c.South
  } else if(direction == "W") {
    neighbor = c.West
  } else if(direction == "E") {
    neighbor = c.East
  }

  return neighbor
}

func (c *City) Alien_arrive() {
  fmt.Println("do something")
}

func (c *City) Alien_depart() {
  fmt.Println("do something")
}

func (c *City) Get_aliens() {
  fmt.Println("do something")
}

func (c *City) City_info() string {
  summary := c.Name

  if(c.CanTravel("N")){
    summary += " north=" + c.North
  }

  if(c.CanTravel("S")){
    summary += " south=" + c.South
  }

  if(c.CanTravel("E")){
    summary += " east=" + c.East
  }

  if(c.CanTravel("W")){
    summary += " west=" + c.West
  }

  return summary
}
