package playWright

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

// App 어플리케이션
func App(Function func(playwright.Page, *[]map[string]string) error, reqDataset *[]map[string]string) error {
	pw, err := playwright.Run()
	if err != nil {
		err = fmt.Errorf("어플리케이션 실행간 오류가 발생했습니다 : %v", err)
		log.Println(err)
		return err
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		err = fmt.Errorf("브라우저 실행간 오류가 발생했습니다 : %v", err)
		log.Println(err)
		return err
	}
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
	page, err := browser.NewPage(playwright.BrowserNewContextOptions{UserAgent: &userAgent})
	if err != nil {
		err = fmt.Errorf("페이지 이동간 오류가 발생했습니다 : %v", err)
		log.Println(err)
		return err
	}
	// 동작
	if err := Function(page, reqDataset); err != nil {
		err = fmt.Errorf("동작 진행간 오류가 발생했습니다 : %v", err)
		log.Println(err)
	}
	if err = browser.Close(); err != nil {
		err = fmt.Errorf("브라우저 종료간 오류가 발생했습니다 : %v", err)
		log.Println(err)
		return err
	}
	if err = pw.Stop(); err != nil {
		err = fmt.Errorf("어플리케이션 종료간 오류가 발생했습니다 : %v", err)
		log.Println(err)
		return err
	}
	return nil
}
