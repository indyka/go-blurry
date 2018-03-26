package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/lazywei/go-opencv/opencv"
)

func main() {

	dirname := "."

	f, err := os.Open(path.Join(dirname, "images"))

	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {

		log.Fatal(err)
	}

	for _, file := range files {
		image := opencv.LoadImage(path.Join(path.Join(dirname, "images"), file.Name()))

		if image == nil {
			log.Fatal(err)
		}

		laplace := opencv.CreateImage(image.Width(), image.Height(), image.Depth(), image.Channels())
		opencv.Laplace(image, laplace, 3)
		_, sigma := opencv.MeanStdDevWithMask(laplace, nil)

		fmt.Println(file.Name(), ": ", sigma.Val()[0]*sigma.Val()[0])
	}

}
