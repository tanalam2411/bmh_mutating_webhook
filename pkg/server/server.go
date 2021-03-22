package server

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Server contains the functions handling server requests
type Server struct {
	ServerTLSConf *tls.Config
	ClientTLSConf *tls.Config
	CaPEM         []byte
}

func (s Server) GetCA(w http.ResponseWriter, req *http.Request) {
	if len(s.CaPEM) == 0 {
		fmt.Fprintf(w, "No certificate found\n")
		return
	}

	// if base64 parameter is set, return in base64 format
	req.ParseForm()
	if _, hasParam := req.Form["base64"]; hasParam {
		fmt.Fprintf(w, string(base64.StdEncoding.EncodeToString(s.CaPEM)))
		return
	}

	fmt.Fprintf(w, string(s.CaPEM))
}

/*
Pod - RSOrchastrator


*/

func PrintRequestBody(r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Failed to read request body: %v", err)
	}

	var bodyInterface interface{}

	if err = json.Unmarshal(body, &bodyInterface); err != nil {
		fmt.Println("Failed to unmarshal to bodyInterface: %v", err)
	}

	bodyJson, err := json.Marshal(bodyInterface)

	if err != nil {
		fmt.Println("Failed to Marhsal: %v", err)
	}

	fmt.Printf("\n\n\n---------------\n\n\n")
	fmt.Println(string(bodyJson))
	fmt.Printf("\n\n\n---------------\n\n\n")
}

func (s Server) PostWebhook(w http.ResponseWriter, r *http.Request) {
	//PrintRequestBody(r)
	var request AdmissionReviewRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON body in invalid format: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	if request.APIVersion != "admission.k8s.io/v1" || request.Kind != "AdmissionReview" {
		http.Error(w, fmt.Sprintf("wrong APIVersion or kind: %s - %s", request.APIVersion, request.Kind), http.StatusBadRequest)
		return

	}
	fmt.Printf("debug: %+v\n", request.Request)
	fmt.Printf("debug object: %+v\n", request.Request.Object)
	response := AdmissionReviewResponse{
		APIVersion: "admission.k8s.io/v1",
		Kind:       "AdmissionReview",
		Response: Response{
			UID:     request.Request.UID,
			Allowed: true,
		},
	}

	// add label if we're creating a pod
	//if request.Request.Kind.Group == "" && request.Request.Kind.Version == "v1" && request.Request.Kind.Kind == "Pod" && request.Request.Operation == "CREATE" {
	//	patch := `[{"op": "add", "path": "/metadata/labels/myExtraLabel", "value": "webhook-was-here"}]`
	//	patchEnc := base64.StdEncoding.EncodeToString([]byte(patch))
	//	response.Response.PatchType = "JSONPatch"
	//	response.Response.Patch = patchEnc
	//}

	out, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON output marshal error: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	fmt.Printf("Got request, response: %s\n", string(out))
	fmt.Fprintln(w, string(out))
}
