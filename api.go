package vault

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//DownloadRequest - input structure for download request of trustore
type DownloadRequest struct {
	StorePassword string            `json:"storePassword"`
	Data          map[string]string `json:"data"`
}

//DownloadResponse - output response for download request of trustore
type DownloadResponse struct {
	Vault    string `json:"vault"`
	FileName string `json:"fileName"`
}

//DownloadTruststore - download trustore request
func DownloadTruststore(w http.ResponseWriter, req *http.Request) {
	downloadRequest := DownloadRequest{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	err = json.Unmarshal([]byte(data), &downloadRequest)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	randomFileName := RandomString(16) + ".vault"
	filePath := TEMPTRUSTORESPATH + string(os.PathSeparator) + randomFileName
	v := CreateVault(filePath, downloadRequest.StorePassword)
	for key, value := range downloadRequest.Data {
		err = v.Upsert(key, value)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
	}
	v.Close()
	file, err := os.Open(filePath)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	fileData, err := ioutil.ReadAll(file)
	file.Close()
	sEnc := b64.StdEncoding.EncodeToString(fileData)
	defer os.Remove(filePath)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	downloadResponse := DownloadResponse{
		FileName: randomFileName,
		Vault:    sEnc,
	}
	fmt.Println(downloadResponse)
	payload, err := json.Marshal(&downloadResponse)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	fmt.Println(string(payload))
	w.Write(payload)
}
