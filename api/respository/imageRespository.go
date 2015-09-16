// this file has thes function below
// 1. access qiniu cloud storage
// 2. create user space to store image
// 3. create thumbnail
// 4. edit imagevia qiniu api
package respository

import (
	_ "log"
	"os"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

const (
	// access qiniu api access key
	ACCESS_KEY_KEY string = "Qiniu_Access_Key"
	SECURE_KEY_KEY string = "Qiniu_Secure_Key"
)

var (
	// access qiniu api secure key
	Access_Key string = ""
	Secure_Key string = ""
)

type ImageRespository struct{}

func init() {
	Access_Key = os.Getenv(ACCESS_KEY_KEY)
	Secure_Key = os.Getenv(SECURE_KEY_KEY)
	if Access_Key == "" || Secure_Key == "" {
		panic("can't get the Access_Key,Secure_Key.")
	}
}

// require the upload token
func (this *ImageRespository) GetUPToken() string {
	kodo.SetMac(Access_Key, Secure_Key)
	zone := 0
	c := kodo.New(zone, nil)
	bucket := "photoswall"
	policy := &kodo.PutPolicy{
		Scope:   bucket,
		Expires: 3600,
	}

	upToken := c.MakeUptoken(policy)

	return upToken
}

// upload image
func (this *ImageRespository) UploadImage(imgPath string) (bool, error) {
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	ctx := context.Background()
	err := uploader.PutFile(ctx, nil, this.GetUPToken(), "", imgPath, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// https://7xlpj4.com1.z0.glb.clouddn.com/bucket/photoswall
