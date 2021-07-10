package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini配置文件解析器

//MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//RedisConfig ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	//0.参数的校验
	// 0.1 传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("data参数必须是指针类型")
	}
	// 0.2 传进来的data参数必须是结构体类型（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("data参数必须是结构体类型")
	}
	//1.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	// fmt.Println(lineSlice)
	var structName string
	//2.一行一行读数据
	for index, line := range lineSlice {
		//去掉字符串首尾空格
		line = strings.TrimSpace(line)

		// 2.1 如果是注释或者空行就跳过
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头的就表示是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			//把这一行首尾的[]去掉，取到中间的内容,把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			//根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体,把字段名记下来
					structName = field.Name
					// fmt.Println(sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[开头的就是=分割的键值对
			//1.以等号分割这行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			idx := strings.Index(line, "=")
			key := strings.TrimSpace(line[:idx])
			val := strings.TrimSpace(line[idx+1:])
			//2.根据structName，去data里面把对应的嵌套结构取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("dota中的%s，应该是个结构体", structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			//3.遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) //tag信息是存储在类型信息中
				fieldType = field
				if field.Tag.Get("ini") == key {
					//找到了对应的字段
					fieldName = field.Name
					break
				}
			}
			//4.如果key==tag,给这个字段赋值
			//4.1根据fieldName 去取出这个字段
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字符
				continue
			}
			fileobj := sValue.FieldByName(fieldName)
			//4.2对其赋值
			switch fieldType.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
				{
					valueInt, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						err = fmt.Errorf("line:%d, value type error", index+1)
						return err
					}
					fileobj.SetInt(valueInt)
				}
			case reflect.String:
				{
					fileobj.SetString(val)
				}
			case reflect.Bool:
				{
					valueBool, err := strconv.ParseBool(val)
					if err != nil {
						err = fmt.Errorf("line:%d, value type error", index+1)
						return err
					}
					fileobj.SetBool(valueBool)
				}
			case reflect.Float32, reflect.Float64:
				{
					valueFloat, err := strconv.ParseFloat(val, 64)
					if err != nil {
						err = fmt.Errorf("line:%d, value type error", index+1)
						return err
					}
					fileobj.SetFloat(valueFloat)
				}

			}

		}

	}
	return nil
}

func main() {
	var cfg Config
	fmt.Println(cfg)
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
	}
	fmt.Println(cfg)
}
