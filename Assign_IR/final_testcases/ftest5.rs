 struct T{
     a : i32,
     b : i32,
     c : i32,
     d : i32 ,
} 

fn main (){

	let x : T =T{a:4,b:5,c:6,d:7};;
	let y : T =T{a:4,b:5,c:6,d:7};;
	printf("a=%d",x.a);
	printf("b=%d",x.b);
	printf("c=%d",x.c);
	printf("d=%d",x.d);
	x.a=100;
	y.a=145;
	printf("New a in y=%d",y.a);
	printf("New a=%d",x.a);
    return 0;
}