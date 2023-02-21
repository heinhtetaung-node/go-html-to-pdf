package main

import (
	"fmt"
	"log"
	"os"

	u "github.com/RamonMamon/html-pdf-demo/pdfGenerator"
)

type Translation struct {
	CompanyName string
}

type Package struct {
	Order uint8
	Name  string
	Price string
}

type Order struct {
	Packages []Package
}

type Content struct {
	PolicyHolder string
	// Order              Order
	Packages           []Package
	Language           string
	LeadID             string
	Date               string
	InsuredPerson      string
	InsurerName        string
	PackageOrder       string
	InsuranceKind      string
	PackagePrice       string
	NetPremium         string
	Vat                string
	TotalDiscount      string
	TotalPremium       string
	Note               string
	PaymentMethod      string
	CreditCard         string
	PromptPay          string
	AuthorizedPerson   string
	CompanyName        string
	PaymentReceiptData map[string]interface{}
}

func main() {

	r := u.NewRequestPdf("")

	//html template path
	// templatePath := "templates/receipt.html"

	//path for download pdf
	engPath := "output/eng-sample.pdf"

	// thaiPath := "output/thai-sample.pdf"

	if _, err := os.Stat("output"); os.IsNotExist(err) {

		err := os.Mkdir("output", 0755)

		if err != nil {
			log.Fatal(err)
		}
	}

	//html template data
	// thaiData := Content{
	// 	LeadID:        "L9852803",
	// 	Date:          "15/06/2022",
	// 	InsuredPerson: "Insured name",
	// 	InsurerName:   "บริษัท กรุงเทพประกันภัย จำกัด (มหาชน)",
	// 	PackageOrder:  "1",
	// 	InsuranceKind: "ประกันภาคสมัครใจประกันภันชั้น 3คุ้มครองเฉพาะรถคู่กรณี",
	// 	PackagePrice:  "1,000,000",
	// 	NetPremium:    "1,000",
	// 	Vat:           "10%",
	// 	TotalDiscount: "0",
	// 	TotalPremium:  "0",
	// 	Note:          "เอกสารนี้เป็นเพียงการออกใบเสร็จรับเงินชั่วคราวเท่านั้น",
	// 	PaymentMethod: "บัตรเครดิต / บัตรเดบิตพร้อมเพย์ คิวอาร์",
	// 	CreditCard:    "checked",
	// 	// PromptPay:        "checked",
	// }

	p1 := Package{Order: 1, Price: "10000", Name: "Full-Stack Developer"}
	p2 := Package{Order: 2, Price: "7000", Name: "Back-end Developer"}
	p3 := Package{Order: 3, Price: "7000", Name: "Back-end Developer"}

	englishData := Content{
		PolicyHolder: "Hein",
		// Order: Order{
		// 	Packages: []Package{p1, p2},
		// },
		Packages: []Package{
			p1,
			p2,
			p3,
		},
		Language:           "en",
		LeadID:             "L9852803",
		Date:               "15/06/2022",
		InsuredPerson:      "Insured name",
		InsurerName:        "Bangkok Insurance Public Company Limited",
		PackageOrder:       "1",
		InsuranceKind:      "Voluntary Insurance\nType 3\nThird party coverage",
		PackagePrice:       "1,000,000",
		NetPremium:         "1,000",
		Vat:                "10%",
		TotalDiscount:      "0",
		TotalPremium:       "0",
		Note:               "This document is a temporary receipt document only.",
		PaymentMethod:      "Credit / Debit Card PromptPay QR",
		CreditCard:         "checked",
		PromptPay:          "",
		AuthorizedPerson:   "",
		CompanyName:        "",
		PaymentReceiptData: map[string]interface{}{"PolicyHolder": "Hein"},
	}

	if err := r.ParseTemplate("templates/eng-receipt/index.html", englishData); err == nil {
		ok, _ := r.GeneratePDF(engPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}

	// if err := r.ParseTemplate("templates/thai-receipt/index.html", thaiData); err == nil {
	// 	ok, _ := r.GeneratePDF(thaiPath)
	// 	fmt.Println(ok, "pdf generated successfully")
	// } else {
	// 	fmt.Println(err)
	// }
}
