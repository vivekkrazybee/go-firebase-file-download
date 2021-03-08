package main

import (
	"context"
	"fmt"
	"io/ioutil"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// [START cloud_storage_golang]
	config := &firebase.Config{}
	opt := option.WithCredentialsFile("your-credential-file.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		fmt.Println(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	bucket, err := client.Bucket("Your-bucket-name")
	if err != nil {
		fmt.Println(err)
	}

	var a = []string{"location/2021-03-07/abc.html"}
	for _, v := range a {
		//create a firebase bucket object with the given filepath
		obj := bucket.Object(v)
		reader, err := obj.NewReader(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

		defer reader.Close()

		contentBytes, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println(err)
		}

		//your new file name
		fileName := "cdf.html"
		err = ioutil.WriteFile(fileName, contentBytes, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
