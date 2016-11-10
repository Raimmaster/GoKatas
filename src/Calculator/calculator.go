package Calculator

import (
  "strings"
  "strconv"
  "fmt"
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
	var delimiter string
	var isSetDelimiter bool
	if(strings.ContainsAny(numbers, "//")){
		fmt.Println("It contains it!")
		delimiter = numbers[2:3]
		numbers = numbers[4:len(numbers)]
		isSetDelimiter = true
	}

	var adders []string
	if(!isSetDelimiter){
  	adders = strings.FieldsFunc(numbers, Split)
	}else {
		adders = strings.Split(numbers, delimiter)
	}
  
  num = 0

  for _, value := range adders {
    val, erro := strconv.Atoi(value)
      if erro == nil{
        num += val        
      }
  }

  return num
}

func Split(r rune) bool {
	return r == ',' || r == '\n'
}
