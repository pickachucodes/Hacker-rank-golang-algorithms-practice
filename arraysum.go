package main
import ("fmt")
var n,t int
func main(){
    fmt.Scan(&n)
    var sum int
    for i:=0;i<n;i++{
        fmt.Scan(&t)
        sum += t
    }
    fmt.Println(sum)
}
