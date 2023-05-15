package playWright

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

// DomInterface 돔 인터페이스
type DomInterface interface {
	Type(selector string, text string, options ...playwright.PageTypeOptions) error
	WaitForSelector(selector string, options ...playwright.PageWaitForSelectorOptions) (playwright.ElementHandle, error)
	QuerySelectorAll(selector string) ([]playwright.ElementHandle, error)
}

// ActionOption 액션 옵션
type ActionOption struct {
	NotFoundPass bool
}

// GoTo 페이지 이동
func GoTo(pageUrl *string, page playwright.Page) error {
	if _, err := page.Goto(*pageUrl, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		err = fmt.Errorf("페이지 전환을 실패하였습니다 : %v", err)
		log.Println(err)
		return err
	}
	return nil
}

// SendText 텍스트 입력
func SendText(selector, value string, dom DomInterface) error {
	noWaitAfter := false
	if err := dom.Type(
		selector,
		value,
		playwright.PageTypeOptions{NoWaitAfter: &noWaitAfter},
	); err != nil {
		return err
	}
	return nil
}

// Click 클릭
func Click(selector string, dom DomInterface, options ...ActionOption) error {
	btn, err := dom.WaitForSelector(selector)
	if err != nil {
		// 엘리먼트 존재 여부 확인
		elements, err := dom.QuerySelectorAll(selector)
		if len(elements) < 1 && len(options) == 1 && options[0].NotFoundPass {
			return nil
		}
		return err
	}
	if err := btn.Click(); err != nil {
		return err
	}
	return nil
}

// Screenshot 스크린샷 촬영
func Screenshot(page playwright.Page) {
	if _, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("screenshot.png"),
	}); err != nil {
		log.Println("스크린샷 촬영 중 오류 발생, 오류 내용 : " + err.Error())
	}
	log.Println("스크린샷 촬영 완료")
}
