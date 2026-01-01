package cmd

type fileInfo struct {
	rawFileName   string
	size          int64
	fileName      string
	fileExtension string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
