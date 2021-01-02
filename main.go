/* Golang Script to get the metadata of EC2-IAM Credentials */
package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "os"
import "time"

const metadata_url = "http://169.254.169.254/latest/meta-data/iam/security-credentials/"
const filename = "metadata.txt"

type metacreds struct {    Token string `json:"Token"`
    SecretAccessKey string `json:"SecretAccessKey"`
    AccessKeyId string `json:"AccessKeyId"`
}

func getcreds(body []byte) (*metacreds, error) {
    var s = new(metacreds)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

func main() {
currentTime := time.Now()
resp, err := http.Get(metadata_url)
if err != nil {
        // handle err
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)

role_name := string(body)
resp1, err := http.Get(metadata_url+role_name)
defer resp1.Body.Close()
creds, err := ioutil.ReadAll(resp1.Body)
s, err := getcreds([]byte(creds))
_, ferr := os.Stat(filename)
    if !(os.IsNotExist(ferr)) {
	os.Remove(filename)
	}

file, err := os.Create(filename)
    if err != nil {
            return
        }
 defer file.Close()
 file.WriteString("## File Generated Date Time: "+currentTime.Format("2006-01-02-15h-04m-05s")+"\n")
 file.WriteString("export AWS_ACCESS_KEY_ID="+s.AccessKeyId+"\n")
 file.WriteString("export AWS_SECRET_ACCESS_KEY="+s.SecretAccessKey+"\n")
 file.WriteString("export AWS_SESSION_TOKEN="+s.Token+"\n")
 fmt.Println("FileName: metadata.txt is generated")
}
