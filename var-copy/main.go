package main


func addHundred(num *int) {
  *num += 100
}

func main() {
  x := 1
  addHundred(&x) 
}