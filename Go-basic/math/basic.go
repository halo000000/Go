package math

func Add(a...int)int{
	var num int;
	for _,i:=range a{
		num+=i
	}

	return num
}

func Mines(a...int)int{
	var num int;
	for _,i:=range a{
		num-=i
	}

	return num
}

func Moltiply(a int,b int)int{
	var num=a*b
	return num
}
func Divied(a int,b int)int{
	var num=a/b
	return num
}


