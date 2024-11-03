package models

import (
    "time"
    "enigma-new/database"
    "enigma-new/services"
)

/*type Key struct {
    gorm.Model
    Key       string    `gorm:"not null"`
	IsCurrent  bool   `gorm:"boolean"`
    
}*/

type Key struct {
    ID        uint   `gorm:"primary_key" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Key       string `gorm:"not null" json:"key"`
    IsCurrent bool   `gorm:"boolean" json:"is_current"`
}



type InputDecrypt struct {
	FirstWord string `json:"first_word"`
	Document  string `json:"document"`
}

type DecryptResult struct {
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}


func DecryptDocument(input InputDecrypt) DecryptResult {
    var key Key

    if err := database.DBOpen(); err != nil {
        return DecryptResult{Error: "Falha no banco de dados", Result: ""}
    }
    defer database.DBClose()

    tx := database.DB.Begin()
    if tx.Error != nil {
        return DecryptResult{Error: "Erro ao iniciar a transação", Result: ""}
    }

    if input.FirstWord == "" {
        if err := tx.Last(&key).Error; err != nil {
            tx.Rollback()
            return DecryptResult{Error: "Nenhuma chave atual encontrada", Result: ""}
        }
    } else {
        key.Key = input.FirstWord
        if err := tx.Create(&key).Error; err != nil {
            tx.Rollback()
            return DecryptResult{Error: "Erro ao registrar a nova chave", Result: ""}
        }

        var previousKey Key
        if err := tx.Where("is_current = ?", true).First(&previousKey).Error; err == nil {
            previousKey.IsCurrent = false
            if err := tx.Save(&previousKey).Error; err != nil {
                tx.Rollback()
                return DecryptResult{Error: "Erro ao atualizar a chave anterior", Result: ""}
            }
        }

        key.IsCurrent = true
        if err := tx.Save(&key).Error; err != nil {
            tx.Rollback()
            return DecryptResult{Error: "Erro ao definir a nova chave como atual", Result: ""}
        }
    }

    if err := tx.Commit().Error; err != nil {
        return DecryptResult{Error: "Erro ao confirmar a transação", Result: ""}
    }

    result, err := services.DecryptDocumentWithKey(input.Document, key.Key)
    if err != nil {
        return DecryptResult{Error: "Erro ao descriptografar o documento", Result: ""}
    }

    return DecryptResult{Error: "Documento descriptografado com sucesso", Result: result}
}

func ListKeys(page int) (DecryptResult,[]Key) {
    if err := database.DBOpen(); err != nil {
        return DecryptResult{Error: "Falha no banco de dados"},nil
    }
    defer database.DBClose()

    var keys []Key

    if page <= 0 {
        page = 1 
    }

    offset := (page - 1) * 5

    if err := database.DB.Limit(5).Offset(offset).Find(&keys).Error; err != nil {
        return DecryptResult{Error: "Erro ao buscar chaves"},nil
    }

    return DecryptResult{Error:"",Result: ""},keys
}

func GetCurrentKey() (Key,error) {
    var key Key
    if err := database.DBOpen(); err != nil {
        return Key{}, err
    }

    defer database.DBClose()

    if err := database.DB.Where("is_current = ?", true).First(&key).Error; err != nil {
        return Key{}, err
    }

    return key, nil
}







