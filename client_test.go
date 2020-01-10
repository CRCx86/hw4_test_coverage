package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"strconv"
	"testing"
)

// код писать тут

type Users struct {
	XMLName xml.Name   `xml:"root"`
	Users   []*UserXML `xml:"row"`
}

type UserXML struct {
	Id         int    `xml:"id" json:"id"`
	GUID       string `xml:"guid"`
	IsActive   bool   `xml:"isActive"`
	Balance    string `xml:"balance"`
	Picture    string `xml:"picture"`
	Age        int    `xml:"age" json:"age"`
	EyeColor   string `xml:"eyeColor"`
	First_name string `xml:"first_name"`
	Last_name  string `xml:"last_name"`
	//Name          string `xml:"-" json:"name"`
	Company       string `xml:"company"`
	Email         string `xml:"email"`
	Phone         string `xml:"phone"`
	Address       string `xml:"address"`
	About         string `xml:"about" json:"about"`
	Registered    string `xml:"registered"`
	FavoriteFruit string `xml:"favoriteFruit"`
	Gender        string `xml:"gender" json:"gender"`
}

func TestSearchClient_FindUsers(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}

	sr := SearchRequest{
		Limit:      10,
		Offset:     0,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsIs,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()

}

func TestBadToken(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     0,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestOffset(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     -1,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestBadLimit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      -1,
		Offset:     1,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestLimitLess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     1,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestBadJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(BadJson))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     1,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func BadJson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(""))
}

func TestBadUnpack(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(BadUnpack))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     1,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func BadUnpack(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(""))
}

func TestLimitGreat(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      27,
		Offset:     1,
		Query:      "Boyd",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestTimeOut(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ReadUsers))

	sc := SearchClient{
		AccessToken: "123",
		URL:         "http://127.0.0.1:53663",
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     0,
		Query:      "",
		OrderField: "0",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestInternalBadRequestOrderField(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(InternalBadRequestOrderField))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     0,
		Query:      "",
		OrderField: "abc",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func InternalBadRequestOrderField(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	b := SearchErrorResponse{Error: "ErrorBadOrderField"}
	marshal, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Write(marshal)
}

func TestInternalBadRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(InternalBadRequest))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     0,
		Query:      "",
		OrderField: "abc",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func InternalBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	b := SearchErrorResponse{ErrorBadOrderField}
	marshal, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Write(marshal)
}

func TestServerError(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(InternalServerError))

	sc := SearchClient{
		AccessToken: "123",
		URL:         ts.URL,
	}
	sr := SearchRequest{
		Limit:      0,
		Offset:     0,
		Query:      "",
		OrderField: "abc",
		OrderBy:    OrderByAsc,
	}

	users, err := sc.FindUsers(sr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(users)

	ts.Close()
}

func TestError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SomeError))
	defer ts.Close()

	sc := SearchClient{
		AccessToken: "123",
		URL:         "unknow",
	}

	req := SearchRequest{}
	_, err := sc.FindUsers(req)
	if err == nil {
		t.Errorf("Error")
		return
	}
}

func SomeError(w http.ResponseWriter, r *http.Request){
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accesstoken") == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	users, err := ReadDummyUsers(r.FormValue("query"), r.FormValue("limit"), r.FormValue("offset"), r.FormValue("order_field"), r.FormValue("orderBy"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marshal, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(marshal)

}

func ReadDummyUsers(query string, limit string, offset string, order_field string, orderBy string) ([]UserXML, error) {

	xmlFile, err := os.Open("dataset.xml")

	if err != nil {
		return nil, err
	}

	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var v Users
	err = xml.Unmarshal(xmlData, &v)
	if err != nil {
		return nil, err
	}

	users := make([]UserXML, 0)

	if limit != "0" {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		if lim > 1 {
			for i := 0; i < len(v.Users); i++ {
				users = append(users, *v.Users[i])
			}
			return users[:lim], nil
		}

	}

	if query != "" {
		for i := 0; i < len(v.Users); i++ {
			if v.Users[i].First_name == query {
				users = append(users, *v.Users[i])
			}
		}
	}

	count, err := strconv.Atoi(offset)
	if offset != "" {
		for i := 0; i < len(v.Users); i++ {
			if i < count {
				continue
			}
			users = append(users, *v.Users[i])
		}
	}

	if orderBy != "" {
		if order_field == "1" {
			sort.SliceStable(users, func(i, j int) bool {
				return users[i].Id < users[j].Id
			})
		} else if order_field == "-1" {
			sort.SliceStable(users, func(i, j int) bool {
				return users[i].Id > users[j].Id
			})
		}
	}

	if order_field != "" {
		atoi, err := strconv.Atoi(order_field)
		if err != nil {
			return nil, err
		}
		fmt.Println(atoi)
	}

	return users, nil
}
