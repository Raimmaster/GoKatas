package Calculator

import (
  "testing"
)

func TestAddZeroNumbers(t *testing.T){
   var num int = Add("")
   expectedValue := 0
   if num != expectedValue {
     t.Fatalf("Received: %d, expected %d \n", num, expectedValue)
   }
}

func TestAddOneNumber(t *testing.T){
  var num int = Add("55")
  expectedValue := 55

  if num != expectedValue {
    t.Fatalf("Received %d, expected %d \n", num, expectedValue)
  }
}
