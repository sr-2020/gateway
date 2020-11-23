package usecases

import (
	"github.com/mxschmitt/playwright-go"
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"log"
	"testing"
)

func _TestCheck(t *testing.T) {
	convey.Convey("Check authorization", t, func() {
		cfg := config.LoadConfig()

		pw, err := playwright.Run()
		if err != nil {
			log.Fatalf("could not launch playwright: %v", err)
		}
		browser, err := pw.Chromium.Launch()
		if err != nil {
			log.Fatalf("could not launch Chromium: %v", err)
		}
		page, err := browser.NewPage()
		if err != nil {
			log.Fatalf("could not create page: %v", err)
		}

		_, err = page.Goto("https://rc-web.evarun.ru", playwright.PageGotoOptions{})
		if err != nil {
			log.Fatalf("could not goto: %v", err)
		}

		if err = page.Type("#mat-input-0", cfg.Login); err != nil {
			log.Fatalf("could not type: %v", err)
		}

		if err = page.Type("#mat-input-1", cfg.Password); err != nil {
			log.Fatalf("could not type: %v", err)
		}

		if err = page.Press("#mat-input-1", "Enter"); err != nil {
			log.Fatalf("could not press: %v", err)
		}

		if _, err = page.Screenshot(playwright.PageScreenshotOptions{
			Path: playwright.String("screenshots/rc-web2.png"),
		}); err != nil {
			log.Fatalf("could not create screenshot: %v", err)
		}

		name, err := page.InnerText(".leading-snug")
		if err != nil {
			log.Fatalf("could not create screenshot: %v", err)
		}

		convey.So(name, convey.ShouldEqual, `Владыка Лев, Митрополит Иркутский и Ангарский`)

		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}

	})
}
