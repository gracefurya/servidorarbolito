package handlers

import (
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 //2MB

//RecibirArchivo el archivo que el usuario intente subir se guardara
func RecibirArchivo(r *http.Request, uploadPath, fileName string) error {
	//fmt.Println("recibiendo...")
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return err
	}
	//fmt.Println("recibiendo...2", r.MultipartForm)
	file, fileHeader, err := r.FormFile("picture")
	if err != nil {
		return err
	}
	defer file.Close()

	fileSize := fileHeader.Size
	//fmt.Println("recibiendo...3")
	if fileSize > maxUploadSize {
		return errors.New("EXCEEDED SIZE LIMIT")
	}
	//fmt.Println("recibiendo...4")
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	detectedFileType := http.DetectContentType(fileByte)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
		//case "application", "pdf":
		break
	default:
		return errors.New("INVALID FILE TYPE")
	}

	fileEnding, err := mime.ExtensionsByType(detectedFileType)

	if err != nil {
		return errors.New("CANT READ FILE TYPE")
	}
	newPath := filepath.Join(uploadPath, fileName+fileEnding[0])
	newFile, err := os.Create(newPath)
	if err != nil {
		return errors.New("ERROR_TO_CREATE_NEW_FILE")
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileByte); err != nil || newFile.Close() != nil {
		return errors.New("Error to wirte file")
	}

	return nil
}
