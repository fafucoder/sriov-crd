package testing

import (
	sriovfake "github.com/fafucoder/sriov-crd/pkg/client/clientset/versioned/fake"
	sriovclient "github.com/fafucoder/sriov-crd/pkg/client/clientset/versioned/typed/sriov/v1"
	"github.com/fafucoder/sriov-crd/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

type tmpSysFs struct {
	dirRoot      string
	dirList      []string
	fileList     map[string][]byte
	netSymlinks  map[string]string
	devSymlinks  map[string]string
	vfSymlinks   map[string]string
	originalRoot *os.File
}

var ts = tmpSysFs{
	dirList: []string{
		"sys/class/net",
		"sys/bus/pci/devices",
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1/net/enp175s0f1",
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.0/net/enp175s6",
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.1/net/enp175s7",
		"sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0/net/ens1",
		"sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0/net/ens1d1",
	},
	fileList: map[string][]byte{
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1/sriov_numvfs":   []byte("2"),
		"sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0/sriov_numvfs":   []byte("0"),
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1/sriov_totalvfs": []byte("2"),
		"sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0/sriov_totalvfs": []byte("0"),
	},

	netSymlinks: map[string]string{
		"sys/class/net/enp175s0f1": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1/net/enp175s0f1",
		"sys/class/net/enp175s6":   "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.0/net/enp175s6",
		"sys/class/net/enp175s7":   "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.1/net/enp175s7",
		"sys/class/net/ens1":       "sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0/net/ens1",
		"sys/class/net/ens1d1":     "sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0/net/ens1d1",
	},
	devSymlinks: map[string]string{
		"sys/class/net/enp175s0f1/device": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1",
		"sys/class/net/enp175s6/device":   "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.0",
		"sys/class/net/enp175s7/device":   "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.1",
		"sys/class/net/ens1/device":       "sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0",
		"sys/class/net/ens1d1/device":     "sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0",

		"sys/bus/pci/devices/0000:af:00.1": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1",
		"sys/bus/pci/devices/0000:af:06.0": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.0",
		"sys/bus/pci/devices/0000:af:06.1": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.1",
		"sys/bus/pci/devices/0000:05:00.0": "sys/devices/pci0000:00/0000:00:02.0/0000:05:00.0",
	},
	vfSymlinks: map[string]string{
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1/virtfn0": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.0",
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.0/physfn":  "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1",

		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1/virtfn1": "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.1",
		"sys/devices/pci0000:ae/0000:ae:00.0/0000:af:06.1/physfn":  "sys/devices/pci0000:ae/0000:ae:00.0/0000:af:00.1",
	},
}

// CreateTmpSysFs create mock sysfs for testing
func CreateTmpSysFs() error {
	originalRoot, err := os.Open("/")
	ts.originalRoot = originalRoot

	tmpdir, err := ioutil.TempDir("/tmp", "sriovcrd-testfile-")
	if err != nil {
		return err
	}

	ts.dirRoot = tmpdir
	//syscall.Chroot(ts.dirRoot)

	for _, dir := range ts.dirList {
		if err := os.MkdirAll(filepath.Join(ts.dirRoot, dir), 0755); err != nil {
			return err
		}
	}
	for filename, body := range ts.fileList {
		if err := ioutil.WriteFile(filepath.Join(ts.dirRoot, filename), body, 0644); err != nil {
			return err
		}
	}

	for link, target := range ts.netSymlinks {
		if err := createSymlinks(filepath.Join(ts.dirRoot, link), filepath.Join(ts.dirRoot, target)); err != nil {
			return err
		}
	}

	for link, target := range ts.devSymlinks {
		if err := createSymlinks(filepath.Join(ts.dirRoot, link), filepath.Join(ts.dirRoot, target)); err != nil {
			return err
		}
	}

	for link, target := range ts.vfSymlinks {
		if err := createSymlinks(filepath.Join(ts.dirRoot, link), filepath.Join(ts.dirRoot, target)); err != nil {
			return err
		}
	}

	utils.SysBusPci = filepath.Join(ts.dirRoot, utils.SysBusPci)
	utils.NetDirectory = filepath.Join(ts.dirRoot, utils.NetDirectory)
	return nil
}

func createSymlinks(link, target string) error {
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	if err := os.Symlink(target, link); err != nil {
		return err
	}

	return nil
}

// RemoveTmpSysFs removes mocked sysfs
func RemoveTmpSysFs() error {
	err := ts.originalRoot.Chdir()
	if err != nil {
		return err
	}
	if err = syscall.Chroot("."); err != nil {
		return err
	}
	if err = ts.originalRoot.Close(); err != nil {
		return err
	}
	if err = os.RemoveAll(ts.dirRoot); err != nil {
		return err
	}
	return nil
}

//Create Sriov Client
func SriovClient() sriovclient.K8sCniCncfIoV1Interface {
	return sriovfake.NewSimpleClientset().K8sCniCncfIoV1()
}
