package domain

type CompanyType uint8

const (
	Corporations CompanyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
)
