package convert

// // 这个文件主要提供方法，把数据读取过来

// // LoadYamlFile is 读取一个文件
// func LoadYamlFile(path string) (schemes []*Scheme, err error) {

// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Error(err)
// 		return schemes, err
// 	}
// 	schemes, err = DecodeYaml(file)
// 	if err != nil {
// 		log.Error(err)
// 		return schemes, err
// 	}
// 	return schemes, nil
// }

// // LoadFiles is 读取多个问天系统
// func LoadFiles() {

// }

// // DecodeYaml is load
// func DecodeYaml(file *os.File) (schemes []*Scheme, err error) {

// 	dec := yaml.NewDecoder(file)
// 	for {
// 		s := &Scheme{}
// 		err = dec.Decode(s)

// 		if err != nil {
// 			break
// 		}
// 	}
// 	// if err != nil {
// 	// 	log.Error(err)
// 	// 	return schemes, err
// 	// }
// 	return schemes, nil
// }
