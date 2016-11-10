package Calculator

import (
  "strings"
  "strconv"
)

func Add(numbers string) int {
  if strings.Compare(numbers, "") == 0 {
    return 0
  }

  num, err := strconv.Atoi(numbers)
  if err == nil {
    if num >= 0 {
      return num
    }
  }
  
  adders := strings.Split(numbers, "\n")
  
  num = 0

  for _, value := range adders {
    val, erro := strconv.Atoi(value)
      if erro == nil{
        num += val        
      }
  }

  return num
}
