package ini_loader
import "fmt"
import "os"
func assert(b bool) {
	if b {
		return
	}
	fmt.Print("Error in test")
	os.Exit(0)
}

func main(){
	m := map[string]map[string]string{}
	m["Section-A"] = map[string]string{}
	m["Section-B"] = map[string]string{}
	m["Section-A"]["var"]="true"
	m["Section-A"]["foo"]="bar"
	m["Section-B"]["greeting"]="Hello World"
	dic := Read_ini("example.ini")
	assert(m["Section-B"]["greeting"] == dic["Section-B"]["greeting"])
	assert(m["Section-A"]["foo"] == dic["Section-A"]["foo"])
	assert(m["Section-A"]["var"] == dic["Section-A"]["var"])
	
	fmt.Print("Test Successful")
}
