package main

func myTest() *int{
     var a int = 10
     var b = &a
     return b
}

func main(){
     var c = myTest()
     _ = c
}	