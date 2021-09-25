package main

import "fmt"
import "strconv"

func main(){
  var kata1 string = "12"
  var kata2 string = "6"
  var a, err1 = strconv.Atoi(kata1)
  var b, err2 = strconv.Atoi(kata2)
  if err1==nil{
    fmt.Println("Penjumlahan : ",a,"+",b,"=",a+b)
    fmt.Println("Pengurangan : ",a,"-",b,"=",a-b)
    fmt.Println("Perkalian : ",a,"*",b,"=",a*b)
    fmt.Println("Pembagian : ",a,"/",b,"=",a/b)
    // fmt.Println(a)
  }
  if err2==nil{
    // fmt.Println(b)
  }

}
