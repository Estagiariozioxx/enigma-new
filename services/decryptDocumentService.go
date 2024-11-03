package services

import (
	"errors"
	//"log"
	//"log"
	"strings"

)
func DecryptDocumentWithKey(document, key string) (string, error) {
    if key == "" {
        return "", errors.New("chave não pode ser vazia")
    }
	 chars := "abcdefghijklmnopqrstuvwyzàáãâéêóôõíúç"

	document = strings.ToLower(document)

	var result []rune

	keyChar := key[0]
	keyDocument := document[0]
	difer := int(keyDocument) - int(keyChar)

	for _, char := range document {

		indexChars := indexOf(chars, char)

		
		if indexChars == -1{
			result = append(result, char)

		}else{
			newIndexChars:=indexChars
			
			newIndexChars-=difer

			newIndexChars = (newIndexChars + len(chars)) % len(chars) //para garantir que fique dentro do alfabeto
			result = append(result,rune(chars[newIndexChars]))
		}
		
    }
    return string(result), nil
}

func indexOf(alphabet string, char rune) int {
    for i, c := range alphabet {
        if c == char {
            return i
        }
    }
    return -1 
}

