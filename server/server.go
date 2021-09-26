package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	clientscheme "k8s.io/client-go/kubernetes/scheme"
)

func StartServer(addr string) error {
	http.HandleFunc("/serialize", validatePodSpec)
	return http.ListenAndServe(addr, nil)
}

func jsonError(e error, responseCode int, w http.ResponseWriter) {
	w.WriteHeader(responseCode)
	data := []byte(fmt.Sprintf(`
	{
		"error": %q
	}
	`, e.Error()))

	w.Write(data)
}

func validatePodSpec(w http.ResponseWriter, r *http.Request) {
	data := []byte{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		jsonError(fmt.Errorf("can't read request body: %w", err), 500, w)
		return
	}

	fmt.Println(string(data))

	scheme := runtime.NewScheme()
	clientscheme.AddToScheme(scheme)

	jsonSerializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme, scheme, json.SerializerOptions{})
	yamlSerializer := yaml.NewDecodingSerializer(jsonSerializer)

	object, _, err := yamlSerializer.Decode(data, nil, nil)

	if err != nil {
		jsonError(fmt.Errorf("error decoding body: %w", err), 400, w)
		return
	}

	protobufSeralizer := protobuf.NewSerializer(scheme, scheme)

	if err := protobufSeralizer.Encode(object, w); err != nil {
		jsonError(fmt.Errorf("error encoding object: %w", err), 400, w)
		return
	}

	w.WriteHeader(200)
}
