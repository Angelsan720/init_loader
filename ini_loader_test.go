package ini_loader
import "os"
import "fmt"
func assert(b bool) {
	if b {
		return
	}
	fmt.Print("Error in test")
	Exit(0)
}

func main(){
	m := map[string]map[string]string{}
	m["Section-A"] = map[string]string{}
	m["Section-B"] = map[string]string{}
	m["Section-A"]["var"]="true"
	m["Section-A"]["foo"]="bar"
	m["Section-B"]["greeting"]="Hello World"
	assert(m == Read_ini("example.ini"))
	m["Section-C"] = map[string]string{}
	assert(m != Read_ini("example.ini")
	fmt.Print("Test Successful")

}
