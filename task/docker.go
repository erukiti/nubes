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
func dockerRun(Cimage *C.char) {
	var err error
	fmt.Println("docker run")
	image := C.GoString(Cimage)
	client, err := docker.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		fmt.Println(err)
		return
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
		return
	}
	fmt.Println("created")
	hostConfig := docker.HostConfig{}
	err = client.StartContainer(container.ID, &hostConfig)
	fmt.Println("started")
}

//export dockerCreateContainer
func dockerCreateContainer(mrb *C.mrb_state, Ropts C.mrb_value) *C.char {
	config := docker.Config{}
	keys := C.mrb_hash_keys(mrb, Ropts)
	var name string
	name = ""
	for i := 0; i < int(C._RARRAY_LEN(keys)); i++ {
		key := C.mrb_ary_ref(mrb, keys, C.mrb_int(i))
		val := C.mrb_hash_get(mrb, Ropts, key)
		switch mrbSymToString(mrb, key) {
		case "Hostname":
			config.Hostname = mrbToString(mrb, val)
		case "Domainname":
			config.Domainname = mrbToString(mrb, val)
		case "User":
			config.User = mrbToString(mrb, val)
		case "Memory":
			config.Memory = mrbToUint64(mrb, val)
		case "MemorySwap":
			config.MemorySwap = mrbToUint64(mrb, val)
		case "CpuShares":
			config.CpuShares = mrbToInt64(mrb, val)
		case "Image":
			config.Image = mrbToString(mrb, val)
		case "AttachStdin":
			config.AttachStdin = mrbToBool(mrb, val)
		case "AttachStdout":
			config.AttachStdout = mrbToBool(mrb, val)
		case "AttachStderr":
			config.AttachStderr = mrbToBool(mrb, val)
		case "Tty":
			config.Tty = mrbToBool(mrb, val)
		case "OpenStdin":
			config.OpenStdin = mrbToBool(mrb, val)
		case "StdinOnce":
			config.StdinOnce = mrbToBool(mrb, val)
		case "VolumesFrom":
			config.VolumesFrom = mrbToString(mrb, val)
		case "WorkingDir":
			config.WorkingDir = mrbToString(mrb, val)
		case "NetworkDisabled":
			config.AttachStderr = mrbToBool(mrb, val)
		case "Name":
			name = mrbToString(mrb, val)
		}
	}
	var opts docker.CreateContainerOptions
	if name == "" {
		opts = docker.CreateContainerOptions{
			Config: &config,
		}
	} else {
		opts = docker.CreateContainerOptions{
			Name:   name,
			Config: &config,
		}
	}
	client, err := docker.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	container, err := client.CreateContainer(opts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("created")
	return C.CString(container.ID)
}
