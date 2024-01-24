package main

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"encoding/json"
	firebase "firebase.google.com/go"
	"gocv.io/x/gocv"
	"google.golang.org/api/option"
)

type FirebaseFileAPI struct {
	Name            string    `json:"name"`
	Bucket          string    `json:"bucket"`
	Generation      string    `json:"generation"`
	Metageneration  string    `json:"metageneration"`
	ContentType     string    `json:"contentType"`
	TimeCreated     time.Time `json:"timeCreated"`
	Updated         time.Time `json:"updated"`
	StorageClass    string    `json:"storageClass"`
	Size            string    `json:"size"`
	Md5Hash         string    `json:"md5Hash"`
	ContentEncoding string    `json:"contentEncoding"`
	Crc32C          string    `json:"crc32c"`
	Etag            string    `json:"etag"`
	DownloadTokens  string    `json:"downloadTokens"`
}

// Funktion zum Hochladen eines Bildes nach Firebase Storage
func uploadImageToFirebaseStorage(imageData image.Image) (string, error) {
	ctx := context.Background()

	// Erstelle eine neue Firebase-App mit den bereitgestellten Optionen
	opt := option.WithCredentialsFile("/Users/mhammel/GolandProjects/inno-lab/server/cmd/server_firebase/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return "", fmt.Errorf("fehler beim Erstellen der Firebase-App: %v", err)
	}

	// Erstelle einen Storage-Client
	client, err := app.Storage(ctx)
	if err != nil {
		return "", fmt.Errorf("fehler beim Erstellen des Storage-Clients: %v", err)
	}

	// Erstelle einen eindeutigen Dateinamen, z.B., basierend auf der aktuellen Zeit
	fileName := fmt.Sprintf("%d.jpg", time.Now().UnixNano())

	// Erstelle einen Bucket-Handle
	var bucketName = "inno-lab-85f72.appspot.com"
	bucketHandle, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("fehler beim Abrufen des Bucket-Handles: %v", err)
	}

	// Erstelle einen Storage-Handle für das Bild
	object2000 := bucketHandle.Object(fileName)
	imageHandle := object2000.NewWriter(ctx)

	// Kopiere den Bildinhalt in den Storage-Writer

	if err := jpeg.Encode(imageHandle, imageData, nil); err != nil {
		return "", fmt.Errorf("fehler beim Kopieren des Bildinhalts: %v", err)
	}

	// Schließe den Storage-Writer und committe die Änderungen
	if err := imageHandle.Close(); err != nil {
		return "", fmt.Errorf("fehler beim Schließen des Storage-Handles: %v", err)
	}

	// Erhalte den öffentlichen URL-Link des gespeicherten Bildes
	attrs, err := object2000.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("fehler beim Abrufen der Objektattribute: %v", err)
	}

	imageName := attrs.Name

	url := "https://firebasestorage.googleapis.com/v0/b/inno-lab-85f72.appspot.com/o/" + imageName

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("fehler beim http request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var object FirebaseFileAPI
	err = json.NewDecoder(resp.Body).Decode(&object)
	if err != nil {
		return "", err
	}

	url += "?alt=media&token=" + object.DownloadTokens

	return url, nil

}

// Funktion zum Laden der Firebase-Konfiguration aus einer JSON-Datei

func getAPicture() (image.Image, error) {
	camera, err := gocv.OpenVideoCapture(0)
	if err != nil {
		return nil, fmt.Errorf("fehler beim erkennenn der kamera: %v", err)
	}
	defer func(camera *gocv.VideoCapture) {
		err := camera.Close()
		if err != nil {

		}
	}(camera)

	img := gocv.NewMat()
	defer func(img *gocv.Mat) {
		err := img.Close()
		if err != nil {

		}
	}(&img)

	if !camera.Read(&img) {
		return nil, fmt.Errorf("fehler beim lesen der kamera")
	}

	imageForImage, err := img.ToImage()
	if err != nil {
		return nil, fmt.Errorf("fehler beim erstellen des byte arrays: %v", err)
	}

	return imageForImage, nil
}

func saveImageAsJPEG(img image.Image, directory string) (string, error) {
	// Erstelle einen eindeutigen Dateinamen, z.B., basierend auf der aktuellen Zeit
	fileName := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
	filePath := filepath.Join(directory, fileName)

	// Öffne die Datei zum Schreiben
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("fehler beim Öffnen der Datei zum Schreiben: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Speichere das Bild als JPEG
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return "", fmt.Errorf("fehler beim Speichern des Bildes als JPEG: %v", err)
	}

	return filePath, nil
}

func main() {

	thepicture, err := getAPicture()
	if err != nil {
		log.Fatal(err)
	}

	url, err := uploadImageToFirebaseStorage(thepicture)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("URL des gespeicherten Bildes: %s\n", url)

	// Speichere das Bild als JPEG in gewünschtem Verzeichnis
	// err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Bild erfolgreich gespeichert: %s\n", savedFilePath)

}
