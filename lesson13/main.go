package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	v1 "github.com/YarikGuk/GoDcModule"
	v2 "github.com/YarikGuk/GoDcModule/v2"
	"log"
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
	В рамках задачи будем работать с картотекой известного врача. Нужно будет написать модуль с несколькими версиями: v1.0.0, v1.1.0, v2.0.0, v2.1.0. Модуль должен прочитать файл со следующим содержимым:

	{"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
	{"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
	{"name":"Лисичка","age":3,"email":"alice@mail.ru"}

	v1.0.0 должна создавать файл с содержимым:

	[{«name»:»Ёжик","age":10,"email":"ezh@mail.ru"},
	{"name":"Зайчик","age":2,"email":"zayac@mail.ru"},
	{«name":"Лисичка","age":3,"email":"alice@mail.ru"}]

	v1.1.0 должна сортировать данные по полю age по
	возрастанию:
	[{«name":"Зайчик","age":2,"email":"zayac@mail.ru"},
	{«name":"Лисичка","age":3,"email":"alice@mail.ru"}{«name»:»
	Ёжик","age":10,"email":"ezh@mail.ru"}]

	v2.0.0 должна создавать файл с содержимым:
	<?xml version="1.0" encoding="UTF-8"?>
	<patients>
	<Patient>
	<Name>Ёжик</Name>
	<Age>10</Age>
	<Email>ezh@mail.ru</Email>
	</Patient>
	<Patient>
	<Name>Зайчик</Name>
	<Age>2</Age>
	<Email>zayac@mail.ru</Email>
	</Patient>
	<Patient>
	<Name>Лисичка</Name>
	<Age>3</Age>
	<Email>alice@mail.ru</Email>
	</Patient>
	</patients>
	v2.1.0 должна сортировать данные по полю age по возрастанию:
	<?xml version="1.0" encoding="UTF-8"?>
	<patients>
	<Patient>
	<Name>Зайчик</Name>
	<Age>2</Age>
	<Email>zayac@mail.ru</Email>
	</Patient>
	<Patient>
	<Name>Лисичка</Name>
	<Age>3</Age>
	<Email>alice@mail.ru</Email>
	</Patient>
	<Patient>
	<Name>Ёжик</Name>
	<Age>10</Age>
	<Email>ezh@mail.ru</Email>
	</Patient>
	</patients>
	Модуль должен содержать функцию Do, которая принимает два
	строковых параметра: путь файла откуда прочитать данные, путь
	файла, в который записать данные в требуемом формате; и
	возвращать ошибку. Пример использования модуля:
	package main
	import (
	format ...
	)
	func main() {
	err := format.Do("patients", «result")
	if err != nil {
	…
	}
	}
	Должна быть возможность подключить любую из версий: v1.0.0, v1.1.0,
	v2.0.0, v2.1.0.
	*/

	v1.Do("patient", "result_v1_1_0.json")
	if err != nil {
		log.Fatalf("Success v1.1.0: %v", err)
	}
	fmt.Println("Processed with v1.1.0")
	v2.Do("patient", "result_v2_1_0.xml")
	if err != nil {
		log.Fatalf("Error v2.0.0: %v", err)
	}
	fmt.Println("Success v2.1.0")
}
