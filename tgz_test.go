package tgz

import (
	"net/http"
	"os"
	"path"
	"testing"
)

func TestUnTgz(t *testing.T) {
	url := "https://kubernetes-charts-incubator.storage.googleapis.com/schema-registry-1.1.7.tgz"
	pwd, err := os.Getwd()
	HandleError(err)
	out := path.Join(pwd, "out/")

	resp, err := http.Get(url)
	HandleError(err)

	UnTgz(resp.Body, "/", out)
}
