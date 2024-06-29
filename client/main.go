package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

// Response is the structure of the server's response
type Response struct {
    Response string `json:"response"`
}

func sendMessage(message string) (string, error) {
    url := "http://load-balancer:80/message"
    requestBody, err := json.Marshal(map[string]string{"message": message})
    if err != nil {
        return "", err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var serverResponse Response
    if err := json.NewDecoder(resp.Body).Decode(&serverResponse); err != nil {
        return "", err
    }

    return serverResponse.Response, nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter a message:\n    ")
        message, _ := reader.ReadString('\n')
        message = strings.TrimSpace(message)
        if message == "" {
            continue
        }

        response, err := sendMessage(message)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        fmt.Println("Server response:\n    ", response)
        time.Sleep(1 * time.Second) 
    }
}
