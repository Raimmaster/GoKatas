package Calculator

import (
  "strings"
  "strconv"
)

func Add(numbers string) int {
  // adders := strings.Split(numbers, ",")
  if strings.Compare(numbers, "") == 0 {
    return 0
  }

  num, err := strconv.Atoi(numbers)
  if err != nil {
    if num >= 0 {
      return num
    }
  }

  return 0
}
