package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type Selpg struct {
	s         int
	e         int
	pageLen   int
	pageType  bool
	inputFile string
	printDest string
}

var selpg Selpg

func init() {
	selpg = Selpg{
		s:         -1,
		e:         -1,
		pageLen:   72,
		pageType:  false,
		inputFile: "",
		printDest: "",
	}
	flag.IntVar(&selpg.s, "s", -1, "specify start page(>=1)")
	flag.IntVar(&selpg.e, "e", -1, "specify end page(>=s)")
	flag.IntVar(&selpg.pageLen, "l", 72, "specify length of one page")
	pageType_ := flag.Bool("f", false, "-f can be set true or false,if true ,it will auto paging by find ‘\\f’,but you can't set -f and -l at the selpgme time.")
	printDest_ := flag.String("d", "", "specify print dest.")
	selpg.pageType = *pageType_
	selpg.printDest = *printDest_
	flag.Usage = func() {
		fmt.Printf("Usage of seplg:\n")
		fmt.Printf("seplg -s num1 -e num2 [-f -l num3 -d str1 file]\n")
		flag.PrintDefaults()
	}
}

func main() {
	// 提取命令行参数
	flag.Parse()
	// 检查输入参数的错误
	checkCmdError()
	// 获取non-flag参数
	if len(flag.Args()) == 1 {
		selpg.inputFile = flag.Args()[0]
	}
	// 如果pageType是true，那么自动识别分页符'\f'进行分页，如果为false，那么根据指定的行数进行分页
	if selpg.pageType == false {
		pageByFixedLines(selpg, selpg.inputFile != "", selpg.printDest != "")
	} else {
		autoPaging(selpg, selpg.inputFile != "", selpg.printDest != "")
	}
}

// 检查输入参数的错误
func checkCmdError() {
	if selpg.s == -1 || selpg.e == -1 || selpg.s > selpg.e || selpg.s < 1 || selpg.e < 1 {
		flag.Usage()
		return
	}

	if selpg.pageLen != 72 && selpg.pageType == true {
		flag.Usage()
		return
	}
	if len(flag.Args()) > 1 {
		flag.Usage()
		return
	}
}

// pageType为false时调用，根据指定行数分页
func pageByFixedLines(selpg Selpg, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	curPage := 1
	curLines := 0
	if file {
		fileInput, err := os.OpenFile(selpg.inputFile, os.O_RDWR, os.ModeType)
		defer fileInput.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(fileInput)
		for line.Scan() {
			if curPage >= selpg.s && curPage <= selpg.e {
				os.Stdout.Write([]byte(line.Text() + "\n"))
				stdin.Write([]byte(line.Text() + "\n"))
			}
			curLines++
			if curLines %= selpg.pageLen; curLines == 0 {
				curPage++
			}
		}
	} else {
		tmpScan := bufio.NewScanner(os.Stdin)
		for tmpScan.Scan() {
			if curPage >= selpg.s && curPage <= selpg.e {
				os.Stdout.Write([]byte(tmpScan.Text() + "\n"))
				stdin.Write([]byte(tmpScan.Text() + "\n"))
			}
			curLines++
			if curLines %= selpg.pageLen; curLines == 0 {
				curPage++
			}
		}
	}
	if curPage < selpg.e {
		fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
	}
	if pipe {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}

}

// pageTypr为true调用，自动识别换行符分页
func autoPaging(selpg Selpg, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	curPage := 1
	if file {
		fileInput, err := os.OpenFile(selpg.inputFile, os.O_RDWR, os.ModeType)
		defer fileInput.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(fileInput)
		for line.Scan() {
			flag := false
			for _, c := range line.Text() {
				if c == '\f' {
					if curPage >= selpg.s && curPage <= selpg.e {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					curPage++
				} else {
					if curPage >= selpg.s && curPage <= selpg.e {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && curPage >= selpg.s && curPage <= selpg.e {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	} else {
		tmpScan := bufio.NewScanner(os.Stdin)
		for tmpScan.Scan() {
			flag := false
			for _, c := range tmpScan.Text() {
				if c == '\f' {
					if curPage >= selpg.s && curPage <= selpg.e {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					curPage++
				} else {
					if curPage >= selpg.s && curPage <= selpg.e {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && curPage >= selpg.s && curPage <= selpg.e {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	}
	if curPage < selpg.e {
		fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
	}
	if pipe {

		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}
