package vm

import (
	"os"

	"github.com/mewkiz/pkg/errutil"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func getNum(v interface{}) (float64, error) {
	num, ok := v.(float64)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to num")
	}

	return num, nil
}

func getChar(v interface{}) (rune, error) {
	char, ok := v.(rune)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to char")
	}

	return char, nil
}

func getBool(v interface{}) (bool, error) {
	boo, ok := v.(bool)
	if !ok {
		return false, errutil.Newf("Cannot convert current value to bool")
	}

	return boo, nil
}

func getInt(v interface{}) (int, error) {
	in, ok := v.(int)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to address")
	}

	return in, nil
}
