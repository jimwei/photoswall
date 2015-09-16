// this file has thes function below
// 1. access qiniu cloud storage
// 2. create user space to store image
// 3. create thumbnail
// 4. edit imagevia qiniu api
package respository

import "os"

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

func init() {
	Access_Key = os.Getenv(ACCESS_KEY_KEY)
	Secure_Key = os.Getenv(SECURE_KEY_KEY)
	if Access_Key == "" || Secure_Key == "" {
		panic("can't get the Access_Key,Secure_Key.")
	}
}

type ImageRespository struct{}
