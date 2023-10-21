package models

type CustomInput struct {
	UserName   string
	NickName   *string
	Age        int
	Handicaped bool
	Tag        rune
}

type Output struct {
	UserName   string
	NickName   string
	Age        string
	Handicaped string
	Tag        string
}
