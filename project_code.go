package main

import (
	"fmt"	
  "os"
  "time"
  "os/exec"
)

var(  
	answer, age string
	anthem int
	//age string
	bones int
	choice string 
	covid string
	c0py string
	coin int
	cartier int
	CFtemp int
	email string
	fabric string
	Fname string
	gas string
	geography string
	grade int
	iron string
	invasion int
	justin string
	jobs int
	plusone int
	province string
	pi float32
	physics string
	Sname string
	speed string
	small string
	sharks string
	teeth int
	num int
	Nitrogen int
	holocaust string
	//For Keeping a count to how many questions the user has got correct while doing the quiz
	questionscorrect int = 0; 
)

const (
	GREEN = "\033[32m"
	YELLOW = "\033[33m"
	GREY = "\033[37m"
	BLUE = "\033[34m"
	RED = "\033[31m"
	BLACK = "\033[90m"
	MAGENTA = "\033[35m"
	CYAN = "\033[36m"
	RESET = "\033[0m"
	CLEAR = "\033[H\033[2]"
)


func main() {
	fmt.Println(CLEAR)
	t := time.Now()
		fmt.Printf(GREEN+">>>>>>>>>>>>>>>>>>>>>>\n")
		y := t.Year()
		mon := t.Month()
		d := t.Day()
		h := t.Hour()-5-1
	  //To change the hour again use (-5-1)
	  //Or use (+6)
		m := t.Minute()
		s := t.Second()
		w := t.Weekday()
		 
		fmt.Println(YELLOW+"The Year is:",y)
		fmt.Println("The Month is:",mon)
		fmt.Println("The Day of the month is:",d)
		fmt.Println("Hour   :",h)
		fmt.Println("Minute :",m)
		fmt.Println("Second :",s)
		fmt.Println("The day of the week is:",w) 
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>\n")
	  
	fmt.Println(CYAN+"WELCOME TO THE QUIZ GAME")
	  
	  fmt.Printf(GREY+"What is your first and last name?\n ")
	  fmt.Scan(&Fname)
	  fmt.Scan(&Sname)
	  fmt.Printf("Hello %v %v welcome to the quizing game!\n",Fname, Sname)
	  if len(Fname) <= 7 && len(Sname) >= 7 {
		fmt.Println("The MD5# is starting..........")
		duration, _ := time.ParseDuration("6s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run() 
	  }else{
		os.Exit(1)
	  }
	  
	  //Asks the user their grade to determine if they are eligible to play the game.
	  fmt.Println("What grade are you in?:")
	  fmt.Scan(&grade)
	  //Allows the user to continue playing if they are in the eligible grade
	  if grade == 12{
		fmt.Println("You are good to go")  
	  }else if grade == 11{
		fmt.Println("You are good to go")
	  }else if grade == 10{
	   fmt.Println("You are good to go")  
	  }else if grade == 9{
		fmt.Println("You are good to go")  
	  }else if grade < 9{
		fmt.Println("You are still too young for this.")
		return
		os.Exit(1)
	  }else if grade > 12{
		fmt.Println(RED+"Invalid grade")
		os.Exit(0)
	  }
	//Choice for choosing which subjects the user would like to do the quiz on
	fmt.Println("What subject would you like to do the quiz on:?")
	fmt.Println("1-Math")
	fmt.Println("2-Science")
	fmt.Println("3-History")
	fmt.Println("4-Anything")
	fmt.Println("5-Quit")
	fmt.Scanln(&choice)
	
	//An If statement for the user to choose what they want
	if choice == "1"{
	  Math()
	} else if choice == "2"{
	  Science()
	}else if choice == "3"{
	   History()
	}else if choice == "4"{
	  Anything()
	}else if choice == "5"{
	  fmt.Println(RED+"GOODBYE")
	  os.Exit(0)
		}
	}
	func Math(){
	  fmt.Println("Question 1:","What is the length of the hypotenuse(c) if leg a equals to 6cm and leg be is 3cm?")
	fmt.Scan(&answer)
	  if answer == "6.7cm"{
		questionscorrect += 1 
		fmt.Println("Correct")  
	  }else {
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 2:","Justin goes to the mall with $120. He spends $54.67 on clothes, $13.49 on school supplies, and $8.14 on lunch. How much does he have left?")
	  fmt.Scan(&justin)
	  if justin == "$43.70"{
		questionscorrect += 1
		fmt.Println("Correct")  
	  }else {
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 3:","The price of fabric is $7.60 per meter. Lance bought 5.5 meters of fabric. What is the cost of the fabric lance bought?")
	  fmt.Scan(&fabric)
	  if fabric == "$41.80"{
		fmt.Println("You are correct.")
		questionscorrect += 1
	  }else {
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 4:","A gallon of gas costs $2.26. Rob buys 13.5 gallons of gas. How much did he pay?")
	  fmt.Scan(&gas)
	  if gas == "$29.16"{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect\n.")
	  }
	
	 fmt.Println(RED+"Wait for next questions") 
	  fmt.Println(CYAN+"It's not the end, the questions still continue.....")
	duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	 
	fmt.Println(GREY+"Question 6:","What is 25%% of 200?")
	  fmt.Scan(&num)
	  if num == 50{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
	  }
	 fmt.Println("Question 7:","If a driver has travelled 180 miles and it took them 3 hours to make that distance, what  is the speed the car is travelling at:")
	  fmt.Scan(&speed)
	  if speed == "60mph"{
		fmt.Println("Correct")
		questionscorrect += 1
		///for clearing everything on the screen and then printing what is below. 
	  }else{
		fmt.Println("Incorrect")
	  }   
	  fmt.Println("Question 8:","The price of fabric is $7.60 per meter. Lance bought 5.5 meters of fabric. What is the cost of the fabric lance bought?")
	  fmt.Scan(&fabric)
	  if fabric == "$41.80"{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 9: What is 20 + 45 + 156 equal to?")
	  fmt.Scan(&plusone)
	  if plusone == 221 {
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
	  }
	//Stores the points of questions the user has gotten correct out of 9.
	numqu := 9
	  fmt.Printf("You scored %v out of %v", questionscorrect,numqu) 
	  percent := (float32(questionscorrect) / float32(numqu))*100
	  fmt.Printf("\nYou have %.0f%%\n", percent)
	  
	//Prints out a message in accordance with the percenatge the user has accumulated from the short quiz so far.
	  if percent < 30{
		fmt.Println("You can do better, don't give up")
	  }else if percent > 50{
		fmt.Println("Good job trying hard")
	  }else if percent > 75{
		fmt.Println("Excellent")
	  }else if percent == 100{
		fmt.Println("Excelent")
	  }  
	fmt.Println(RED+"The End")
	}
	func Science(){
	fmt.Println("Question 1: At what temperature are Celsius and Fahrenheit equal?")
	fmt.Scan(&CFtemp)
	  if CFtemp == -10{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Fun fact")
	fmt.Println("Question 2: How many bones are in the human body?")
	  fmt.Scanln(&bones)
	  if bones == 206{
		fmt.Println("Correct")
		questionscorrect += 1
		fmt.Println("Did you know: Bones are essential when it comes to balance and coordination.")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("Fun Fact: There are typically around 270 bones in human infants, which fuse to become 206 to 213 bones in the human adult.")
	  
	  }
	fmt.Println("Question 3: The concept of gravity was discovered by which famous physicist?")
	  fmt.Scan(&physics)
	  if physics == "Sir Isaac Newton"{
		fmt.Println("Correct")
	  }else if physics == "Isaac Newton"{
		questionscorrect += 1
		fmt.Println("Correct")
	  }else{
		fmt.Println("Incorrect")
	  }
	fmt.Println("Question 4: What is the chemical symbol of iron?")
	  fmt.Scan(&iron)
	  if iron == "Fe"{
		questionscorrect += 1
		fmt.Println("Correct")
	  }else{
		fmt.Println("Incorrect")
	
	fmt.Println(RED+"Wait for next questions") 
	  fmt.Println(CYAN+"It's not the end, the questions still continue.....")
	duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		
	  }
	fmt.Println("Question 5: Are sharks mammals?")
	fmt.Scan(&sharks)
	  if sharks == "no"{
		questionscorrect += 1
		fmt.Println("correct")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("")
	  }
	fmt.Println("Question 6: What percentage of nitrogen is in our atmosphere?")
	  fmt.Scan(&Nitrogen)
	  if Nitrogen == 78{
		questionscorrect += 1
		fmt.Print("Correct")
	  }else{
		fmt.Println("Incorrect")
	  }
	fmt.Println("Question 7: How many teeth does an adult human have?")
	  fmt.Scan(&teeth)
	  if teeth == 32{
		questionscorrect += 1
		fmt.Println("Correct")
	  }else{
		fmt.Println("Incorrect")
	  } 
	  //This is the total number of questions on the section for science
	c:= 7
	  fmt.Printf("You scored %v out of %v", questionscorrect,c) 
	  percent := (float32(questionscorrect) / float32(c))*100
	  fmt.Printf("\nYou have %.0f%%\n", percent)  
	if percent < 20{
		fmt.Println("You can do better, don't give up")
	  }else if percent > 50{
		fmt.Println("Good job trying hard")
	  } else if percent > 100{
		fmt.Println("Excellent")
	  }
	  fmt.Println("The End")
	  
	}
	func History(){
	  fmt.Println("Question 1:","Is Canada a continent?")
	  fmt.Println("yes or no")
	  fmt.Scan(&geography)
	  if geography == "no"{
		fmt.Println("Correct")
		questionscorrect += 1
		fmt.Println("Note: ")
	  }else if geography == "yes"{
		fmt.Println("Incorrect.")
		fmt.Println("Fun fact: Asia, Africa, North America, South America, Antarctica, Europe, and Australia are continents.")
	}
	  fmt.Println("Questions 2: How many trips did Jacques Cartier make to canada?. Write a number please!")
	  fmt.Scan(&cartier)
	  if cartier == 3{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("Fun fact: He made three voyages to Canada, Cartier became the first European to explore the St. Lawrence Gulf and St. Lawrence River.")
	  }
	  fmt.Println("Questions 3: Who is responsible for the Holocaust?")
	  fmt.Scan(&holocaust)
	  if holocaust == "Adolf Hitler"{
		fmt.Println("correct")
		questionscorrect += 1
		fmt.Println("Did you know: His supremacy fueled the murder of some 6 million Jews")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("Fun Fact:Adolf Hitler, the leader of Germany’s Nazi Party, was one of the most powerful and notorious dictators of the 20th century.")
	  }
	  
	  fmt.Println("Questions 4: Which province joined Canada in 1949?")
	  fmt.Scan(&province)
	  if province == "NewFoundLand"{
		questionscorrect += 1
		fmt.Println("Correct")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("Fun Fact: In 1864, Newfoundland delegates attended the Quebec Conference and signed the resolutions which became of foundation of the 1867 British North America Act. But it was not until over 80 years later, in 1949, that Newfoundland became a Canadian province.")
	fmt.Println("Clearing Screen in 6s....")
		// exit after 6 seconds
		duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run() 
	  }
	  fmt.Println("Questions 5: In what year did the United States invade Canada?")
	  fmt.Scan(&invasion)
	  if invasion == 1812{
		questionscorrect += 1
		fmt.Println("Correct")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println(RED+"Wait for next questions") 
	  fmt.Println(CYAN+"It's not the end, the questions still continue.....")
	duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	  }
	  fmt.Println("Questions 6: In which year did Canada replace 1 dollar bills with coins?")
	  fmt.Scan(&coin)
	  if coin == 1989{
		questionscorrect += 1
		fmt.Println("Correct")
		fmt.Println("Fun Fact: Coins last longer than plain bank notes.")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("Fun Fact: The $1 and the $2 notes stopped being issued in 1989 and 1996, respectively, and were replaced with coins.")
	  }
	  fmt.Println("Questions 7: O Canada was officially adopted as the anthem in which year?")
	  fmt.Scan(&anthem)
	  if anthem == 1980{
		questionscorrect += 1
		fmt.Println("Correct")
	  }else{
		fmt.Println("Incorrect")
		fmt.Println("Fun fact: The music for O Canada was composed in 1980 by Calixa Lavallée — a well-known composer at the time — and the French lyrics were written by Sir Adolphe-Basile Routhier.")
	  }
	c:= 7
	  fmt.Printf("You scored %v out of %v", questionscorrect,c) 
	  percent := (float32(questionscorrect) / float32(c))*100
	  fmt.Printf("\nYou have %.0f%%\n", percent)
	if percent < 20{
		fmt.Println("You can do better, don't give up")
	  }else if percent > 50{
		fmt.Println("Good job trying hard")
	  } else if percent > 100{
		fmt.Println("Excellent")
	  }
	  fmt.Println(RED+"The End")
	}
	func Anything(){
	  fmt.Println("Question 1: In what year did the Covid pandemic begin?")
	  fmt.Scan(&covid)
	  if covid == "2019"{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 2: What goes up but never comes down?")
	  fmt.Scan(&age)
	  if age == "age"{
		fmt.Println("Correct")
		questionscorrect+= 1
	  }else{
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 3:What year was the very first model of the iPhone released?")
	  fmt.Scan(&jobs)
	if jobs == 2007{
	  fmt.Println("Correct")
	  questionscorrect += 1
	}else{
	  fmt.Println("Incorrect")
	}
	  fmt.Println("Question 4:What is often seen as the smallest unit of memory?")
	  fmt.Scan(&small)
	  if small == "Kilobyte"{
		fmt.Println("Correct")
		questionscorrect += 1
	  }else if small == "kilobyte"{
		fmt.Println("correct")
		questionscorrect += 1
	  }else{
		fmt.Println("Incorrect")
	  }
	  fmt.Println("Question 5:What’s the shortcut for the “copy” function on most computers?")
	  fmt.Scan(&c0py)
	  if c0py == "ctrl+c"{
		fmt.Println("Correct")
		questionscorrect += 1
		fmt.Println(RED+"Wait for next questions") 
	  fmt.Println(CYAN+"It's not the end, the questions still continue.....")
	duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	  }else{
		fmt.Println("Incorrect")
	  }  
	fmt.Println(RED+"Wait for next questions") 
	  fmt.Println(CYAN+"It's not the end, the questions still continue.....")
	duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
		
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	  fmt.Println("Question 6: Which email service is owned by Microsoft?")
	  fmt.Scan(&email)
	  if email == "Hotmail"{
	  fmt.Println("Correct")
	  questionscorrect += 1
	}else{
	 fmt.Println("Incorrect")
	}
	  fmt.Println("Question 7: What is the number of pi to two decimal place?")
	  fmt.Scanln(&pi)
	   if pi == 3.14{
		 questionscorrect += 1
		 fmt.Print("Correct")
		 fmt.Println("Fun Fact: We've used computers to calculate pi to more than 22 trillion digits.")
	  }else{
		 fmt.Println("Incorrect")
}
//Stores the points of questions the user has gotten correct.
anything := 7
//if the user is in the eligible grade the program continues them and they get their first question
  //Prints out the percentage to a whole number wihout the decimal point by rounding it.
  fmt.Printf("You scored %v out of %v", questionscorrect,anything) 
  percent := (float32(questionscorrect) / float32(anything))*100
  fmt.Printf("\nYou have %.0f%%\n", percent)
//Prints out a message in accordance with the percenatge the user has accumulated from the short quiz
  if percent < 20{
    fmt.Println("You can do better, don't give up")
  }else if percent > 50{
    fmt.Println("Good job trying hard")
  } else if percent > 100{
    fmt.Println("Excellent")
  }
  fmt.Println(RED+"The End")
 }
