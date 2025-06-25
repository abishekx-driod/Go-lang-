package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"

	// "path"
	"time"

	// "io"
	"os"
)
func main(){
	var filename string
	fmt.Scan(&filename)
	b,_:=ifExistorNot(filename)
	
	createFileIfNotExist(filename)
	fmt.Println(ifExistorNot(filename))
	
	if !b{
		 createFile(filename)
	}
	fmt.Println(ifExistorNot(filename))
	fmt.Println(WriteInAFile(filename,[]byte("Hello My name is mano From Third CSE")))
	// var bb[100]byte
	// f,_:=os.Open(filename)
	// for{
	// 	_,err:=f.Read(bb[:])
	// 	if err==io.EOF{
	// 		break
	// 	}

	// }
	//  fmt.Printf("%s",bb)
	// f.Close()
	AppendFile("onr.txt",[]byte("At Karpagam Institute"))
	 bbb,_:=os.ReadFile("onr.txt")
	fmt.Printf("%s",bbb)
	size,modTime,_:=getFileStatus(filename)
	fmt.Println(size,modTime)

	var src string
	var dest string
	fmt.Scan(&src)
	fmt.Scan(&dest)
	copyfile(src,dest)

	//delete file
	os.Remove(filename)
 listedFiles,_:=makedir("./")
 fmt.Println(listedFiles)
	
}
func AppendFile(path string,data []byte){
	f,err:=os.OpenFile(path,os.O_WRONLY|os.O_APPEND,0664,)
	if err!=nil{
		panic(err)
	}
	f.Write(data)
	f.Close()
}
func WriteInAFile(path string,b []byte)error{
	return os.WriteFile(path,b,0644);
}
func ifExistorNot(filename string)(bool,error){
	_,err:=os.Lstat(filename)
	if err!=nil{
		if errors.Is(err,os.ErrNotExist){
			return false,nil
		}
		return false,err
	}
	return true,nil
}
func createFile(filename string) error{
	file,err:=os.Create(filename)
	if err!=nil{
		return err
	}
	defer file.Close()
	return nil
}
func createFileIfNotExist(filename string)error{
	file,err:=os.OpenFile(filename,os.O_RDWR |os.O_EXCL|os.O_CREATE,0666)
	if err!=nil{
		return err;
	}
	defer file.Close()
	return nil
}
// file status
func getFileStatus(path string)(int64,time.Time,error){
	info,err:=os.Lstat(path)
	if err!=nil{
		return -1,time.Time{},err
	}
	size:=info.Size()
	modTime:=info.ModTime()
	return size,modTime,nil
}
func copyfile(src string,dest string){
	f1,err:=os.Open(src)
	if err!=nil{
		panic(err)
	}
	defer f1.Close()
	f2,err:=os.Create(dest)
	if err!=nil{
		panic(err)
	}
	defer f2.Close()
	io.Copy(f2,f1)
}
func makedir(dirpath string)([]fs.DirEntry,error){
	entries,err:=os.ReadDir(dirpath)
	listery:=make([]fs.DirEntry,0)
	if err!=nil{
		return nil,err
	}
	for _,entry:=range entries{
		listery = append(listery, entry)
	}
	return listery,nil
}