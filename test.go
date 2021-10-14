package main

import (
	"fmt"
	// "os"
	// "bufio"
	"io/ioutil"
	// "math/rand"
	"time"
	// "os/exec"
	"strconv"
	// "runtime"
	// "sync"
	// "math/big"
	// "bytes"
	// "strings"
	"reflect"
	// "syscall"
)

func main(){
	// t1:=time.Now()
	// files, _ := ioutil.ReadDir("./data")
	// fmt.Println(time.Now().Sub(t1))

	t1:=time.Now()
	for i:=0;i<128;i++{
		// 2m27.964033412s
		// fd,_:=ioutil.ReadFile("./data/Treasure_"+strconv.Itoa(i)+".data")
		// fmt.Println(len(fd))
		// fmt.Println(reflect.TypeOf(&fd))

		// 3m40.929148611s
		// f, _ := os.Open("./data/Treasure_"+strconv.Itoa(i)+".data")
		// defer f.Close()
		// fd, _ := ioutil.ReadAll(f)
		// fmt.Println(len(fd))

		//3
		//3.192445ms
		f,_ := os.Open("./data/Treasure_"+strconv.Itoa(i)+".data")
		read:=bufio.NewReader(f)
		fmt.Println(reflect.TypeOf(read))
		fmt.Println(i)



	}
	fmt.Println(time.Now().Sub(t1))

	t1=time.Now()
	// for _,_= range files{
	// 	// ioutil.ReadFile("./data/"+fi.Name())

	// 	// fmt.Println(reflect.TypeOf(file1))
	// 	// fmt.Println("./data/"+fi.Name())
	// }
	fmt.Println(time.Now().Sub(t1))
}