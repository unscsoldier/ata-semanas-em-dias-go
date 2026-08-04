package main

import (
  "fmt"  
)

func main() {
  var semanas int;
  var dias int;  
  fmt.Scan(&semanas);
  
  dias = semanas * 7;
  
  fmt.Println(dias);
}
