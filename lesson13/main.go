package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenant   string `json:"tenat"`
}

type XmlContract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenant   string `xml:"tenat"`
}

type XmlContracts struct {
	XmlContract []XmlContract `xml:"contract"`
}

type contracts2 struct {
	List []contract2 `xml:"contract"`
}
type contract2 struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenant   string `xml:"tenat"`
}

func main() {

	/* Задача 13.1
	Необходимо распарсить json {"number":1,"landlord":"Остап
	Бендер","tenat":"Шура Балаганов»}

	p.s. Последние кавычки елочкой некорректые, заменил на обычные
	*/
	contract := Contract{}
	str := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	err := json.Unmarshal([]byte(str), &contract)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON'а:", err)
	}
	fmt.Printf("%+v", contract)
	fmt.Println()

	/* Задача 13.2
	Необходимо представить в виде json структуру contract

	contract{
	Number: 2,
	Landlord: "Остап",
	tenat: "Шура",
	}
	*/
	newContract := Contract{
		Number:   2,
		Landlord: "Остап",
		Tenant:   "Шура",
	}
	res, err := json.Marshal(newContract)
	if err != nil {
		fmt.Println("Ошибка при создании  JSON'а:", err)
	}
	fmt.Printf("%v", string(res))
	fmt.Println()

	/* Задача 13.3
	Необходимо распарсить xml
	*/

	xmlTemp := `	<?xml version="1.0" encoding="UTF-8"?>
	<contracts>
	<contract>
	<number>1</number>
	<sign_date>2023-09-02</sign_date>
	<landlord>Остап</landlord>
	<tenat>Шура</tenat>
	</contract>
	<contract>
	<number>2</number>
	<sign_date>2023-09-03</sign_date>
	<landlord>Бендер</landlord>
	<tenat>Балаганов</tenat>
	</contract>
	</contracts>`

	xmlContracts := XmlContracts{}
	err = xml.Unmarshal([]byte(xmlTemp), &xmlContracts)
	if err != nil {
		fmt.Println("Ошибка при парсинге XML:", err)
	}
	fmt.Printf("%+v", xmlContracts)
	fmt.Println()

	/* Задача 13.4
	Необходимо представить в виде xml структуру contracts
	*/

	c := contract2{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenant:   "Шура Балаганов",
	}
	contracts2 := contracts2{}
	contracts2.List = append(contracts2.List, c)

	res, err = xml.Marshal(contracts2)
	if err != nil {
		fmt.Println("Ошибка при создании XML:", err)
	}
	fmt.Println(string(res))

	/* Задача 13.5
	В рамках задачи будем работать с картотекой известного врача.
	Нужно будет написать модуль с несколькими версиями:
	v1.0.0, v1.1.0, v2.0.0, v2.1.0.
	*/

}
