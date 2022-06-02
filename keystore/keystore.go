package keystore

// export keystore
func CreateKeystore(priv, password string) (string, error) {

	encrypted, err := EncryptV3([]byte(priv), password) // 生成 key_store
	if err != nil {
		return "", err
	}

	// wallets, _ := filepath.Abs("wallets")
	// filename := wallets + "/" + "test" + ".json"
	// f, err := os.Create(filename)
	// defer f.Close()
	// if err != nil {
	// 	return false, err
	// }
	// _, err = f.Write(encrypted)
	// if err != nil {
	// 	return false, err
	// }

	return string(encrypted), nil
}

// import keystore
func Importkeystore(keystore, password string) (string, error) {

	// wallets, _ := filepath.Abs("wallets")
	// filename := wallets + "/" + file

	// f, err := os.OpenFile(filename, os.O_RDONLY, 0600)

	// defer f.Close()

	// if err != nil {
	// 	return "", err
	// }
	// contentByte, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	return "", err
	// }
	found, err := DecryptV3([]byte(keystore), password) // 通过密码解出 私钥
	if err != nil {
		return "DecryptV3", err
	}

	return string(found), nil
}
