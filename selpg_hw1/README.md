文件树
---
* test.py:  python脚本，用于生成输入内容input.txt，以及作为第三方程序产生输出然后传给selpg
* selpg.go:  主程序，包含所有代码
* selpg:  build之后产生的可执行文件
* input.txt:  输入文件
* output.txt:  输出文件

设计说明
---
根据原作者的[文档说明](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)，使用go语言实现。  

使用 flag 包对命令行参数进行解析，使用os，bufio对文件、os.Stdin进行读取，输出则调用os.Stdout.Write()，标准错误输出使用 fmt.Fprintf(os.Stderr, "error")。  

整体流程：首先构造Selpg结构体，里面包含了必要的信息。在init()中初始化变量，在main()中执行flag.Parse()解析命令行输入的参数，并且检查参数是否合法。然后根据输入内容的类型，决定调用哪种输出读取方式。  

关键代码部分
---
1. 对于标准输入以及文件类型，关键代码如下：
```
    line := bufio.NewScanner(fileInput)  // fileInput是读取文件，替换为 os.Stdin 即为读取标准输入
    for line.Scan() {
            flag := false
            for _,c := range line.Text() {
                if c == '\f' {
                    if cur_page >= sa.s && cur_page <= sa.e {
                        flag = true
                        os.Stdout.Write([]byte("\n"))
                        stdin.Write([]byte("\n"))  // 此处的stdin是写入管道的输入，用于打印机
                    }
                    cur_page++;
                } else {
                    if cur_page >= sa.s && cur_page <= sa.e {
                        os.Stdout.Write([]byte(string(c)))
                        stdin.Write([]byte(string(c)))
                    }
                }
            }
            if flag != true && cur_page >= sa.s && cur_page <= sa.e {
                os.Stdout.Write([]byte("\n"))
                stdin.Write([]byte("\n"))
            }
            flag = false
    }
```
2. 涉及管道部分
```
    cmd := exec.Command("cat", "-n")
    stdin, err:= cmd.StdinPipe()
    if err != nil {
        panic(err)
    }
    ···
    stdin.Close()
    cmd.Stdout = os.Stdout
    cmd.Start()
```

使用以及测试
---
节选四个用例:  

1. ```./selpg -s 1 -e 1 -l 33 input.txt > output.txt ```

![](https://github.com/CODEJY/ServiceComputing/blob/master/selpg_hw1/ScreenShot/1.png)
2. ``` python test.py | ./selpg -s 1 -e 1 -l 43```

![](https://github.com/CODEJY/ServiceComputing/blob/master/selpg_hw1/ScreenShot/2.png)
3. ``` python test.py | ./selpg -s 1 -e 1 -l 43 > output.txt```

![](https://github.com/CODEJY/ServiceComputing/blob/master/selpg_hw1/ScreenShot/3.png)
4. ``` python test.py | ./selpg -s 1 -e 1 -f=true```

![](https://github.com/CODEJY/ServiceComputing/blob/master/selpg_hw1/ScreenShot/4.png)


