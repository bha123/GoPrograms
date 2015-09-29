package main
import "fmt"
import "math"
func main(){
	// Initializing variables
	var  val int
	k := 0
	l := 0
	numb := 0
	sum := 0
	temp_1 := 0
	// TITLE OF PROGRAM
	fmt.Println("#########################################")
	fmt.Println("             ARMSTRONG NUMBER            ")
	fmt.Println("                YES  OR NO               ")
	fmt.Println("#########################################")
	fmt.Println("Author Bharath.P (nmania03@gmail.com)")
	fmt.Println("Enter the value you want to check as ARMSTRONG NUMBER:")
	fmt.Scan(&val)
	k = val
	l = val	
	// Reading the length of the values 
	// Finding the number of digits in the number
	for l != 0{
			l = l/10
			numb++
	}
	fmt.Println("The Number of digits in the given Input ",numb)
	//Checking if the given number is armstrong or not 
	for k != 0{
		sum = k%10
		k   = k/10
		temp_1 = temp_1 + int(math.Pow(float64(sum), float64(numb)))	
	}
	if temp_1 ==int(val){
			fmt.Println("The Given Number IS  Armstrong Number")
	} else {
		fmt.Println("The Given Number is NOT an Armstrong Number")
	}
	
}

	