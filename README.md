# Veritas

[Hochschule Heilbronn](https://www.hs-heilbronn.de),
[Hochschule Pforzheim](https://designpf.hs-pforzheim.de)

## Dependencies

- MacOS
- Go >=1.21.6
- JavaScript (Node/Bun/...)
- OpenCV 4
- Firebase Storage Project (Google Cloud credentials expected at `/Users/Shared/veritas/serviceAccountKey.json`)

Firebase API URL is hardcoded at `server/internal/firebase/firebase.go uploadImageToFirebaseStorage`.
This relates to the firebase project.

## Build

1. Go to `extension` folder.
2. Run `npm install`.
3. Run `npm run build`.
4. Install the folder `extension` as a Chrome extension.
5. Go to `server` folder.
6. Run `go run ./cmd/server`.
7. Connect the **three** Raspberry Pies to the server.
   If the server Macbook name is `MBP-von-IT` use the adresses:
   `http://MBP-von-IT:8080/image/0`, `http://MBP-von-IT:8080/image/1` and `http://MBP-von-IT:8080/image/2`.
