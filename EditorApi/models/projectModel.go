package models

import "time"

type User struct {
	Id            uint
	Username      string
	Password      string           `valid:"Required"`
	Email         string           `valid:"Email"`
	Demandes      []*Demande       `orm:"reverse(many)"`
	Vainqueur     *Vainqueur       `orm:"reverse(one)"`
	Language      *Language        `orm:"rel(fk)"`
	UserChallenge []*UserChallenge `orm:"reverse(many)"`
}

type Language struct {
	Id          uint
	NLanguage   uint
	NomLanguage string
	BaseCode    string
	User        []*User `orm:"reverse(many)"`
}

type Demande struct {
	Id        uint
	Cible     uint
	Reponse   bool
	User      *User      `orm:"rel(fk)"`
	Challenge *Challenge `orm:"reverse(one)"`
}

type Challenge struct {
	Id               uint
	Demande          *Demande            `orm:"null;rel(one);on_delete(set_null)"`
	Vainqueur        *Vainqueur          `orm:"reverse(one)"`
	ExoChallengeRand []*ExoChallengeRand `orm:"reverse(many)"`
}

type Vainqueur struct {
	Id        uint
	Challenge *Challenge `orm:"null;rel(one);on_delete(set_null)"`
	User      *User      `orm:"null;rel(one);on_delete(set_null)"`
}

type ExoChallengeRand struct {
	Id            uint
	Challenge     *Challenge       `orm:"rel(fk)"`
	Exercice      []*Exercice      `orm:"reverse(many)"`
	UserChallenge []*UserChallenge `orm:"reverse(many)"`
}

type Exercice struct {
	Id               uint
	Titre            string
	Enonce           string
	Difficulte       string
	TempResolution   time.Duration
	ExoChallengeRand []*ExoChallengeRand `orm:"rel(m2m)"`
	Exercice         []*Testeur          `orm:"reverse(many)"`
}

type Testeur struct {
	Id             uint
	Fonction       string
	ResultaAttendu string
	Exercice       *Exercice `orm:"rel(fk)"`
}

type UserChallenge struct {
	Id               uint
	Code             string
	TempExec         time.Duration
	Satus            bool
	User             *User             `orm:"rel(fk)"`
	ExoChallengeRand *ExoChallengeRand `orm:"rel(fk)"`
}

/*
type User struct {
	Id             uint
	Username       string
	Password       string           `valid:"Required"`
	Email          string           `valid:"Email"`
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
	Enonce     *Enonce    `orm:"reverse(one)"`
	Solution   []*Testeur `orm:"reverse(many)"`
}

type Enonce struct {
	Id          uint
	Exemple     string
	InputOutput string
	Exercice    *Exercice `orm:"null;rel(one);on_delete(set_null)"`
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


*/
