package main
import (
	"os"
	"os/exec"
	"log"
)
func main(){
	languages  := [12]string{"c","c++","python","dlang","julia","scala","golang","javascript","dart","rust","R","java"};
	extensions := [12]string{".c",".cpp",".py",".d",".jl",".scala",".go",".js",".dart",".rs",".r",".java"};
	for _,arg := range os.Args[1:]{
		execute("mkdir",arg)
		pathOuter := pwd()
		os.Chdir(arg)
		pathInner := pwd()
		for i,lang := range languages{
			execute("mkdir",lang)
			os.Chdir(lang)
			execute("touch",string(arg+extensions[i]))
			os.Chdir(pathInner)
		}
		os.Chdir(pathOuter)	
	}	
}
func execute(param1 string, param2 string){
	execCmd := exec.Command(param1,param2)
	err     := execCmd.Run()
	check(err)
}
func check (e error){
	if e != nil{
		log.Fatal(e)
	}
}
func pwd() string{
	path, err := os.Getwd()
	check(err)
	return path
}	
	
	
	
	
	
	
	
	
	
	
	
		
		
