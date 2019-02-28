package CharacterCount

import (
	
	"os"
	"io"
	"bufio"
	"io/ioutil"
	"unsafe"
	"strings"
	"unicode/utf8"
)

//CountByLine 按行统计中文字符 bufio
//map[中字]:次数-字符集   int-中文总字数   
func CountByLine( fileName string ) (map[rune]int,int) {
	var invalide,totalCount int
	dictionary := make(map[rune]int)
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		panic("fileName can't be nil!")
	}
	
	//打开文件
	f, err := os.Open(fileName)
	if err != nil{
		panic(err)
	}
	//函数结束时 关闭文件
	defer f.Close()

	//添加通道
	r := bufio.NewReader(f)
	for {
		//读取数据，直到指定字符出现，包括指定字符
		line, err := r.ReadString('\n')
		if err != nil || err == io.EOF {
			if line == "" {
				break
			}
		}
		//消除字符串中的空白符：\t,\n,\r,空格等
		line = strings.TrimSpace(line)

		for i := 0 ; i < len(line); {
			r,size := utf8.DecodeRuneInString(line[i:])
			if r >= 0x4E00 && r <= 0x9FA5{
				dictionary[r]++
				totalCount++
				//fmt.Printf("中文\t%c:%#[1]x\n",r)
			}else{
				invalide++
				//fmt.Printf("符号\t%c:%#[1]x\n",r)
			}
			i+=size
		}		
	}
	return dictionary,totalCount
}

//countByAll_Demo 一次性读取文本，然后统计中文字符 ioutil
//map[中字]:次数-字符集   int-中文总字数 
//有部分操作可以避免
func countByAll_Demo( fileName string ) (map[rune]int,int){
	var invalide,totalCount int
	dictionary := make(map[rune]int)
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		panic("fileName can't be nil!")
	}

	r,err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	//转换成字符串且去除空白符
	strR := strings.TrimSpace(string(r))
	//转换成slice  []rune
	content := []rune(strR)
	for _,value := range(content){
		if value >= 0x4E00 && value <= 0x9FA5{
			dictionary[value]++
			totalCount++
		}else{
			invalide++
		}
	}

	return dictionary,totalCount
}


//CountByAll 一次性读取文本，然后统计中文字符 ioutil
//map[中字]:次数-字符集   int-中文总字数 
//优化后
func CountByAll( fileName string ) (map[rune]int,int){
	var invalide,totalCount int
	dictionary := make(map[rune]int)
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		panic("fileName can't be nil!")
	}

	r,err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//转换成字符串且去除空白符
	strR := strings.TrimSpace(bytes2str(r))
	//range会隐式的unicode解码
	for _,value := range(strR){
		if value >= 0x4E00 && value <= 0x9FA5{
			dictionary[value]++
			totalCount++
		}else{
			invalide++
		}
	}

	return dictionary,totalCount
}

//字节slice转string 优化后的函数
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}