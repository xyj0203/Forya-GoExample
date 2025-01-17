package main

func main() {
	//str := ""
	//slice := strings.Split(str, ",")
	//fmt.Printf("%d", len(slice))
	copy1()
}

func copy1() {
	// 先打开文件对象
	f, _ := OpenFile("foo.dat")

	// 绑定到了 f 对象
	// func Close() error
	var Close = func() error {
		return (*File).Close(f)
	}

	// 绑定到了 f 对象
	// func Read(offset int64, data []byte) int
	var Read = func(offset int64, data []byte) int {
		return (*File).Read(f, offset, data)
	}
}

type File struct {
	fd int
}

// 打开文件
func OpenFile(name string) (f *File, err error) {
	// ...
}

// 关闭文件
func CloseFile(f *File) error {
	// ...
}

// 读文件数据
func ReadFile(f *File, offset int64, data []byte) int {
	// ...
}
