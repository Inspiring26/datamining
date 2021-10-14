package main
import (
	"fmt"
	// "os"
	// "bufio"
	"io/ioutil"
	// "math/rand"
	// "time"
	// "os/exec"
	"strconv"
	"runtime"
	"sync"
	"math/big"
	"bytes"
	"strings"
	// "reflect"
	// "syscall"
)

var (
	// cstZone = time.FixedZone("CST", 8*3600)
	cpunum = runtime.NumCPU()
)


func main(){
	runtime.GOMAXPROCS(cpunum)
	ch := make(chan *[]byte,30)
	wg := sync.WaitGroup{}
	for i := 0; i < cpunum; i++{
		go WgSortLogs(ch, &wg)	
	}


	for i:=0; i<128;i++  {

		fmt.Println("./data/Treasure_"+strconv.Itoa(i)+".data")
		file2,_ := ioutil.ReadFile("./data/Treasure_"+strconv.Itoa(i)+".data")
		ch <- &file2
		wg.Add(1)
	}
	wg.Wait()
	close(ch)
}

func WgSortLogs(ch chan *[]byte,wg *sync.WaitGroup){
	for true{
		tmp,ok := <-ch
		if !ok{
			break
		}
		// fmt.Println(reflect.TypeOf(tmp))
		ReadFile(*tmp)
		wg.Done()
	}
}

func ReadFile(fname []byte){

	status := 1
	big1024 := big.NewInt(1024)
	var buffer bytes.Buffer
	sp := strings.Split(string(fname),"\n")

	for _,line := range sp[:len(sp)-1]{//优化 这要切片影响效率
		big2,_ :=  new(big.Int).SetString(line[90:len(line)-2], 10)
		buffer.Reset()
		for _,s := range line[15:79]{
			// 58是:号
			if s<58{
				buffer.WriteRune(s)
			}
		}
		big1,_ := new(big.Int).SetString(buffer.String(), 10)
		status = 1

		bigadd := big.NewInt(0)
		bigadd.Add(big1,big1024)
		status *= bigadd.Cmp(big2)


		bigsub := big.NewInt(0)
		bigsub.Sub(big1,big1024)
		status *= bigsub.Cmp(big2)


		bigmul := big.NewInt(0)
		bigmul.Mul(big1,big1024)
		status *= bigmul.Cmp(big2)

		bigrem := big.NewInt(0)
		bigrem.Rem(big1,big1024)
		status *= bigrem.Cmp(big2)
		if status==0{
			// fmt.Println(big2)
		}
	}
	fname=nil
}