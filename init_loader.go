/*///////////////////////////////////////////////////////////////////////////////////
//This is a simple init file loader I wrote in golang to read some settings for
//a project i've been working on. Its meant to be simple to understand and read.
//Feedback is appreciated but I don't promise that it will be responded to.
//One feature thats missing is to be able to write out a config given a dictionary
//in other words the inverse.
//
//https://github.com/Angelsan720/init_loader/
//
///////////////////////////////////////////////////////////////////////////////////*/


package init_loader

import "io/ioutil"
import "log"
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func iniToDic(content string)map[string]map[string]string{
	m := map[string]map[string]string{}
	var start int
	var finish int
	var key string
	var value string
	var s string
	for i:=0 ;i < len(content) ;i++{
		if (content[i] == '['){
			start:=i+1
			for i:=i ; i < len(content) ; i++{
				if content[i]==']'{
					finish=i
					break
				}
			}
			s = content[start:finish]
			if _ , ok := m[s]; !ok{
				m[s] = map[string]string{}
			}
		}
		if content[i]=='='{
			for j:=i ; content[j]!='\n'; j--{
				start = j
			}
			key=content[start:i]
			for j:=i ; content[j]!='\n'; j++{
                                finish = j
                        }

                        value=content[i+1:finish+1]
		m[s][key] = value
		}
	}
	return m
}
func Read_ini(s string) map[string]map[string]string {
        content, err := ioutil.ReadFile(s)

        check(err)

        if err != nil {
                log.Fatal(err)
        }

        return iniToDic(string(content))
}
