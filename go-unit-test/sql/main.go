package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// 批量创建课程
func main() {

	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for j := 0; j < 100; j++ {
	//			createClazz()
	//		}
	//	}()
	//}
	//wg.Wait()
	str := "/123"
	fmt.Println(strings.TrimPrefix(str, "/"))

}

// createClazz 创建课程
func createClazz() int {
	chapterID := createChapter()
	courseID := createCourse()
	// 请求

	// 示例 JSON 数据
	jsonData := `{"type":1,"title":"无本人-课程","cover_img":"","course_id":"","chapters":[{"id":310,"course_ids":["815911b8976411ef87be0242ac170017"]}],"description":"","category_id":"63f00abada8511e9a2dd0242ac180006","is_reminded":false,"required_privilege_type":1,"optional_privilege_type":2,"allow_accelerate":1,"pop_duration":0,"required_privileges":[{"id":"b0e48dd6ff3f11e8a178787b8aca2306","model":"staff","name":"cattyhuang(黄荣敏)","display_name":"cattyhuang(黄荣敏)","type":"staff","is_resigned":false}],"optional_privileges":[],"start_reminded_at":null,"end_reminded_at":null,"interval":null,"allow_comment":true,"credential":{"cer_id":"","title":"","enabled":false,"is_notify":false},"enable_credit":false,"is_chapter_hidden":false,"point_plus":{"enabled":false,"plus":0},"enable_ai_video_text":false,"enable_ai_clazz_assistance":false,"tags":[]}`
	var clazzReq ClazzReq
	err := json.Unmarshal([]byte(jsonData), &clazzReq)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	for i := range clazzReq.Chapters {
		clazzReq.Chapters[i].ID = chapterID
		clazzReq.Chapters[i].CourseIds = []string{courseID}
	}

	payloadBytes, err := json.Marshal(clazzReq)
	if err != nil {
		// handle err
	}

	requestbody := bytes.NewReader(payloadBytes)

	_, err = sendRequest("POST", "http://km.lx.net/api/v1/settings/classes", requestbody)

	if err != nil {
		return 0
	}

	return 1
}

func createCourse() string {
	data := CourseReq{
		MediaType: 9,
		Title:     "签名任务-批量",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}

	requestbody := bytes.NewReader(payloadBytes)

	respBody, err := sendRequest("POST", "http://km.lx.net/api/v1/settings/courses", requestbody)

	if err != nil {
		fmt.Println("createCourse error: ", err)
	}

	jsonData := JSONData{}
	json.Unmarshal(respBody, &jsonData)
	course := Course{}
	dataBytes, err := json.Marshal(jsonData.Data)
	json.Unmarshal(dataBytes, &course)
	return course.ID
}

// createChapter 创建章节
func createChapter() int {

	data := ChapterReq{
		Name: "第一章",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}

	requestbody := bytes.NewReader(payloadBytes)

	respBody, err := sendRequest("POST", "http://km.lx.net/api/v1/settings/class-chapters", requestbody)

	if err != nil {
		fmt.Println("createChapter error: ", err)
	}

	jsonData := JSONData{}
	json.Unmarshal(respBody, &jsonData)
	chapter := Chapter{}
	dataBytes, err := json.Marshal(jsonData.Data)
	json.Unmarshal(dataBytes, &chapter)
	return chapter.ID
}

type ChapterReq struct {
	Name string `json:"name"`
}

type CourseReq struct {
	MediaType int    `json:"media_type"`
	Title     string `json:"title"`
}

type JSONData struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
}

