package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var (
 active = 0
 files []fs.DirEntry
 path string 
 blackText = color.New(color.FgBlack)
 activeText = blackText.Add(color.BgWhite)
)



func PrintDir(){
	clearScreen()

	fmt.Println(path)

	for i, file := range files{
		if i == active {
			activeText.Println(file.Name())
		} else{
			fmt.Println(file.Name())
		}
	}
}

func ReadDir() {
	DirEntry, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)	
	}
	
	files = DirEntry
}

func clearScreen() {
	c := exec.Command("cmd", "/C", "cls")
	c.Stdout = os.Stdout
	_ = c.Run() 
}

func GetPath(){
	var err error
	path, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

func KeyPressHandler(key keyboard.Key){
	switch key {
	case keyboard.KeyArrowUp :
		if(active > 0){
			active--
		} else {
			active = len(files) - 1
		}
		PrintDir()
	case keyboard.KeyArrowDown:
		if active < len(files) - 1 {
			active++
		} else {
			active = 0
		}
		ReadDir()
		PrintDir()
	case keyboard.KeyArrowLeft:
		path = filepath.Dir(path)
		active = 0
		ReadDir()
		PrintDir()
	case keyboard.KeyEnter, keyboard.KeyArrowRight:
		if files[active].IsDir() {
			path = path + "\\" + files[active].Name()
			active = 0
			ReadDir()
			PrintDir()
		}	
	case keyboard.KeyEsc:
		fmt.Println("End...")
		os.Exit(0)
	}
}

func main() {

	
	GetPath()
	ReadDir()
	PrintDir()

	
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		KeyPressHandler(key)
		
	}	

}