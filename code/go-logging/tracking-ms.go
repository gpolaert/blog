package main

import (
	"net/http"
	"io"
	"github.com/Sirupsen/logrus"
	"math/rand"
	"fmt"
	"io/ioutil"
)

var logger *logrus.Entry

func init() {
	logrus.SetFormatter( &logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logger = logrus.WithField("appname", "go-logging")
}

const aResponseMessage = "hello world from [microservive %d]\n\tsession: %s,\n\ttrack: %s,\n\tparent: %s\n-------\n"

func helloMicroService1(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}

	// This service is responsible to received all incoming user requests
	// So, we are checking if it's a new user session or a another call from
	// an existing session
	session := r.Header.Get("x-session")
	if ( session == "") {
		session = generateSessionId()
		// log something for the new session
	}

	// Track Id is unique per request, so in each case we generate one
	track := generateTrackId()


	// Call your 2nd microservice, add the session/track
	reqService2, _ := http.NewRequest("GET", "http://localhost:8082/", nil)
	reqService2.Header.Add("x-session", session)
	reqService2.Header.Add("x-track", track)
	resService2, _ := client.Do(reqService2)

	// Read the response
	m, _ := ioutil.ReadAll(resService2.Body)

	// Add meta to the response
	w.Header().Set("x-session", session)
	w.Header().Set("x-track", track)

	logger.WithField("session", session).WithField("track", track).Debugf("hello from ms 1")

	// Write the response body
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf(aResponseMessage, 1, session, track, "") + string(m))

}

func helloMicroService2(w http.ResponseWriter, r *http.Request) {

	// Like for the microservice, we check the session and generate a new track
	session := r.Header.Get("x-session")
	track := generateTrackId()

	// This time, we check if a track id is already set in the request,
	// if yes, it becomes the parent track
	parent := r.Header.Get("x-track")
	if (session == "") {
		w.Header().Set("x-parent", parent)
	}

	// Add meta to the response
	w.Header().Set("x-session", session)
	w.Header().Set("x-track", track)
	if (parent == "") {
		w.Header().Set("x-parent", track)
	}

	logger.WithField("session", session).WithField("parent", parent).WithField("track", track).Debugf("hello from ms 2")


	// Write the response body
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf(aResponseMessage, 2, session, track, parent))

}

func main() {

	// register the service 1 endpoints
	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/", helloMicroService1)

	// register the service 2 endpoints
	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/", helloMicroService2)

	// start servers
	go func() {
		http.ListenAndServe("localhost:8081", serverMuxA)
	}()
	http.ListenAndServe("localhost:8082", serverMuxB)



	// OUTPUT
	// hello world [microservive 1]
	//session: GRxnyDNY,
	//	track: uRKWsmZn,
	//	parent:
	//-------
	//	hello world [microservive 2]
	//session: GRxnyDNY,
	//	track: OYgSLPSs,
	//	parent: uRKWsmZn
	//-------

	// {"appname":"go-logging","level":"debug","msg":"hello from ms 2","parent":"UzWHRihF","session":"eUBrVfdw","time":"2017-03-02T15:29:26+01:00","track":"DPRHBMuE"}
	// {"appname":"go-logging","level":"debug","msg":"hello from ms 1","session":"eUBrVfdw","time":"2017-03-02T15:29:26+01:00","track":"UzWHRihF"}

}

func generateSessionId() string {
	return RandStringBytesRmndr(8)
}
func generateTrackId() string {
	return RandStringBytesRmndr(8)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}
