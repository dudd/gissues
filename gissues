/**********************************
*** Author : dudd
*** Date   : 2018-09-27
*** Desc   : A tool for manage github issues
***********************************/

package gissue

import (
	"net/url"
	"strings"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bytes"
		"time"
		"os"
	"os/exec"
	"runtime"
	"bufio"
	"strconv"
)


var commands = map[string]string {
	"windows": "cmd /c start",
	"darwin":  "open",
	//"linux":   "xdg-open",
	"linux":   "gedit",

}


func initFile(fname string, content interface{}) error{
	f, err := os.Create(fname)
	if err != nil {
		fmt.Printf("Create %s err, %s\n", fname, err)
		return err
	}

	b, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		fmt.Printf("Marshal content err, %s\n", err)
		return err
	}

	f.WriteString(string(b))
	f.Close()
	return err
}


func readFile(fname string, body interface{}) (error){
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Open %s err, %s\n", fname, err)
		return err
	}

	err = json.Unmarshal(b, body)
	if err != nil {
		fmt.Printf("Unmarshal err, %s\n", err)
		return err
	}

	return err
}


func getInput(issueNums map[int]int) (int){

	input := 0
	for ;; {
		fmt.Printf("Please input issue number : ")
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input, err = strconv.Atoi(strings.Split(s, "\n")[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		if issueNums[input] == 0 {
			fmt.Println("Can't find this issues.")
			continue
		}
		break
	}

	return input
}


func editIssue(fname string, body interface{}) error {

	// Open text editor to edit request body.
	err := initFile(fname, body)
	if err != nil {return err}
	cmd := exec.Command(commands[runtime.GOOS], fname)
	cmd.Run()

	// Read body content from file and send request.
	err = readFile(fname, body)
	return err
}


// My http client, just for exercise. You should use gin, beggo etc.
func MyHttp(method, reqUrl string, result, body interface{}) (error){
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {return err}

	req.Header.Add("Accept", headerAccept)

	if req.Body == nil && body != nil {
		byts, err := json.Marshal(body)
		if err !=nil {return err}
		req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		req.ContentLength = int64(len(byts))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "token " + token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err !=nil {return err}

	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, result)
	fmt.Println(string(b))

	//err = json.NewDecoder(resp.Body).Decode(result)
	if err !=nil {return err}

	resp.Body.Close()
	return err
}


func SearchIssues(args []string) {
	var result *SearchResult
	reqUrl := searchIssue + "?q=" + url.QueryEscape(strings.Join(args, " "))

	err := MyHttp("GET",  reqUrl, &result, nil)
	if err != nil {
		log.Printf("SearchIssues error, %s\n", err)
	}

	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))

	fmt.Printf("Total_count: %d\n\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("No.%-5d title:%s\turl:%s\n", item.Number, item.Title, item.Url)
	}
}


func ListRepoIssus(owner, repo string) (ListResult, map[int]int){
	var result ListResult
	issueNums := make(map[int]int)
	reqUrl := fmt.Sprintf(repoIssues, owner, repo)

	err := MyHttp("GET", reqUrl, &result, nil)
	if err != nil {
		log.Printf("ListRepoIssus error, %s\n", err)
		return nil, nil
	}

	for _, item := range result {
		fmt.Printf("No.%-5d title:%-20s\t%s\n", item.Number, item.Title, item.Url)
		issueNums[int(item.Number)] ++
	}

	//conv, err := json.MarshalIndent(result, "", "  ")
	//fmt.Println(string(conv))
	return result, issueNums
}


func GetSingleIssue(owner, repo string, num int) ([]byte) {
	var result Issue
	reqUrl := fmt.Sprintf(singleIssue, owner, repo, num)

	err := MyHttp("GET", reqUrl, &result, nil)
	if err != nil {
		log.Printf("GetSingleIssue error, %s\n", err)
	}

	conv, err := json.MarshalIndent(result, "", "  ")
	//fmt.Println(string(conv))
	return conv
}


func CreateIssue(owner, repo string){
	var result Issue
	body := IssueReq{}
	fileName := "/tmp/tmp-" + time.Now().String()
	reqUrl := fmt.Sprintf(repoIssues, owner, repo)

	err := editIssue(fileName, &body)
	if err != nil {
		fmt.Printf("CreateIssue error when call editIssue, %s\n", err)
		return
	}

	err = MyHttp("POST", reqUrl, &result, body)
	if err != nil {
		log.Printf("CreateIssue error, %s\n", err)
	}

	conv, err := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(conv))
}


func EditIssue(owner, repo string){
	var result Issue
	body := IssueReq{}
	fileName := "/tmp/tmp-" + time.Now().String()
	_, issueNums := ListRepoIssus(owner, repo)
	input := getInput(issueNums)
	issue := GetSingleIssue(owner, repo, input)
	reqUrl := fmt.Sprintf(singleIssue, owner, repo, input)

	err := json.Unmarshal(issue, &body)
	if err != nil {
		fmt.Printf("EditIssue error %s", err)
		return
	}

	err = editIssue(fileName, &body)
	if err != nil {
		fmt.Printf("EditIssue error when call editIssue, %s\n", err)
		return
	}

	body.Milestone, body.Assignee = nil, nil
	err = MyHttp("PATCH", reqUrl, &result, body)
	if err != nil {
		log.Printf("EditIssue error, %s\n", err)
	}
	conv, err := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(conv))
}


func main(){
	EditIssue("dudd", "github-issues-tools")
	//CreateIssue("dudd", "github-issues-tools")
	//GetSingleIssue("dudd", "github-issues-tools", 1)
	//SearchIssues(os.Args[1:])
	//ListRepoIssus("dudd", "github-issues-tools")
}
