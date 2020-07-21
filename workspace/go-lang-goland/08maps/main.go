package main

import "fmt"

func main() {
	states := make(map[string] string)
	fmt.Println(states)
	states["HA"]="DF"
	states["AB"]="CD"
	states["EF"]="YZ"
	for k,_:= range states{
		fmt.Println(k)
	}
	delete(states,"EF")
	for k,_:= range states{
		fmt.Println(k)
	}
	i:=0
	keys:= make([]string,len(states))
	for k:=range states{
		keys[i]=k
		i++
	}
	for i:=range keys{
		fmt.Println(keys[i])
	}

}
