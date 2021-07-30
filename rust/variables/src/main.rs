fn main() {
    // 常量使用const
    // const MAX_POINTS: u32 = 100_000;

    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);

    // 标量与复合类型
    // 标量是一个单独的值

    // 标量：整形、浮点型、布尔、字符
    // 复合类型：元组和数组

    // 元组是一个将多个其他类型的值组合进一个复合类型的主要方式。元组长度固定：一旦声明，其长度不会增大或缩小
    // 使用包含在圆括号中的逗号分隔的值列表来创建一个元组。元组中的每一个位置都有一个类型，而且这些不同值的类型也不必是相同的。

    let tup: (i32, f64, u8) = (500, 6.4, 1);
    // 访问元组、解构
    let (x, y, z) = tup;
    println!("The value of y is:{}", y);
    println!("The value of x is:{}", x);
    println!("The value of z is:{}", z);

    // 或者使用索引点
    let five_hundred = tup.0;
    let six_point_four = tup.1;
    let one = tup.2;

    println!("1 is {}", five_hundred);
    println!("2 is {}", six_point_four);
    println!("3 is {}", one);
    // 数组：包含的元素类型必须相同
    let a = [1, 2, 3, 4, 5, 6];


    // let months = ["January", "February", "March", "April", "May", "June", "July",
    //     "August", "September", "October", "November", "December"];

    // 变量名为 a 的数组将包含 5 个元素，这些元素的值最初都将被设置为 3。
    // let b = [3; 5];

    // 使用索引访问
    let first = a[0];
    let second = a[1];

    println!("First is {}", first);
    println!("Second is {}", second);

    // 底层rust会检测数组越界操作
    // let index = 0;
    // let element = a[index]; // 这里运行时会报错

    functions(23, 32);
    let tmp = functions2(3);
    println!("tmp is :{}", tmp);

    branches();
    loop_functions();
    println!("----------------------------");
    while_functions();
    println!("----------------------------");
    for_functions();
}


fn functions(x: i32, y: i32) { // 函数体定义原码前后无区别
    println!("This is another function. value is {} {}", x, y);

    // rust是一门基于表达式的语言，表达式计算并产生一个值
    let a = {
        let b = 1;
        b + 1     // 表达式的结尾没有分号。如果在表达式的结尾加上分号，它就变成了语句，而语句不会返回值。
    };

    println!("a is : {}", a);
}

// 有返回值的函数

fn functions2(x: i32) -> i32 {
    x + 1
}


fn branches() {
    let number = 3;
    if number < 5 {  // 判定值必须是boole类型，不是一个非1的数字
        println!("condition was true");
    } else {
        println!("condition was false");
    }

    // let语句中使用if

    let condition = false;
    let tmp = if condition {    // if 和 else的返回语句类型必须是一样的
        5
    }else {
        //"six" // error
        6
    };

    println!("tmp is {}",tmp);
}

fn loop_functions(){
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {}", result);
}

fn while_functions(){
    let mut number = 3;

    while number != 0 {
        println!("{}!", number);

        number = number - 1;
    }

    println!("LIFTOFF!!!");
}

fn for_functions(){
    let a = [10,20,30,40,50];
    let mut index = 0;
    while index < 5 {
        println!("the value is: {}", a[index]);

        index = index + 1;
    }

    for element in a.iter() {
        println!("the value is:{}",element);
    }

    for number in (1..4).rev() {
        println!("{}!", number);    // 3.2.1
    }
    println!("LIFTOFF!!!");
}

