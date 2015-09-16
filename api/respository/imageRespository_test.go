package respository

import "testing"

// test get qiniu up token
func TestGetUPToken(t *testing.T) {
	imgRes := new(ImageRespository)
	upToken := imgRes.GetUPToken()
	if upToken == "" {
		t.Error("can't get up token.")
	} else {
		t.Log("the up token is", upToken)
	}

}

// test qiniu upload image file func
func TestUploadFile(t *testing.T) {
	imgRes := new(ImageRespository)
	_, err := imgRes.UploadImage(`D:\GoExercise\src\photoswall\fiddle\1.jpg`)
	if err != nil {
		t.Error("upload file fail.", err.Error())
	} else {
		t.Log("upload file successfully!")
	}
}
