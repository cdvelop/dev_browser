package dev_browser

func Add() *Browser {

	b := Browser{}
	b.captureArguments()

	return &b
}
