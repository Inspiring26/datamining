package main
import (
	"fmt"
	"os"
	"bufio"
	// "io/ioutil"
	// "math/rand"
	"time"
	// "os/exec"
	"strconv"
	"runtime"
	// "sync"
)

var (
	cstZone = time.FixedZone("CST", 8*3600)
	cpunum = runtime.NumCPU()-1
)


func main(){
	ReadFile()
}

func ReadFile(){
	// var b bytes.Buffer
	file,err := os.Open("./data/Treasure_0.data")
	if err!=nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line[15:79])
		tmpmagic := line[90:len(line)-2]
		magic,_ := strconv.ParseInt(tmpmagic,10,128)
		fmt.Println(magic)
		// 后续用缓冲写
		stingInt:=""
		for _,s := range line[15:79]{
			// 58是:号
			if s<58{
				stingInt+=string(s)
			}
		}
		intData,_ := strconv.Atoi(stingInt)
		fmt.Println(intData)
		break
	}
	file.Close()
}