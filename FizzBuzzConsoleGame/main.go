package main

import(
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
) 

func main() {
  fmt.Println("I've created simple FizzBuzz game as my first contact with GO.\n")
  fmt.Println("You count from 1 to n. When number is divided by 3 you replace it with Fizz, when is divided by 5 you replace it by Buzz, when is divided by both you replace it by FizzBuzz. Example: 1,2,Fizz,4,Buzz,Fizz,7..\n")
  
  //set to true if game is lost
  isLost:=false
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter n number (game length): ")
  input, _ := reader.ReadString('\n')
  input = strings.TrimSuffix(input, "\n")
  gameLength, err := strconv.Atoi(input)
  
  for err!=nil{
    fmt.Print("Incorect input, try again. Enter n number (game length): ")
    input, _ = reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    gameLength, err = strconv.Atoi(input)
  }
  
  for i:=1; i<=gameLength;i++{
    game:=""
    if i%3==0 {game+="Fizz"}
    if i%5==0 {game+="Buzz"}
    if game=="" {game= strconv.Itoa(i)}
    fmt.Print("To exit press q. Enter game value: ")
    input, _ = reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
  
    if input=="q"{
      fmt.Print("Exit..")
      break
    }

    if game==input{
      fmt.Print("Good job! ")
    }else{
      fmt.Print("You've lost, sorry..")
      isLost=true
      break
    }
}

if input!="q" && isLost==false{
    fmt.Print("Game has ended. You won.")
  }
}
