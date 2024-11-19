package goutil

import "os"

func WriteTmpFile(prefix, extension, content string) (string, error) {
	f, err := os.CreateTemp(os.TempDir(), prefix+"_*."+extension)
	if err != nil {
		return "", err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			Logger().Error(err)
		}
	}(f)

	_, err = f.WriteString(content)
	return f.Name(), err
}
