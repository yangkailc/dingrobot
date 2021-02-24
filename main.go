package main


import (
       "github.com/yangkailc/dingrobot/pkg/util"
       "fmt"
)
func  main(){
      fmt.Printf("%t\n", util.CompareTimestamp(1614158858))
      fmt.Printf("%t\n", util.CompareSign("1", "1", "1"))
      fmt.Printf("%t\n", util.CompareTimestamp(1614158858) && util.CompareSign("1", "1", "1"))
}
