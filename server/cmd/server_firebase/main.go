package main

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"time"

	"encoding/json"
	firebase "firebase.google.com/go"
	"gocv.io/x/gocv"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseAPIKey            string `json:"FIREBASE_API_KEY"`
	FirebaseAuthDomain        string `json:"FIREBASE_AUTH_DOMAIN"`
	FirebaseProjectID         string `json:"FIREBASE_PROJECT_ID"`
	FirebaseStorageBucket     string `json:"FIREBASE_STORAGE_BUCKET"`
	FirebaseMessagingSenderID string `json:"FIREBASE_MESSAGING_SENDER_ID"`
	FirebaseAppID             string `json:"FIREBASE_APP_ID"`
}

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
	defer resp.Body.Close()

	var object FirebaseFileAPI
	err = json.NewDecoder(resp.Body).Decode(&object)
	if err != nil {
		return "", err
	}

	url += "?alt=media&token=" + object.DownloadTokens

	return url, nil

}

// Funktion zum Laden der Firebase-Konfiguration aus einer JSON-Datei
func loadFirebaseConfig(configFilePath string) (Config, error) {
	var config Config

	// Lese die Konfigurationsdatei
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, fmt.Errorf("fehler beim Lesen der Konfigurationsdatei: %v", err)
	}

	// Dekodiere JSON-Daten in die Config-Struktur
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("fehler beim Dekodieren der Konfigurationsdaten: %v", err)
	}

	return config, nil
}

func getAPicture() (image.Image, error) {
	camera, err := gocv.OpenVideoCapture(0)
	if err != nil {
		return nil, fmt.Errorf("fehler beim erkennenn der kamera: %v", err)
	}
	defer camera.Close()

	img := gocv.NewMat()
	defer img.Close()

	if !camera.Read(&img) {
		return nil, fmt.Errorf("fehler beim lesen der kamera")
	}

	imageForImage, err := img.ToImage()
	if err != nil {
		return nil, fmt.Errorf("fehler beim erstellen des byte arrays: %v", err)
	}

	return imageForImage, nil
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
}
