package firebase

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"time"

	"encoding/json"

	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"
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

var (
	app    *firebase.App
	client *storage.Client
)

// Funktion zum Hochladen eines Bildes nach Firebase Storage
func uploadImageToFirebaseStorage(imageData image.Image) (string, error) {
	ctx := context.Background()

	// Erstelle eine neue Firebase-App mit den bereitgestellten Optionen
	opt := option.WithCredentialsFile("/Users/mhammel/GolandProjects/inno-lab/server/cmd/server_firebase/serviceAccountKey.json")
	if app == nil {
		var err error
		app, err = firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return "", errors.Join(errors.New("failed to create firebase app"), err)
		}
	}

	// Erstelle einen Storage-Client
	if client == nil {
		var err error
		client, err = app.Storage(ctx)
		if err != nil {
			return "", errors.Join(errors.New("failed to create firebase storage client"), err)
		}
	}

	// Erstelle einen eindeutigen Dateinamen, z.B., basierend auf der aktuellen Zeit
	fileName := fmt.Sprintf("%d.jpg", time.Now().UnixNano())

	// Erstelle einen Bucket-Handle
	var bucketName = "inno-lab-85f72.appspot.com"
	bucketHandle, err := client.Bucket(bucketName)
	if err != nil {
		return "", errors.Join(errors.New("failed to create firebase bucket handle"), err)
	}

	// Erstelle einen Storage-Handle für das Bild
	object2000 := bucketHandle.Object(fileName)
	imageHandle := object2000.NewWriter(ctx)
	defer imageHandle.Close()

	// Kopiere den Bildinhalt in den Storage-Writer

	if err := jpeg.Encode(imageHandle, imageData, nil); err != nil {
		return "", errors.Join(errors.New("failed to encode image"), err)
	}

	// Erhalte den öffentlichen URL-Link des gespeicherten Bildes
	attrs, err := object2000.Attrs(ctx)
	if err != nil {
		return "", errors.Join(errors.New("failed to get image attributes"), err)
	}

	imageName := attrs.Name

	url := "https://firebasestorage.googleapis.com/v0/b/inno-lab-85f72.appspot.com/o/" + imageName

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Join(errors.New("failed to get image url"), err)
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

func takePicture() (image.Image, error) {
	vidCapture, err := gocv.OpenVideoCapture(0)
	if err != nil {
		return nil, errors.Join(errors.New("failed to open video capture"), err)
	}
	defer vidCapture.Close()

	mat := gocv.NewMat()
	defer mat.Close()

	if !vidCapture.Read(&mat) {
		return nil, errors.New("failed to read frame from video capture")
	}

	imageForImage, err := mat.ToImage()
	if err != nil {
		return nil, errors.Join(errors.New("failed to convert mat to image"), err)
	}

	return imageForImage, nil
}

func GetWebcamUrl() (string, error) {
	img, err := takePicture()
	if err != nil {
		return "", errors.Join(errors.New("failed to take picture"), err)
	}

	url, err := uploadImageToFirebaseStorage(img)
	if err != nil {
		return "", errors.Join(errors.New("failed to upload image to firebase storage"), err)
	}

	return url, nil
}
