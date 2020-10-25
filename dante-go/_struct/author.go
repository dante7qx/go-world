package _struct

import "fmt"

type Author struct {
	Name string
	Age int8
	Description string
}

func init() {
	fmt.Println("初始化Author......")
	fmt.Println("初始化Address......")
}

/**
	方法 func (variable_name variable_data_type) function_name() [return_type] {}
 */
func (author Author) SelfBook(address Address, book string) string {
	return fmt.Sprintf("%s, %s在%d岁创作了《%s》。", buildAddress(address), author.Name, author.Age, book)
}

/**
	函数 func function_name( [parameter list] ) [return_types] {}
 */
func buildAddress(address Address) string {
	return fmt.Sprintf("%s%s%s%d", address.Country, address.Province, address.Street, address.Code)
}
