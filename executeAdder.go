package executeadder

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AddToFile(prefix, oldFileName string, newFileName string) {
	f, err := os.Open(oldFileName)
	check(err)

	data := make([]byte, 10000)
	count, err := f.Read(data)
	check(err)
	f.Close()

	finalContent := prefix + string(data[:count])
	os.WriteFile(newFileName, []byte(finalContent), 0644)
	check(err)
}
