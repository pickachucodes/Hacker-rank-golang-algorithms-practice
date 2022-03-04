func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
    var diagonal1 int32
    var diagonal2 int32
    
    for i:=0;i<len(arr);i++{
        diagonal1 +=arr[i][i]
        diagonal2 += arr[i][len(arr)-i-1]
    }
   m:= diagonal1-diagonal2
   floater := float64(m)
  value := math.Abs(floater)
  return int32(value)
}
