package main 
import "fmt"
func main(){
	var arr []int=[]int{1,2,3,4,5,6,7,8}
	ch:=make(chan int)
	n:=3
	chunk:=0
	if n%3==0{
        chunk=int(n/3)
	}else{
		chunk=int(n/3)+1
	}
	for i:=0;i<n;i++{
		start:=i*chunk
		end:=i+chunk
		if(end>len(arr)){
			end=len(arr)
		}
		go func(brr []int){
		   sum:=0;
		   for _,V:=range brr{
			sum+=V
		   }
		   ch<-sum
		}(arr[start:end])
	}
	total:=0
	for i:=0;i<n;i++{
      total+=<-ch
	}
	fmt.Println(total)
}