package main

import(
	"fmt"
	"encoding/csv"
	"flag"
	"os"
	"time"
)

/**Format for the problems */
type problem struct{
	ques string
	ans string
}

func main() {
	
	/* Command line arguments */
	fileName := flag.String("csv","problems.csv","File to read the quiz !!")
	timeLimit := flag.Int("timelimit",30,"Time limit for the quiz ( seconds)")

	flag.Parse()

	/*Open the csv file and read all the entries to slic of struct*/
	problems := readFile(*fileName)

	/*Flash the quiz with timer*/
	quiz(problems,*timeLimit)
	
}

/**Conduct the quiz*/
func quiz(problems []problem,timeLimit int){

	timer := time.NewTimer(time.Duration(timeLimit)*time.Second)
	correct := 0
	total := len(problems)
	label :
	for i,x := range problems{
		fmt.Printf("Probblem #%d: %s --> ",i+1,x.ques)
		answerCh := make(chan string)
		go func(){
			var ans string
			fmt.Scanf("%s ",&ans)
			answerCh <- ans
		}()
		select{
			case <-timer.C :
				fmt.Printf("\nTime Up !!!")
				break label

			case ans := <-answerCh :
				if ans == x.ans{
					correct++
				}
		}
	}
	fmt.Printf("\nYou Scored %d out of %d.\n",correct,total)
}

/**Open and read the csv file*/
func readFile(fileName string) []problem{
	file,err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("Error !! Check if the file %s exists and you have permissions to access it !",fileName))
	}
	x := csv.NewReader(file)
	lines,er := x.ReadAll()
	if er != nil{
		exit(fmt.Sprintf("Error !! Couldn't read the file."))
	}
	return refactor(lines)
}
/**Refactor the csv in form of slice of struct*/
func refactor(lines [][]string) []problem{
	
	problems := make([]problem,len(lines))
	for i,x := range lines{
		problems[i].ques = x[0]
		problems[i].ans = x[1]
	}
	return problems
}

/**Generic function for exit*/
func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}