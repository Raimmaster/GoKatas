package Calculator

import (
  "testing"
)

func TestAddZeroNumbers(t *testing.T){
   num, _ := Add("")
   expectedValue := 0
   if num != expectedValue {
     t.Fatalf("Received: %d, expected %d \n", num, expectedValue)
   }
}

func TestAddOneNumber(t *testing.T){
  num, _ := Add("55")
  expectedValue := 55

  if num != expectedValue {
    t.Fatalf("Received %d, expected %d \n", num, expectedValue)
  }
}

func TestAddTwoNumbers(t *testing.T){
  num, _ := Add("10,25")
  expectedValue := 35

  if num != expectedValue {
      t.Fatalf("Received %d, expected %d \n", num, expectedValue)
  }
}

func TestUnknownAmountOfNumbers(t *testing.T){
  num, _ := Add("10,25,30,55,66")
  expectedValue := 186

  if num != expectedValue {
      t.Fatalf("Received %d, expected %d \n", num, expectedValue)
	}
}

func TestSupportingDifferentDelimiters(t *testing.T){
	num, _ := Add("//-\n10-25-30-55-66")
  expectedValue := 186

  if num != expectedValue {
      t.Fatalf("Received %d, expected %d \n", num, expectedValue)
  }
}

func TestForNotHavingNegativeNumbers(t *testing.T){
	_, err := Add("10,25,30,55,66")
  
  if err != nil{
  	t.Fatalf(err.Error())
  }
}

func TestNumbersLowerThanThousand(t *testing.T){
  num, _ := Add("1000,55")
  expectedValue := 55
  if num != expectedValue {
    t.Fatalf("Received %d, expected %d \n", num, expectedValue)
  }
}

func TestMultipleLengthDelimiter(t *testing.T){
  num, _ := Add("//[***]\n10***25***30***55***66")
  expectedValue := 186

  if num != expectedValue {
      t.Fatalf("Received %d, expected %d \n", num, expectedValue)
  }
}