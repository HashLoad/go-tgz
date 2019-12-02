package tgz

import (
	"io"
	"net/http"
	"os"
	"path"
	"testing"
)

func makeFile() io.ReadCloser {
	url := "https://kubernetes-charts-incubator.storage.googleapis.com/schema-registry-1.1.7.tgz"

	resp, err := http.Get(url)
	HandleError(err)

	return resp.Body
}

func TestUnTgz(t *testing.T) {
	pwd, err := os.Getwd()
	HandleError(err)
	out := path.Join(pwd, "out/root")

	UnTgz(makeFile(), "/", out)
}

func TestUnTgzInitialPath(t *testing.T) {
	pwd, err := os.Getwd()
	HandleError(err)
	out := path.Join(pwd, "out/cut")

	UnTgz(makeFile(), "/schema-registry/charts/kafka/charts/zookeeper/templates", out)
}
