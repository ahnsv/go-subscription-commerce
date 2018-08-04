package iamport

import (
	"fmt"
	"net/http"

	"github.com/mgsmurf/go-iamport"
)

func Init() {
	client := &http.Client{} // 상황에 맞는 클라이언트 사용
	iam := iamport.NewClient("", "", client)
	pay, err := iam.GetPaymentImpUID("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pay.Amount)
	fmt.Println(pay.MerchantUID)
}
