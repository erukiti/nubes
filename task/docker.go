package task

/*
#include <mruby.h>
#include <mruby/string.h>
#include <mruby/value.h>
#include <mruby/array.h>
#include <mruby/hash.h>

static int _RARRAY_LEN(mrb_value a) {return RARRAY_LEN(a);}

extern void docker_init(mrb_state *mrb);

//#define

*/
import "C"
import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"os"
	"path/filepath"
)

func dockerInit(mruby mruby) {
	C.docker_init(mruby.mrb)
}

//export dockerBuild
func dockerBuild(Ctag, CbaseDir *C.char) {
	var err error
	tag := C.GoString(Ctag)
	baseDir := "tasks/" + C.GoString(CbaseDir)

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relpath, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}

		hdr := &tar.Header{Name: relpath, Size: info.Size()}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		body, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if _, err := tw.Write(body); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := tw.Close(); err != nil {
		fmt.Println("error close")
		return
	}

	client, err := docker.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		fmt.Println(err)
		return
	}

	opts := docker.BuildImageOptions{
		Name:         tag,
		InputStream:  buf,
		OutputStream: os.Stderr,
	}
	if err := client.BuildImage(opts); err != nil {
		fmt.Println(err)
		return
	}

	return
}

//export dockerRun
func dockerRun(Cimage *C.char) *C.char {
	var err error
	fmt.Println("docker run")
	image := C.GoString(Cimage)
	client, err := docker.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	config := docker.Config{
		Hostname:     "hoge",
		Image:        image,
		AttachStdout: true,
		AttachStderr: true,
	}
	opts := docker.CreateContainerOptions{
		// Name:   "hogefuga",
		Config: &config,
	}
	container, err := client.CreateContainer(opts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("created")
	hostConfig := docker.HostConfig{}
	err = client.StartContainer(container.ID, &hostConfig)
	fmt.Println("started")

	return C.CString(container.ID)
}
