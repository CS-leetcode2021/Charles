use std::io;
use rand::Rng;
use std::cmp::Ordering;
// 引用标准库的io作用域

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..101);
    println!("The secret number is:{}", secret_number);

    loop {
        println!("Please input your guess.");
        // 进行比较，然后选择符合的分支输出

        let mut guess = String::new();  // 创建用户输入的地方
        // let 创建变量
        // eg: let foo = 5; // 不可变
        // eg: let mut foo = 5; // 可变

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");
        // 调用.read_line方法从标准输入句柄获取用户输入。我们还向read_line()传递了一个参数：&mut guess

        // 重新创建一个guess的变量，允许用一个新值来隐藏guess之前的值，常在需要转换值类型之类的场景使用
        // 允许我们复用guess变量的名字，而不是被迫创建两个不同变量
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed:{}", guess);
        // {} 占位符

        match guess.cmp(&secret_number) {
            Ordering::Less => println! {"Too small!"},
            Ordering::Equal => {
                println!("You win!");
                break;
            }
            Ordering::Greater => println! {"Too big!"},
        }
    }
}