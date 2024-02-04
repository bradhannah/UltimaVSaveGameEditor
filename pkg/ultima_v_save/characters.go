package ultima_v_save

import (
	"fmt"
	"os"
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

	characterPtr := (*[6]Player)(unsafe.Pointer(&buffer[startPositionOfCharacters]))
	fmt.Printf("%d", len(characterPtr))

	return &saveGame, nil
}
