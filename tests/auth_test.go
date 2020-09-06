package tests

const (
	login         = "37445"
	password      = "9420"
	firebaseToken = "qweett"
)

//func TestAuthLoginWithEmail(t *testing.T) {
//	convey.Convey("Login with email and password success", t, func() {
//
//		client := http.Client{}
//		req, err := http.NewRequest("POST",
//			fmt.Sprintf("%s/api/v1/auth/login", cfg.Host),
//			strings.NewReader(fmt.Sprintf(`{"email":"%s","password":"%s"}`, login, password)))
//		if err != nil {
//			return
//		}
//		req.Header.Set("Content-Type", "application/json")
//
//		resp, _ := client.Do(req)
//		defer resp.Body.Close()
//
//		convey.Convey("Check response", func() {
//			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
//		})
//	})
//}
//
//func TestAuthLoginWithEmailAndFirebaseToken(t *testing.T) {
//	convey.Convey("Login with email and password with firebase token success", t, func() {
//
//		client := http.Client{}
//		req, err := http.NewRequest("POST",
//			fmt.Sprintf("%s/api/v1/auth/login", cfg.Host),
//			strings.NewReader(fmt.Sprintf(`{"email":"%s","password":"%s","firebase_token":"%s"}`, login, password, firebaseToken)))
//		if err != nil {
//			return
//		}
//		req.Header.Set("Content-Type", "application/json")
//
//		resp, _ := client.Do(req)
//		defer resp.Body.Close()
//
//		convey.Convey("Check response", func() {
//			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
//		})
//	})
//}
//
//func TestAuthLoginWithLogin(t *testing.T) {
//	convey.Convey("Login with login and password success", t, func() {
//
//		client := http.Client{}
//		req, err := http.NewRequest("POST",
//			fmt.Sprintf("%s/api/v1/auth/login", cfg.Host),
//			strings.NewReader(fmt.Sprintf(`{"login":"%s","password":"%s"}`, login, password)))
//		if err != nil {
//			return
//		}
//		req.Header.Set("Content-Type", "application/json")
//
//		resp, _ := client.Do(req)
//		defer resp.Body.Close()
//
//		convey.Convey("Check response", func() {
//			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
//		})
//	})
//}
//
//func TestAuthLoginWithLoginAndFirebaseToken(t *testing.T) {
//	convey.Convey("Login with login and password with firebase token success", t, func() {
//
//		client := http.Client{}
//		req, err := http.NewRequest("POST",
//			fmt.Sprintf("%s/api/v1/auth/login", cfg.Host),
//			strings.NewReader(fmt.Sprintf(`{"login":"%s","password":"%s","firebase_token":"%s"}`, login, password, firebaseToken)))
//		if err != nil {
//			return
//		}
//		req.Header.Set("Content-Type", "application/json")
//
//		resp, _ := client.Do(req)
//		defer resp.Body.Close()
//
//		convey.Convey("Check response", func() {
//			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
//		})
//	})
//}
//
//func TestAuthLoginWithWrongCredsAndFirebaseToken(t *testing.T) {
//	convey.Convey("Login with wrong login and password with firebase token unauthorized", t, func() {
//
//		client := http.Client{}
//		req, err := http.NewRequest("POST",
//			fmt.Sprintf("%s/api/v1/auth/login", cfg.Host),
//			strings.NewReader(fmt.Sprintf(`{"login":"1000000","password":"%s","firebase_token":"%s"}`, password, firebaseToken)))
//		if err != nil {
//			return
//		}
//		req.Header.Set("Content-Type", "application/json")
//
//		resp, _ := client.Do(req)
//		defer resp.Body.Close()
//
//		convey.Convey("Check response", func() {
//			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
//		})
//	})
//}
