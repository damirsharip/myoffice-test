package service

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) UrlHandler(file *os.File) {
	var wg sync.WaitGroup
	errorChan := make(chan error)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			if err := checkUrls(url); err != nil {
				errorChan <- err
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(errorChan)
	}()

	for err := range errorChan {
		fmt.Println(err)
	}
}

func checkUrls(link string) error {
	startTime := time.Now()

	response, err := http.Get(link)
	if err != nil {
		return fmt.Errorf("%s error while getting JSON", link)
	}
	defer response.Body.Close()

	diffTime := time.Since(startTime)

	fmt.Println(link, diffTime, response.ContentLength)

	return nil
}
