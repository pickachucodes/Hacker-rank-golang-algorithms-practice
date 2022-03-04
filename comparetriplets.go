package main
import ("fmt")
func main(){
    var alice [3]int
    var bob [3]int
    var a int
    var b int
    fmt.Scan(&alice[0],&alice[1],&alice[2])
    fmt.Scan(&bob[0],&bob[1],&bob[2])
    for i:=0;i<3;i++{
        if alice[i] > bob[i] {
            a++
        }else if bob[i]>alice[i] {
            b++
            
        }
        
    }
    fmt.Println(a,b)
}
