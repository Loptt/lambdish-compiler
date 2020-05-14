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

func getNums(v1, v2 interface{}) (float64, float64, error) {
	num1, ok := v1.(float64)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to num")
	}

	num2, ok := v2.(float64)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to num")
	}

	return num1, num2, nil
}

func getChar(v interface{}) (rune, error) {
	char, ok := v.(rune)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to char")
	}

	return char, nil
}

func getChars(v1, v2 interface{}) (rune, rune, error) {
	char1, ok := v1.(rune)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to char")
	}

	char2, ok := v2.(rune)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to char")
	}

	return char1, char2, nil
}

func getBool(v interface{}) (bool, error) {
	boo, ok := v.(bool)
	if !ok {
		return false, errutil.Newf("Cannot convert current value to bool")
	}

	return boo, nil
}

func getBools(v1, v2 interface{}) (bool, bool, error) {
	bool1, ok := v1.(bool)
	if !ok {
		return false, false, errutil.Newf("Cannot convert current value to bool")
	}

	bool2, ok := v2.(bool)
	if !ok {
		return false, false, errutil.Newf("Cannot convert current value to bool")
	}

	return bool1, bool2, nil
}

func getInt(v interface{}) (int, error) {
	in, ok := v.(int)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to address")
	}

	return in, nil
}
