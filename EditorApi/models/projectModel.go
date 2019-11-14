package models

import "time"

type User struct {
	Id             uint
	Email          string
	Password       string
	Name           string
	Demandes       []*Demande       `orm:"reverse(many)"`
	UserChallenges []*UserChallenge `orm:"reverse(many)"`
	Language       *Language        `orm:"null;rel(one);on_delete(set_null)"`
}

type Demande struct {
	Id        uint
	Cible     uint
	Reponse   bool
	User      *User      `orm:"rel(fk)"`
	Challenge *Challenge `orm:"reverse(one)"`
}

type Challenge struct {
	Id             uint
	Demande        *Demande    `orm:"null;rel(one);on_delete(set_null)"`
	Exercices      []*Exercice `orm:"reverse(many)"`
	IdVainqueur    uint
	UserChallenges []*UserChallenge `orm:"reverse(many)"`
}

type Exercice struct {
	Id         uint
	Titre      string
	Difficulte string
	Temps      time.Duration
	Challenge  *Challenge `orm:"rel(fk)"`
	Enonce     *Enonce    `orm:"null;rel(one);on_delete(set_null)"`
	Solution   []*Testeur `orm:"reverse(many)"`
}

type Enonce struct {
	Id          uint
	Exemple     string
	InputOutput string
	Exercice    *Exercice `orm:"reverse(one)"`
}

type UserChallenge struct {
	Id           uint
	Code         string
	TempEexution time.Duration
	Valide       bool
	User         *User      `orm:"rel(fk)"`
	Challenge    *Challenge `orm:"rel(fk)"`
}

type Language struct {
	Id          uint
	LanguageId  uint
	NomLanguage string
	CodeBase    string
	User        *User `orm:"reverse(one)"`
}

type Testeur struct {
	Id              uint
	CodeTest        string
	ResultatAttendu string
	Exercice        *Exercice `orm:"rel(fk)"`
}
