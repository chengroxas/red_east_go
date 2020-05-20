package minterface

type ConsumerInterface interface {
	Init() error
	Perform() error
}
