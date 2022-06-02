package services

import "dropzone-go/utils"

func PrettyData(upload_type, file_name, date_format string) map[string]interface{} {
	r := make(map[string]interface{})
	if upload_type == "image" {
		r = imagedata(file_name, date_format)
	} else if upload_type == "video" {
		r = videodata(file_name, date_format)
	} else if upload_type == "file" {
		r = filedata(file_name, date_format)
	}
	return r
}

func filedata(file_name, date_format string) map[string]interface{} {
	r := make(map[string]interface{})
	r["fullLink"] = utils.GetMinioHost() + "/" + utils.GetMinioBacket() + "/files/" + date_format + "/" + file_name
	return r
}

func videodata(file_name, date_format string) map[string]interface{} {
	r := make(map[string]interface{})
	r["fullLink"] = utils.GetMinioHost() + "/" + utils.GetMinioBacket() + "/video/" + date_format + "/" + file_name
	return r
}

func imagedata(file_name, date_format string) map[string]interface{} {
	r := make(map[string]interface{})
	oSize := scalePic("0", file_name, date_format)
	mSize := scalePic("0.7", file_name, date_format)
	sSize := scalePic("0.5", file_name, date_format)
	r["oSize"] = oSize
	r["mSize"] = mSize
	r["sSize"] = sSize
	return r
}

func scalePic(scale, file_name, date_format string) map[string]interface{} {
	r := make(map[string]interface{})
	proxyHost := utils.GetProxyHost()
	uri := "x" + scale + "/" + utils.GetMinioHost() + "/" + utils.GetMinioBacket() + "/images/" + date_format + "/" + file_name
	fullLink := proxyHost + "/" + uri
	r["scale_uri"] = uri
	r["fullLink"] = fullLink
	return r
}
