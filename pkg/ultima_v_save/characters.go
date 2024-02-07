package ultima_v_save

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

const savedGamFileSize = 4192
const startPositionOfCharacters = 0x02

func getSavedGamRaw(savedGamFilePath string) ([]byte, error) {
	// Open the file in read-only mode and as binary
	file, err := os.OpenFile(savedGamFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	buffer := make([]byte, savedGamFileSize)
	n, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}

	if n != savedGamFileSize {
		return nil, fmt.Errorf("expected file of size 4192 but was %d", n)
	}

	return buffer, nil
}

func GetCharactersFromSave(savedGamFilePath string) (*SaveGame, error) {
	// Open the file in read-only mode and as binary
	buffer, err := getSavedGamRaw(savedGamFilePath)
	if err != nil {
		return nil, err
	}

	var saveGame = SaveGame{}

	// Overlay player characters over memory buffer to easily consume data
	characterPtr := (*[NPlayers]PlayerCharacter)(unsafe.Pointer(&buffer[startPositionOfCharacters]))
	saveGame.Characters = *characterPtr

	return &saveGame, nil
}

func (p *PlayerCharacter) GetNameAsString() string {
	return strings.TrimRight(string(p.Name[:]), string(rune(0)))
}

//func getClassStr(characterClass ultima_v_save.CharacterClass) string {
//	value, bExists := ultima_v_save.CharacterClassMap[characterClass]
//	if bExists {
//		return value
//	}
//	return ""
//}
//
//func getClassByStr(classStr string) (ultima_v_save.CharacterClass, bool) {
//	characterClass, bExists := ultima_v_save.FindKeyByValueT[ultima_v_save.CharacterClass, string](ultima_v_save.CharacterClassMap, classStr)
//	if bExists {
//		return characterClass, true
//	}
//	return ultima_v_save.Avatar, false
//}