type Chapter struct {
	TeamID    string `json:"team_id"`
	Name      string `json:"name"`
	ClassID   string `json:"class_id"`
	Sequence  int    `json:"sequence"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	Model     string `json:"model"`
}

type Course struct {
	Title        string      `json:"title"`
	MediaType    int         `json:"media_type"`
	Summary      string      `json:"summary"`
	Downloadable int         `json:"downloadable"`
	TeamID       string      `json:"team_id"`
	CreatedBy    string      `json:"created_by"`
	SourceFrom   int         `json:"source_from"`
	ID           string      `json:"id"`
	UpdatedAt    string      `json:"updated_at"`
	CreatedAt    string      `json:"created_at"`
	Model        string      `json:"model"`
	CoverAsset   interface{} `json:"cover_asset"`
}

type ClazzReq struct {
	Type                    int           `json:"type"`
	Title                   string        `json:"title"`
	CoverImg                string        `json:"cover_img"`
	CourseID                string        `json:"course_id"`
	Chapters                []Chapters    `json:"chapters"`
	Description             string        `json:"description"`
	CategoryID              string        `json:"category_id"`
	IsReminded              bool          `json:"is_reminded"`
	RequiredPrivilegeType   int           `json:"required_privilege_type"`
	OptionalPrivilegeType   int           `json:"optional_privilege_type"`
	AllowAccelerate         int           `json:"allow_accelerate"`
	PopDuration             int           `json:"pop_duration"`
	RequiredPrivileges      []interface{} `json:"required_privileges"`
	OptionalPrivileges      []interface{} `json:"optional_privileges"`
	StartRemindedAt         interface{}   `json:"start_reminded_at"`
	EndRemindedAt           interface{}   `json:"end_reminded_at"`
	Interval                interface{}   `json:"interval"`
	AllowComment            bool          `json:"allow_comment"`
	Credential              Credential    `json:"credential"`
	EnableCredit            bool          `json:"enable_credit"`
	IsChapterHidden         bool          `json:"is_chapter_hidden"`
	PointPlus               PointPlus     `json:"point_plus"`
	EnableAiVideoText       bool          `json:"enable_ai_video_text"`
	EnableAiClazzAssistance bool          `json:"enable_ai_clazz_assistance"`
	Tags                    []interface{} `json:"tags"`
}
type Chapters struct {
	ID        int      `json:"id"`
	CourseIds []string `json:"course_ids"`
}
type Credential struct {
	CerID    string `json:"cer_id"`
	Title    string `json:"title"`
	Enabled  bool   `json:"enabled"`
	IsNotify bool   `json:"is_notify"`
}
type PointPlus struct {
	Enabled bool `json:"enabled"`
	Plus    int  `json:"plus"`
}

// sendRequest 发送HTTP请求并返回响应体和错误
func sendRequest(method, url string, body *bytes.Reader) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "company_server_type=lexiang; company_code=km; company_display_name=km3; certificate_guide_disabled=true; wishes_guide_disabled=true; x_host_key_access=d93727a919e9b9e4c3f007afddbf0ac02c38d09d_s; x-client-ssid=0613b00b:019269e07e97:0d69a1; token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9seC5uZXRcL2xvZ2luIiwiaWF0IjoxNzMwMjU2NTk5LCJleHAiOjE3MzI4NDg1OTksIm5iZiI6MTczMDI1NjU5OSwianRpIjoiMVFTT3g2MWFmdVV4UE5TNCIsInN1YiI6IjFkMzk3MGNlZjcwZTExZTViNGY3MDgwMDI3YzkzODI0IiwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsImNvbXBhbnlfaWQiOiIxYzFkNjY1YWY3MGUxMWU1OGEwMTA4MDAyN2M5MzgyNCIsInN0YWZmX3V1aWQiOiJhZjdjZDRiMmZmM2YxMWU4YTBkZDc4N2I4YWNhMjMwNiJ9.U_o2H4-CVTCZl99ki4oCpK-qXncHB6DPnhrrhxfTZxA; ti18nLng=zh-CN; XSRF-TOKEN=MBwGRc%252BLtBQovY6zRwu3iq5rJcmGKgE29jry8cm3VDDP5jBeLfWmaiv%252BJLT%252BoqBUZeg4dyxpGJnXl9JXZPo0j%252BGfPvW5idE1rEwd6nHz4fQ%253D")
	req.Header.Set("Origin", "http://km.lx.net")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "http://km.lx.net/classes/create?is_chapter=1&company_from=km")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	req.Header.Set("X-Auth-Type", "api")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Xsrf-Token", "MBwGRc%2BLtBQovY6zRwu3iq5rJcmGKgE29jry8cm3VDDP5jBeLfWmaiv%2BJLT%2BoqBUZeg4dyxpGJnXl9JXZPo0j%2BGfPvW5idE1rEwd6nHz4fQ%3D")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// 将 map[string]interface{} 转换为结构体
func mapToStruct(m map[string]interface{}, s interface{}) error {
	// 获取结构体的反射值
	structValue := reflect.ValueOf(s).Elem()

	// 遍历 map
	for key, value := range m {
		// 查找结构体中的字段
		field := structValue.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			// 将值转换为字段的类型并赋值
			val := reflect.ValueOf(value)
			if val.Type().ConvertibleTo(field.Type()) {
				field.Set(val.Convert(field.Type()))
			} else {
				return fmt.Errorf("cannot convert value of type %s to field %s of type %s", val.Type(), key, field.Type())
			}
		}
	}
	return nil
}
