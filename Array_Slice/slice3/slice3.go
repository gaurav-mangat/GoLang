package main
import "fmt"

func main(){
	sl:=[]string {"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)"}
	for _,v:=range sl{
		fmt.Printf("%v\n",v)
	}
	fmt.Println("----------------------------------------------")
  b:=len(sl)
  for k:=0;k<b;k++{
	fmt.Println(sl[k])
  }

}

