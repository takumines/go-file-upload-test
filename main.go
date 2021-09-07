package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		log.Fatal("許可されていないメソッド")
	}

	// formから送信されたファイルを解析
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		log.Fatal("ファイルのアップロード失敗")
	}

	// アップロードされたファイル名を取得
	uploadedFileName := fileHeader.Filename

	// ファイルを置くパスを作成
	imagePath := "img/" + uploadedFileName

	// imagePathにアップロードされたファイルを保存
	saveImage, err := os.Create(imagePath)
	if err != nil {
		log.Fatal("ファイルの保存に失敗")
	}

	// 保存用ファイルにアップロードされたファイルを書き込む
	_, err = io.Copy(saveImage, file)
	if err != nil {
		log.Fatal("アップロードしたファイルの書き込みに失敗")
	}

	//saveImageとfileを最後に閉じる
	defer saveImage.Close()
	defer file.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	if err := tmp.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
