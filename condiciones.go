package main
import "fmt"
/*
== igual a
!=diferente de
< menor que
> mayor que
>= mayor o igual que
<= menor o igual que
&& AND
|| OR
*/

func main (){
	x :=20
	y :=20
	if x>y {
		fmt.Println(x ,"es mayor que ",y)
	}else if x <y {
		fmt.Println(x, "es menor a ",y)
	}else{
		fmt.Println("x,y son iguales")
	}


}