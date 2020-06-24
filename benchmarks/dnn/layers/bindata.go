// Code generated by "esc -private -pkg=layers -o=bindata.go trans.hsaco gpu_gemm.hsaco relu.hsaco maxpooling.hsaco"; DO NOT EDIT.

package layers

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/gpu_gemm.hsaco": {
		name:    "gpu_gemm.hsaco",
		local:   "gpu_gemm.hsaco",
		size:    13800,
		modtime: 1591217294,
		compressed: `
H4sIAAAAAAAC/+xby28TVxc/cz0eO5MH+b5uCKB2mlYKIGZij+PEyYa8moBIQiAtj1YU3XiuHSfzsMbj
NKnAhLRELJCKqkpd0mWl9m8gQeqiy4Y1i26QUFds2mVdzePaM04MQTwF85M8Z+ace+45v3vnjn3H9177
ZGoCMcwweIjAn8DYJ4J7TQ2PZFcedXQZiMMwtAEPHACwvnKNcosJyrinZzy/ZviyMyihs+4X9eXXKK+g
oPT72blCwtM3yCIEJfVDz+hH+Z19aCnsHvz8+dk489BSOHh2sLQ9Ke8G+bg9KFmfX9yLPzI97hSnfXPQ
uR9cPQuxGjeqG5ken5z9zC3bBQCtnh5rSj6ri1hT7M9CCYtiPreSSaS8evV2AN4rK4oif46YpYKhDwkU
XwjJY0JCuMSfIqZO1NIQLwiiMIM1Ui8jCEKeaBpvn8ytavOG6rP32KbhJaXHMU9hPV/G+brz6SLRx6aE
sYC1loUTXRYuOdYRM++Et7FLChpPzz5dLZKAuaBbNeNc4eugX1/NNKIW8vrQrqZzWC2TUwVdoebRVUcV
LGAHpgVOpuR6xdnsmTJW61WPkxwuq1ZzMvrbRGbpbSKD1eICbk4opxr41VOaeC5K88R6yxiNNKfT4/A5
2tOcUaY5o0xzRpOqMY/V0XIuR8w90lIUc66Is4SSc6t4Dtqj7ybtsXeT9vg7QfsFZX6ioChEd4OfzuVK
xLrwhG+I/r6XH//ia47/+WuIP2PoT/pizuzxtgmzem1ZTZdVqzBpFpS5VT07YuafO8MxQyGzplGs/bS3
JxrYzM+RvEZ0y00+maRPlEnTKBc920RhhShugYRnnjULy9gizQsEa/f400zP42WSMw0aVRBq42CmrM1N
zp4t1doqla5bzgUsSRpqGq9MqNg6b5hLbtZOpXK6n5ckiW8+f7TnfAci3I55NOP7HHBVa/ZE0L7+ZeJH
pxiC4BzWgT0DgxAhQoQIESJEiBAhQoQIESLEmwDGm78zzr+7kYZJ/E4g5le43QLQCsGXCUXf+UcNNpZl
uWq1Wn0T+fNw++46QluszT7Cb4HTCvxWxvlr/qe71+DWJl9lNuzs40z8e4hHrrMbsBZdY9YZxFUAsRWA
Ew8QAIoAyAC//dEOCIpwaxPBjfstLALEsjKKILkN3bwKi2sbIKzfPQ43N78BtNVpx2M6tvYDQMd1dnMN
sZWP4Ybj24EQ2PGjiK9ApK1yh+W6I/Dd/XUWgV2O5Tg52hLPdCC+cqetoztm29oQcNC13d6GIIb+X4lB
1zbXgSAOh7Zb/mf39OyDGEDMlhGA+D2el9f4H65GoGv7Ww5BZfGvDRYOba/FXQ7xfeyxR9WNTQZu3GcO
gpNPC4pXWhFfYRiQ7wDqBjsu2DE4mY3EM17dER5AjkZQxmkLjuP3sdwxgOIDZwEBXN8MR1+IECFChAgR
IkSIECFCvDrQteaP97my1bve78moJy900vly0O/vf6uGLQXPTteVD3fuHu/E2JiQVbGeF5bd9dZCMiEl
pIRwuFfBFu5dJPpSQS+JXxnmUqmIs6Q3a2jFskVE08hqoklUMSUlesmKRUwdq70L2axoGVavqi5rYtE0
FknW6nUD9CUG0v2p/r6B5CDJpLCsKEoaz5MEzuXIQCIxIKeyipJJy0eEw/O4RBTB0AU7vZSUkJKDfYMp
cSBN8GBaFr2ahCMwVdCXiDkkTE2Nv4zEVVXZe9pPfa9j99KVD4P6mKf/vUHf7ukPdwf173n6lQb9+87L
oFh9X4OHribrSEDSDYuApKzqpVUNpLxelhZwaQG8o623TJAssmI5V1grZEHKGppGdAuk0qpm4XmQSgsl
y3TPXAmjo4nLSeeYco59zjHtLkC5PH5xZmT65NgLeU8W8611aba/ovbOC3a2e6vPjY43KhO+8cb49pHQ
cWgX+6daNag/HW812ZBWHHb2S9Rnp+OTyv0N/myD/MDb94EangdUcrvef3X0+PfgQPN9O80qED3fCFU0
2U8TbeBPH0P9XpUNtysUPcU0s3t4Ko/7+96HraQrf4b685Pbpf8m/bn7sODtq7r4lPY708T/nlxv3yf5
/xcAAP//aVmkveg1AAA=
`,
	},

	"/maxpooling.hsaco": {
		name:    "maxpooling.hsaco",
		local:   "maxpooling.hsaco",
		size:    22608,
		modtime: 1593038647,
		compressed: `
H4sIAAAAAAAC/+x8b2wcx3X429m9vb25vd29v7zjH+lE62fJlEiRqwuz0i9JLVk27Vqy5Sh14qaBcOId
KVLHO+J4sqW2t14yFnl2jUYwBAUQBDPol7SBv/RLP5IUUBRF09Y6AkWA1h+KokGAfE0b9ENAFm9mlrw7
iJZdu3BK3wDS476Z9+a9N++9eTsc7pvPnn+OSNLTIJoM/woS/jDKn/2Oi7/H4RDDOaDB06ADBRUAlJZx
nXBDaoeawEuCbq/2wVPtEKxdukCLfJ1QS7bDVjqVKeIr1A7noR36dORT0vn6ffPntYLyCeha5cP2ys9r
BRU+fVN8exLYFbwF/ubJdqi00Gli/jMXzrHh/tr0MX/geAWCO7r5uDMXzk2gX+DYDACEBT4/V5ieLA/n
5wr47+pCfnh4euqGM3pS8P3jJwGoGDs8PExfLVYXZirl01m/fTc7djw7mv0efbFYLRdLC6dpNjucfSk/
V9wdk81mL+RvXKxUSs9Vqm/kqwWKqEs3565USi0jj7QPevpa4QgbeD5fnr6en95l+PJ8sfzM+ewzbb07
kjGJ7Oz3WO+Z6jQTCdsjxCrXrlaL+cIC9RHfujlfbBs1U67tdF6a+cN28txO15nSzHT59CO7Xs2Xrhdf
nCkX/O6zNxmqfQBO7A944aS9y3hy8pXr+dIu63PFqfz1Um1vna5UarXK3OVCvpbfW60jU6VKvjZ0ZG/d
nL11c/bWbaJUuZIvnb0+NVWs7q3gc60KFgrVS/P5yaKvJmfxGQxQvj63n9Zz8mq+jJG1n3S6WpyZvlrb
Txq9MVOoXd1PCs1XKqVi4fL+Wymh2L5bsGtsA768H3V6Yz/ptFCrzhSK+2udhE77ap3m84X9tUio0L5a
oVpl/ktc5s7lF649Tv2Z8hem/Aufp/Kfk+TPzxQKxTKf/OWpqYVi7Tsfo8B47n9//te+4Pl//wuY/6VK
+eNSgvMJ3aYr1Rcm1YXrpdrMRHWmcOlmefJMdfozS/hMpVC8WK3M75zVvFislvPV6UvF6bliucaFHzs5
LnonqpXr86LvuZkbxQIfMCq6L1ZnXs/XinsPaOcu9Pcl/Xb+9eJUteLPms3uxMFL1+cuTVz85sKOreyW
nlfbesZ8bhfyN54r5WvfrlSvcakZU/sr4x97TnY2P3nt8Qdl/qjuSdlnqB9mpqa+nPUDao81xJegfOge
EXaPCLtHhN0jwu4RYfeIsHtE2D0i7B4RflmOCP1fhn9ZqvzuQVn3oKx7UNY9KPstOyg76ex1UGZ/9bEH
ZSMjI/Tj7tNJANArqwBfF/f6xOW3qI8X9w3/Vt8d7//r5df61Kx4PvfjycJ3Dz64K4u7c/540jlpxxEc
tN9dg27rtm7rtm7rts+r+fuQxG53y7sX0fdob8EHcJvd9W7fPhstP6ch1n43XVHU7e3t7d9G/WXZ2Egw
G5ANBZ8ldSMHAN8HsvE1APCI4r4J765bq/Jt2IZl1EID7T0Z1I2L/O68jf0SIS4A2B4hjkzUo57yz/XZ
rLfmQqPxY2isW7riyrrqetY7nif/yZJHYu62qhwH8F4mzyquRzR3W1Hob7bfvLytqBTAOyGfUv0xHxJT
cWVTdVVY3qQKgQDc2qQqAQ1ubYYVAkF8VgioYWoHotTGccTk/SRMgOJzmEAYljflMKdHqEZ1O5S0HIrP
JvJZZpBGdFszdSccMhzkSSDTRPqQqTuB2f6mHBFzUpyDy0SFLJpp2JqQh5qGraMMJgFDyBRBGOWyBKJc
RoQoYwD7M2k7NNDroIz4rGUsO9JjOTRkOFqY2iHTcOSIassoh0mAhBXbU1U2HoIEUHcIcF3wGeWTggRk
uMWgynjoDsovBQgQ7A8QCJjUVgOKo4U0hwQVOxRSHQVubVoq2mt500L7op0UAjIx3EBIs3FOxCsRzZZC
eg7CNKfCQFM1fyIFoa9pUaQ50PR0jLKfPpR1lAea3pLEni0D5fSf/471qwBNja1FprmE9oT+pqcSCEGm
uagQkELaaSpwAP/wEHnL2t060tGoz/8vJIverQfEz7J2p66FNfC0e3UFoBkyCcgRFUhIAe1QY83VthqH
hZ+SCYX5IvNN9Dtcv0MACq7bIQBPpTmU1dPQptBEP5LRbwIE5IBmywFlnAA0SYjblkjc9ki7qGj2A4Xa
ELxfl6G3ievD7IC6zHrLPl0goIBHVNfVtlbMVfm2rNKjGCOe1h5Xpq65sk5dz/TjKuVuU43FlfKs5npE
d7c1jceVRllcBU5Rf8yHiqm5AZO6zIe1XZ9mvqyhzW8xPPp0MIl+vbypxHm/YhLQ8Vn4dMDk9Mxnk5Yd
ziQc9P1AHPksM6hHLZvGLScSiTnIU4FME+nDccsJzvY3MQ7YnAaPB5xbF7LQeMymQh49HrMNlCFOwBIy
mQiTXJZgksuIEGUMYv9Avx0+dID5PT7TgYRt9iUcPRJzME7D8ZiDuQPjG2VVTM32KGXjMfZRd1ybkMgl
tCOXaIyH5aD8cogwf0EYjBs2xhSN6I4S1uxwhDoYRyblMWpqPAfI6E8k5gYjuo1zIh7zkxyxcsQ0choM
NLX4T6QQ9DVNA2kOND1LxJWF8uzGlRlDOVviymL+fhpjg8Z5LC2hTTGOKAFZvVsPY3xpPP4wZkzjbl0X
/UE/jtQ7dRZzFgHVVMFT79WRp57ksRCOEwhEKShBrS2uTF115QnM5yy3v4y5HHOLfBBYbkHoaTqPK4q2
g6YS4TZUggSUILWVoDreGW8+j0WV2g9U3Sah+3UFelkMMTtoPK58OjWogUcoiysd44oaRzFGPH03rt6H
xrqO8uqI9+Mq7W7rKosr+VnV9YjBdGFxpeosrtRTuj/mQ7ZXmbqLvuLvU1THnH9rM6xy30R8KGrZWo9l
4zg5yfvlKPdhhLhnqGKPQBjqidl6X8qJ4HOS+yDCSDxmh5MxxzATDvKUIdNEej0Zc7TZ/qYqYpZafN/B
uSNClnAyYYeFPJFkwjZRhiSBqJAJ40vu4bJoPVxGhCijhv0HD9j6E1kHZcTn8MGUbQ2knIiZcMJRy9aT
CUeN67aKcqCfRFXb03U2Htc4JNYYdcFnti9G+F6DMMR4xNgeHQjytUeIOSYUVJ2w2A91U2dxpes8rnRh
d1ltjyvEY1wFIlZO6YgrvSWuAhF6+lGxpccIyPRuvTO+/LjRjbt1P8524obeqasm5XHH1j/TXFQJeNa9
Oo3r4NF7dTli7cTME74PdsQMyzUZYHs8Qo8aPGZ0zEPQxNwbEPYJBHU74MdMRMRMSOSrDMaMbj9QDVsJ
368HoLephHnM6LQlZrDOCVqA8YJ7mHIQQCZ042kA+BUxXBJUc4zfEYAHRHUDlI5HqY55szei6+lAZKk+
m72z9jw01onYA5GPlkGobrBjpYCSY30hwLyRI4o1LocU24iQ91juVQjELMX2LK2X15qqu2pYgxr8YHPJ
IGAoShosy5bjMUcJLNU9aWkNZUH5NFhad5WtlQDaG/WCv3mIPgVkacUlW28fgsb6bPattV9vN9aZfw0A
eNLymiKnXB0yzXCSwGq6d5DlxhT6/A82l9KcTzhAANTllQhkmlofxkd/MzSAfRc/igBE5uHddYBzDyMm
94Ngn6gXMgRcdevtX2y/ve6qWys/215ad8nWipvYWgHkDwRQzxBRXA00m8jUCRPFpQA2kYkDMP+REmR/
RfsRP1BbXAfw/kf/Pu37jyo9/v3nSfaXs/833n/89x609+77DbwnE7KB/c8AYe85zO9AWpQAbA0abK/4
S6mxLgNlPozvUVmEAR4b35fUjeeRv0Q2dJ+eqG5kNXQbiOJGQrCor6Zvy4S6EEo19DQsmkbgPU96Z43g
uxD8cA1jHtKZhgnCP3q3VlaBDKrCR4KEuKqu25JpOAqTa7k+S99ZQb/AjOHSrbcbUmM9pRN3FfRBiCTs
jG640gRxgxOG20MIlUzi9hgGDZoGqwGlYWD1F0KWVywCXiyVw3wW6uE1kCbj+1FsKKTrQ5qcsDWZjEsA
TSlJQEIecZ53kcciSdgPSMoORdO5UPR+HfNLKs3zi08TlglIETIkRaychLXZSZ7fEK5asUGPJHZyXNs8
WGfivhLXh9Kx2JBkJmzJ7JBFvKdJyNNM5DKJu3XMvQ8Iev0/PlRNAirJNBcJgVn45TKOR3qKMsXJkBQ3
uEwZAE3Wcwk9bVvpZVEr3GsMQ2M9tkpuxwgsApEbsk5cjxxmNub27XWtNBlfJbFBKU7Y3hoJAttTEeIc
JBYb8g4Mch2zfK8xnxD1bx8B2peyaV9iHPG0R+Blvg/TIMBiot9+kHjSlnqO5qQebmP56K6N6QCB+I4+
JLdIiN17yFuTScIlJOW6sNWwdvSQF+UWPbAfx7XqE5OJ26kHyu/C1gqrAfoAVtP9g5g/dcyH0VjOy96p
m09YYGH//xN+1EdAHzgwpGH+PIR1eqYZYbVOL4cZi/HXDhJIxtJ2tMPu8VVyO75jd8vVJiy3x7KoZva6
0bQ1vmoNDWrHLDsqZI0JWbVj1hAZGhryBp/kNj9MIAHQjB5BuLyZ6COQ6DtuJ/qOjiM+MUwghfgoAY8c
dXvg1mYK7X581H5w3La1E7mcdkLYPcftjvGSGiOQiVpsPu2YlVu0LGH3IZeQ4zt253r4drdcnAP7cVyr
TgnZcjt1QR3Q7iiTlSRg9KSGzBP9OSN5n/t5lvu5eYKAgX6e5X5uDfQzP0PfoD2JIdqTyrH9zyCwmOq3
jVSjbhDDXU2kBinmmwSB4IABJKnYco/qKEnNDvRQRzKX6h5JuLPZe2s/22qse4fv1KNHUoDvZuGneC0X
7sM9NMH9Ad/Hkukh3E/DA7hnZpoRVlP2cjiQ4jQZAnQ4MYT1nNoHvFY9RkDv688Zx+7Xke5Biuumox8R
vj+jbuEM8oCmgfySqaFwMpFbTKRsM9Goo44mIWCStKv3pEDrWaqjXFi3qpBhdTLWD6hTGPPuga0Vt39r
RQKeI3Dfx7yD9QPSmWkCWEf8cAvriPtr/7HdWMd6wRR1gie9v4b1ATnI6wP5EL6nZprKQayr+5uBQwQ0
OeZaONfg1go744hhnZFpvtXLabx+Pi/mXzjw/grSv3WA02O+cA9svf1P2411rDtMgCjWHVG0zyBhOTyK
9U72z9Z+IcZExBj6RIL82/bSOq8fuq3buq3buq3buq3buq3b9n/b+dac+M5eeOd3+LwFBHxX9PunXlkB
/3Nru4LwedHvf1eu9NSj55ss5cvT2df5X4tmx8ZGRkdGs0dPLFQnTxRv1IrVcr50olR6fW54vlqZLU7W
TnCC0ZPOyXxhbHLqymjeKY5+5cpoftQeLTrjJ+3c6OjYqdy47VyxcyefAjg/U75WrJ7Onj9/7pPwL5UK
n4b7x92jQGvePt2O1wT+6P9vx5sC/+8d+JTAK19rxw8IfH8H/rDANzrwx3w+uXb8mMC/2zHeEfgfdYz/
hsD/wVfb8ecE/r868L8r8PNn2/GvCPxfd+C/I/DZc+34ywJ/uAM/JfDnO/Bzvl4d+JrA/30H/o8E/s+/
0Y5f9O32LDzifFeFDzrG/6m/jh34O759nm7H3/PlP9OO/5EfjV9vx/+Vjx9tx/+SHd4Gd78rKdpP97i3
+i973FuFkXKlVoSRws3yws05GJkuXx+5ml+4CuJ/xNeqMFIr3qixp/zczCSMTFbm5orlGows3Jyr5a/A
yMLVhVqV/8QhnD07enlslAObgxwHX+HgqxycYoCPsPl4Pm6c/e/A2bNjnNEYZzTGGY1xRmOXx8Y5ECNP
McAH2mMMnGT/cyI+2Om8a3u4VJnMlzpu3LYj97qWe/ncay+dufDCM5/bfaxg693gPb7jufO7hQ76oMjh
pCOv+/D5lrwutXyvNN2Sl369vV3x6f287sPDHWJpHfNnBG/SsQ/4MNtBr3TAg+KONOnYd97toG/Pu7vt
SOu3XmHv78PuxWBY0Pr3tPf6bmugQ3/xGVcYFyw7whXmBf36HtP78HcedS8c+fEECg/J7j7d/4j1m2iV
vbV9i4PXHmO/V/agf03Q/+ox9P8dAAD//5QJv/NQWAAA
`,
	},

	"/relu.hsaco": {
		name:    "relu.hsaco",
		local:   "relu.hsaco",
		size:    13904,
		modtime: 1591217294,
		compressed: `
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
/afZ/xoAAP//bciEz1A2AAA=
`,
	},

	"/trans.hsaco": {
		name:    "trans.hsaco",
		local:   "trans.hsaco",
		size:    9656,
		modtime: 1591217294,
		compressed: `
H4sIAAAAAAAC/+xaTXPbxBt/JCuO4+T/B0605YAIh7QdJCtynDi50Lw0SadJmjbQUphOZyOtbTV6G2kd
YgbSF2inM3SGl+FePgCHHjm1OfABmJ576IHO8AXgiBlJu7bkWn2BdgpUv5n4sZ6XfX7PWqvsrvbi0ZVF
nuOOAEUO7gMXfHk1umaGH0cjeTjUVaEAR2AEipAHACHm1yv3uKQsUD1H49Lw9khSMj5B3EDsulf+yiVl
PC7gCiLV90gXkpLF8U8Zx+o79YDowhPExfkFOPmA6Hl4egisP3noEo/JW0NJKcTiCjT/7OpC6M5+mzfC
+yHSCzDYqY3pZlcXltbfj3z3A8Aw1SNLr2u2hCw9+Gv4SJLqtZ2qUqbtykMAReorSVLxNPZ8w7FnRIaP
xPF3REU8VzyOPRub/kxRFCVxDVm46yOK4nsesn3X8XExuNpoWZuOGXMa69iPbOljoc8KsutNVO82c8LF
9vyKOJ+wdviEPFTxXGid9eohkQB9yByzix1eLRcn7GM100Hk8FjHY8P4JBld7ZhmTaNuz/Q1nUZmEx83
bJ2Zl0xnE5lzzVoNe0mvgALzWiyr3dZ13dtwkYZPNpE502mia9c0ZomwgGuoaZL0uk80yctZuNMkbpOc
MXTSSO8AwybptU+k1z6RXvtcK1Sll30sXvZfLGsZG/UG+RfX9YxutGVD17Ed3SsnajUfkw8eQXBy4vnn
P/uC83/4AvKvOfaj7ovqE47yjNULY7XaNImx5Bn6RsvWZr3632Y47+h43XPczr/kYKqAvPoGrlvYJhH5
qkKNS57TdKlp0djBemRn5nXP2EYEpzskG6flM6Jn0DaueQ5LKoqdYbDWtDaW1k/5na4qV7qW0wkLM6yi
nUUTkTOOtxWRDttUK5NFWZaL6fO/YM52IJd/aB7Mxf4O0OmmSK+PTg9/x8emily8wc7kCTJkyPAfBUfH
Pxeu7nI9D4E+/twP8PVQsNZLPozW42v3aIXeXZsKQr7dbrf/ifXzwO8Fa9LPIb8XPBtzueLePgjWp1/e
Hobvb1+EG3egDdcC9gUofDtSKFwWrsGlgUvcFY7P7wIv7AIs3+MB+AKACvDTzzzw4MKNOzxcvTsk8PB/
QVD5HK/m+OufwYVL10D84vYBuB7aYRAgxxX3wr4dENSbQn6Uh6/uXhF4CHIP8IVdQRBUYTBfBVgP83Bw
9S43ADDI53c5DtSbwI9CEAM85AHUXI6vBvkB3HvRw/3ynexOz5AhQ4YMGTJkyJAhQ4aXG+xd863hSFIB
+6hkK/l99D08W/Uzv9/+aDuBvE8VnXf6I/3zLc/Pi5qJ7Lq4Hb1lFccVWZEV8WBJRwSVLmB7y7B96WPH
2/JdpOGS5lhuk2DJczRL8rAplWWlhHcI9mxklhqaJhGHlExz25Jcz7mANVKKEkwoU5XJ8uTE1Pg0rpaR
qut6BW1iBdVqeEpRptSypuvVinpIPLiJfKyLji0G9MqyIo9PT0yXpakKRtMVVaItiYdgxbC3sDcjrqws
PA/ipqk/Oe3H7usEv+6nryX1g1T/S4/+9XAzZLB7ToHifyn7yiDbDsEg6y3bb1kg1+2m3EB+A+hnoCce
yATvkPAKWYYGsuZYFrYJyH7LImgTZL/hEy/6FkmYm1POj4efamwT+vzC2bXZ1WPzz3LfazC29512XqKz
hwUP9+NwLIyNHyaV2PjhYudC2Lh6BQB+b7cdFs/GD5NiD61CT/79tG2+d7yNJPPwPfyZfJOe4+B7xjeT
w33vpy7G4mdqIP0cTloDEo3NMUXK+ZiBnvpZmknapNKTxqXxq1z/9Ey+G++7GPbeiuRO7Dkn9Pn9luLc
Y6jSc1JnH9N/J1PivxlNdkda/J8BAAD//5VIJrK4JQAA
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
