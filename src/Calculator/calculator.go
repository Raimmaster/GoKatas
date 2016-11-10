package Calculator

import (
  "strings"
  "strconv"
  "errors"
  "fmt"
)
var delimiters []rune = []rune{',', '\n'}

func Add(numbers string) (int, error) {
	//delimiters 
  if strings.Compare(numbers, "") == 0 {
    return 0, nil
  }

  num, err := strconv.Atoi(numbers)
  if err == nil {
    if num >= 0 {
      return num, nil
    }
  }

	var delimiter string
	var isSetDelimiter bool
	if(strings.ContainsAny(numbers, "//")){
		delimiter = numbers[2:3]
		numbers = numbers[4:len(numbers)]
		isSetDelimiter = true
	}

	var adders []string
	if(!isSetDelimiter){
		for _, val := range delimiters {
			negativeSubstring := string(val) + "-"
			fmt.Printf("Sub: %s\n", negativeSubstring)
			if(strings.Contains(numbers, negativeSubstring)){
				return -1, errors.New("Contains a negative value!")
			}
		}
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

  return num, nil
}

func Split(r rune) bool {
	isDelimiter := false
	for _, value := range delimiters {
		if(r == value){
			isDelimiter = true
			break
		}
	}

	return isDelimiter
}
