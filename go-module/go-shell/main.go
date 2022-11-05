package main
import (
  "veasna.com/user"
  "fmt"
)
func main() { 
  fmt.Println("hello, World")
  fmt.Println(user.Get_current_user(), user.Get_current_hostname(), user.Get_current_dir())

}
