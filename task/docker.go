package task

/*
#include <mruby.h>
#include <mruby/string.h>
#include <mruby/value.h>
#include <mruby/array.h>

static int _RARRAY_LEN(mrb_value a) { return (RARRAY(a)->len); }
static int _mrb_fixnum(mrb_value o) { return (int) mrb_fixnum(o); }
static float _mrb_float(mrb_value o) { return (float) mrb_float(o); }

extern void docker_init(mrb_state *mrb);

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

	fmt.Println("build")
	return
}
