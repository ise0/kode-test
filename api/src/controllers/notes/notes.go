package noteController

import (
	noteModel "github.com/ise0/kode-test/src/models/note"
	"github.com/ise0/kode-test/src/services/speller"
)

func AddNew(userId int, text string) (noteModel.Note, error) {
	spellerRes, err := speller.CheckText(text, speller.IGNORE_DIGITS|speller.IGNORE_URLS)
	if err != nil {
		return noteModel.Note{}, err
	}

	rtext := []rune(text)
	for _, v := range spellerRes {
		rtext = append(rtext[:v.Pos], append([]rune(v.S[0]), rtext[v.Pos+v.Len:]...)...)
	}

	n, err := noteModel.New(userId, string(rtext))
	if err != nil {
		return noteModel.Note{}, err
	}

	return n, nil
}
