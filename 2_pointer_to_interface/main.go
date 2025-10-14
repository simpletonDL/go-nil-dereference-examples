package main

type MyError struct{}

func (*MyError) Error() string {
	return "faild"
}

func bad() error {
	e := &MyError{}
	return e
}

func bad2() *MyError {
	return &MyError{}
}

func main() {
	var err error = bad()
	if err == nil {
		print("123")
	} else {
		println()
	}
}
