
package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
   "strings"
   "sync"
)

type DictReq struct {
   TransType string `json:"trans_type"`
   Source string `json:"source"`
}

type AutoGenerated struct {
   Rc int `json:"rc"`
   Wiki struct {
      KnownInLaguages int `json:"known_in_laguages"`
      Description struct {
         Source string `json:"source"`
         Target interface{} `json:"target"`
      } `json:"description"`
      ID string `json:"id"`
      Item struct {
         Source string `json:"source"`
         Target string `json:"target"`
      } `json:"item"`
      ImageURL string `json:"image_url"`
      IsSubject string `json:"is_subject"`
      Sitelink string `json:"sitelink"`
   } `json:"wiki"`
   Dictionary struct {
      Prons struct {
         EnUs string `json:"en-us"`
         En string `json:"en"`
      } `json:"prons"`
      Explanations []string `json:"explanations"`
      Synonym []string `json:"synonym"`
      Antonym []string `json:"antonym"`
      WqxExample [][]string `json:"wqx_example"`
      Entry string `json:"entry"`
      Type string `json:"type"`
      Related []interface{} `json:"related"`
      Source string `json:"source"`
   } `json:"dictionary"`
}

func translate(word string,wg *sync.WaitGroup) {
   defer wg.Done()
   url := "https://api.interpreter.caiyunai.com/v1/dict"
   method := "POST"
   request := DictReq{TransType: "en2zh",Source: word}
   buf, err := json.Marshal(request)
   if err != nil{
      fmt.Println(err)
   }
   data := bytes.NewReader(buf)
   client := &http.Client {}
   req, err := http.NewRequest(method, url, data)
   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Accept", "application/json, text/plain, */*")
   req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
   req.Header.Add("Connection", "keep-alive")
   req.Header.Add("Content-Type", "application/json;charset=UTF-8")
   req.Header.Add("Origin", "https://fanyi.caiyunapp.com")
   req.Header.Add("Referer", "https://fanyi.caiyunapp.com/")
   req.Header.Add("Sec-Fetch-Dest", "empty")
   req.Header.Add("Sec-Fetch-Mode", "cors")
   req.Header.Add("Sec-Fetch-Site", "cross-site")
   req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
   req.Header.Add("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
   req.Header.Add("app-name", "xy")
   req.Header.Add("device-id", "")
   req.Header.Add("os-type", "web")
   req.Header.Add("os-version", "")
   // req.Header.Add("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"")
   req.Header.Add("sec-ch-ua-mobile", "?0")
   req.Header.Add("sec-ch-ua-platform", "Windows")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   if res.StatusCode != 200{
      fmt.Println("bad StatusCode:", res.StatusCode, "body", string(body))
      return
   }
   var dictRes AutoGenerated
   err = json.Unmarshal(body,&dictRes)
   if err != nil{
      fmt.Println(err)
      return
   }
   fmt.Println("彩云小译")
   fmt.Println(word, "UK:", dictRes.Dictionary.Prons.En, "US:", dictRes.Dictionary.Prons.EnUs)
   for _, item := range dictRes.Dictionary.Explanations {
      fmt.Println(item)
   }
}


type AutoGenerated2 struct {
   SessionUUID string `json:"sessionUuid"`
   Translate struct {
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
      SessionUUID string `json:"sessionUuid"`
      Source string `json:"source"`
      Target string `json:"target"`
      Records []struct {
         SourceText string `json:"sourceText"`
         TargetText string `json:"targetText"`
         TraceID string `json:"traceId"`
      } `json:"records"`
      Full bool `json:"full"`
      Options struct {
      } `json:"options"`
   } `json:"translate"`
   Dict struct {
      Word string `json:"word"`
      DetailURL string `json:"detailUrl"`
      Abstract []struct {
         Ps string `json:"ps"`
         Explanation []string `json:"explanation"`
      } `json:"abstract"`
      Data []struct {
         EnHash string `json:"en_hash"`
      } `json:"data"`
      Type string `json:"type"`
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
   } `json:"dict"`
   Suggest struct {
      Data []struct {
         ID string `json:"id"`
         Word string `json:"word"`
         Abstract []struct {
            Exp []string `json:"exp"`
            Ps string `json:"ps"`
         } `json:"abstract"`
         SuggestTranslation string `json:"suggest_translation"`
         PhJSON string `json:"ph_json"`
         TransformPl string `json:"transform_pl"`
         TransformIs interface{} `json:"transform_is"`
         TransformWere interface{} `json:"transform_were"`
         TransformBeen interface{} `json:"transform_been"`
         TransformBeing interface{} `json:"transform_being"`
         TransformEr interface{} `json:"transform_er"`
         TransformEst interface{} `json:"transform_est"`
         Groups string `json:"groups"`
         AmEForm interface{} `json:"AmE_form"`
         BrEForm interface{} `json:"BrE_form"`
         ExamplesJSON string `json:"examples_json"`
         WordOfPhrase interface{} `json:"word_of_phrase"`
         Frq int `json:"frq"`
         Valid int `json:"valid"`
         TransformSing interface{} `json:"transform_sing"`
         PrxPhAmE string `json:"prx_ph_AmE"`
         PrxPhBrE string `json:"prx_ph_BrE"`
      } `json:"data"`
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
   } `json:"suggest"`
   ErrCode int `json:"errCode"`
   ErrMsg string `json:"errMsg"`
}

func translate2(word string,wg *sync.WaitGroup) {
   defer wg.Done()
   url := "https://fanyi.qq.com/api/translate"
   method := "POST"
   payload := strings.NewReader(`source=auto&target=zh&sourceText=` + word +`&qtv=5a3d5ab458f88908
&qtk=9PfIShCScN2DELc%2FuDCt5%2BeP86QHhPu6tSZB6ITACggmPKdhkSh3bwyi6LJqKiEfUhOTcngbToukd5UTNQQ66RbIRRmWQ%2FUaIE1ahXL0vjPQLok3HcfmdvTEZbS%2Fi3f512lBudDDZVAvswD7QZ2B3Q%3D%3D
&ticket=&randstr=&sessionUuid=translate_uuid1652253538206`)
   client := &http.Client {}
   req, err := http.NewRequest(method, url, payload)
   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
   req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
   req.Header.Add("Connection", "keep-alive")
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
   req.Header.Add("Cookie", "pgv_pvid=5948909680; pac_uid=0_935f2dad4e09f; tvfe_boss_uuid=8b4f775e020c647a; eas_sid=21M6a4h7p9P4c8m8J668T1x2z2; RK=oksQmnU1XX; ptcz=3f7da84564a30539d2db26c20ab25bb3ba481b2a2eeb87318fd1dc0755ef6aa3; ptui_loginuin=1260828355@qq.com; fy_guid=366184b4-5d6d-4da7-97d8-c32c19b64fa8; ADHOC_MEMBERSHIP_CLIENT_ID1.0=6ac9805b-e21f-b7e7-8b93-cbda95f4b43a; openCount=1; qtv=5a3d5ab458f88908; qtk=9PfIShCScN2DELc/uDCt5+eP86QHhPu6tSZB6ITACggmPKdhkSh3bwyi6LJqKiEfUhOTcngbToukd5UTNQQ66RbIRRmWQ/UaIE1ahXL0vjPQLok3HcfmdvTEZbS/i3f512lBudDDZVAvswD7QZ2B3Q==; gr_user_id=5b66730c-2ffc-489c-b366-3c6c404734ba; 8507d3409e6fad23_gr_session_id=e72f0b98-87b0-44b7-a42a-a2d69b60f6e7; 8c66aca9f0d1ff2e_gr_session_id=6132b52a-e285-41cc-84bb-6bc05d3deb73; 8c66aca9f0d1ff2e_gr_session_id_6132b52a-e285-41cc-84bb-6bc05d3deb73=true")
   req.Header.Add("Origin", "https://fanyi.qq.com")
   req.Header.Add("Referer", "https://fanyi.qq.com/")
   req.Header.Add("Sec-Fetch-Dest", "empty")
   req.Header.Add("Sec-Fetch-Mode", "cors")
   req.Header.Add("Sec-Fetch-Site", "same-origin")
   req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
   req.Header.Add("X-Requested-With", "XMLHttpRequest")
   // req.Header.Add("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"")
   req.Header.Add("sec-ch-ua-mobile", "?0")
   req.Header.Add("sec-ch-ua-platform", "Windows")
   req.Header.Add("uc", "GLYdIz3WjseABEHSwmZpKBGZ9Afq6AxIJcPpg7ZK1j0=")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   //fmt.Println(string(body))
   var dictRes AutoGenerated2
   err = json.Unmarshal(body,&dictRes)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("腾讯翻译君")
   for _, item := range dictRes.Translate.Records {
      fmt.Println(item.TargetText)
   }
}

func main(){
   var word string
   fmt.Scanln(&word)
   wg := sync.WaitGroup{}
   wg.Add(2)
   translate(word,&wg)
   translate2(word,&wg)
}
