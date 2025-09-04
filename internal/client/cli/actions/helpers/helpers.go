package helpers

import "fmt"

const TEXT_REQUEST_TO_SERVER = "Запрос на сервер..."

func SuccessBadge(str string) {
	fmt.Printf("\033[32m[ok] \033[0m")
	fmt.Println(str)
}

func WarnBadge(str string) {
	fmt.Printf("\033[33m[warring] \033[0m")
	fmt.Println(str)
}

func ErrorBadge(str string) {
	fmt.Printf("\033[31m[error] \033[0m")
	fmt.Println(str)
}

func SecretData(str string) {
	fmt.Printf("\033[34m`%s` \033[0m\n", str)
}
