package user

import (
  "os/user"
  "os"
  "fmt"
)


func handleError(err error) { 
  if err !=nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
func Get_current_hostname() string {
  hostname, err := os.Hostname()
  handleError(err) 
  return hostname
}

func Get_current_user() string {
  
  curUser, err := user.Current()
  handleError(err)
  return curUser.Username
}


func Get_current_dir() string { 
  curdir, err := os.Getwd()
  handleError(err) 
  return curdir
}








