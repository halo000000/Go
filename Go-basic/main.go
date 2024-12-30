package main

import (
	// "errors"
	// "errors"
	"fmt"
	// "image/color"
	// "maps"
	// "sync"
	// "time"
	// "strconv"
	// "strings"
	// "strconv"
	// "os"
	// "ferst-program/the"
	// "ferst-program/math"
)

// func add(a int,b int) int {
// 	return a+b;
// }

// type userData struct{
// 	id int
// 	name string
// 	age int
// 	isAlive bool

// }

// var wg= sync.WaitGroup{

// }



// type data struct{
// 	a int
// 	b int
// }

// func(num *data) add() int{
// 	return num.a * num.b
// }


// type aa struct{
// 	name string
// 	color string
// }

// type bb struct{
// 	aa
// 	id int
// }


// func addN()func(string)string{
// 	a:=""
// 	return func(s string) string {
// 		a+=s
// 		return a
// 	}

// }















func main()  {

	// var name string ="halo";
	// var number int ;
	// fmt.Println("what is this man this is a new line ");
	// fmt.Println("what is this "+name);
	// fmt.Scanln(&number)
	// fmt.Printf("you typed %d",number)


	// var name string ="halo";

	// fmt.Printf("the name  is  %s ",name)

	// name="zhila";

	// fmt.Print("the name is "+name)

	// var is_good bool = false;

	// fmt.Scan(&is_good);


	// if (is_good) {
	// 	fmt.Print("you are good");
	// }else {
	// 	fmt.Print("you are not good");
	// }

	// var a int =12;

	// fmt.Println(a)

	// b:=20;
	// fmt.Println(b)

	// for i := 0; i < 20; i++ {
	// 	fmt.Println("num:",i)
	// }



//  var i int=12;
// for (i>0){
// i--;
// fmt.Println("num:",i);
// }


// for i:=range 10{

// 	fmt.Println(i)
// }


// const wonerName string ="halo";

// fmt.Printf("the owner name is %s",wonerName);

// const num float32=12.99;

// fmt.Printf("the owner name is %.2f",num);

// var n int =12;

// fmt.Printf("num is %dwhat is this ",n);




// fmt.Print("the number is ",add(12,34));

// var Name string = "mangodb";


// fmt.Printf("\nthis is %v a better way of writing ",Name)


// var abc int;

// fmt.Print(abc)


// var names bool ;

// fmt.Scan(&names)

// fmt.Printf("%T",names)

// fmt.Print("\n",names)



// var numbers = [10]int{1,2,3,4,5,6,7,8,9,0}
// var n=[...]int{1,2,3,4,5}

// fmt.Print(len(numbers),n)

// var text = []byte("\nthis is a file write oparation\n ")

// os.WriteFile("./halo.txt",text,0644)
// os.WriteFile("./halo.txt",text,0777)




// dat, err := os.ReadFile("./halo.txt")
// if err!=nil{
// 	fmt.Print(err)
// }
// fmt.Print(string(dat))

// this is how to append to a file 

// file, err := os.OpenFile("./halo.txt",os.O_APPEND,0644)
// if err!=nil{
// 	fmt.Print(err)
// }
// len, err := file.WriteString(" The Go language was conceived in September 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google.")
// if err!=nil{
// 	fmt.Print(err)
// }
// fmt.Print(len)

// POINTER is a memeory address

// var num int =12;

// fmt.Println(num)
// fmt.Println(&num)


// var names= []int{1,2,3,4,5,6}

// names[7]=0;

// fmt.Println(names)




// var a int=1;
// fmt.Printf("%T\n",strconv.Itoa(a))


// var b string="12";
// value,err:=strconv.Atoi(b)
// if err !=nil{
// 	fmt.Print(err)
// }
// fmt.Printf("%T\n",value)




// var names = [4]string{"halo","ali","omar","adam"}



// for index,a:= range names {
// 	fmt.Println(a,index)
// }


// var isRuning bool = true;
// for isRuning {
// fmt.Println("a")

// }


// var numbers = []int{1,2,3,4,5}
// fmt.Println("befor appned")
// for _,i:=range numbers{

// 	fmt.Print(i)
// }

// numbers = append(numbers, 6)

// fmt.Println("\naftar appned")
// for _,i:=range numbers{

// 	fmt.Print(i)
// }


// var fullNmaes = []string{"halo jamal","zhila ahmad","farhang hussen"};


// for _,name := range fullNmaes{
// 	// ferstName,_,_:=strings.Cut(name," ")
// 	// fmt.Println(ferstName)
// 	// fmt.Println(strings.Count(ferstName,"a"))
// 	fmt.Println(strings.Fields(name)[0])
// }


// n,err:= strconv.ParseBool("0")

// if err!=nil{
// 	fmt.Print(err)
// }

// fmt.Printf("n: %v\ntype:%T", n,n)


// for{
// 	var num int;
// 	fmt.Scan(&num)
// 	if num == 0{
// 		fmt.Printf("\nhey you goat it\n")
// 		break
		
// 	}
// 	fmt.Printf("the numbers is not %v",num)
// }



// n:=func (a int,b int)int  {
// 	return a+b;
// }(12,2)

// fmt.Println(n)

fmt.Println("")

// hi.Hi()

// fmt.Println(hi.Make)


// fmt.Printf("math.Add(1, 2, 3, 4, 5, 6): %v\n", math.Add(1, 2, 3, 4, 5, 6))
// fmt.Printf("math.Mines(12, 9, 0): %v\n", math.Mines(12, 9, 0))
// fmt.Printf("math.Divied(10, 2): %v\n", math.Divied(10, 2))
// fmt.Printf("math.Moltiply(12, 2, 4): %v\n", math.Moltiply(12, 2))


// var k = make(map[string]int)


// k["a"]=12
// k["b"]=1

// var names = map[string]string{
// 	"A":"halo",
// 	"b":"zhila",
// }


// var names2=map[string]string{
// 	"A":"halo",
// 	"b":"zhila",

// }



// fmt.Println(names)
// fmt.Println(names2)
// fmt.Println(maps.Equal(names,names2))


// var data = userData{
// 	id:1,
// 	name: "halo",
// 	age: 22,
// 	isAlive: true,


// }



// fmt.Printf("data.age: %v\n", data.age)
// fmt.Printf("data.name: %v\n", data.name)
// fmt.Printf("data.isAlive: %v\n", data.isAlive)


// p :=fmt.Sprintf("the name is %v and it is %v year old ",data.name,data.age)


// fmt.Printf("%v\n", p)

// //  go concerncy 



// fmt.Printf("\nthis is a name of the bigger int")




// wg.Add(1)
// go run_on_difrentTherad()




// fmt.Printf("\nthis code is aftar  the 10s delay ")
// wg.Wait()








// k:=data{a:1,b:2}

// fmt.Printf("k.add(): %v\n", k.add())

// animal:=struct{
// 	name string
// 	color string
// 	speed int


// }{
// 	name: "lion",
// 	color: "orange",
// 	speed: 20,
// }


// fmt.Println(animal)




// al := bb{
// 	id: 1,
// 	aa: aa{
// 		name: "halo",
// 		color: "red",
// 	},
// }



// fmt.Printf("al: %v\n", al)
// // we can dircly acses the embeded type that was aa without caling the aa 
// fmt.Println( al.color)
// fmt.Println( al.name)
// fmt.Println( al.id)





// defer fmt.Println("this will get exacuted before the function ends ")

// a:=errors.New("this is errore")
// fmt.Printf("a: %v\n", a)


// text:=addN()

// fmt.Printf("1st call %v\n", text("a"))
// fmt.Printf("1st call %v\n", text("b"))
// fmt.Printf("1st call %v\n", text("c"))




// pointers 

// x:=1
// y:=2
// A:=x
// B:=&x

// fmt.Printf("x: %v\n", x)
// fmt.Printf("y: %v\n", y)
// fmt.Printf("A: %v\n", A)
// fmt.Printf("B: %v\n", B)
// // we changed the value of x 
// *B=4
// fmt.Printf("x: %v\n", x)
// fmt.Printf("*B: %v\n", *B)

// messages := make(chan string)
// go func ()  {
// 	fmt.Print("this runs in a nother thread ")
// 	messages <- "hi from the 2nd thread"
// }()


// msg:= <-messages
// fmt.Printf("\nmsg: %v\n", msg)

parm := getPathParams("/api/v1/users", 3)

fmt.Println( parm)

}

// func run_on_difrentTherad(){
// 	time.Sleep(10*time.Second)

// 	fmt.Printf( "\naftar 10s ")
// 	wg.Done()
// }
