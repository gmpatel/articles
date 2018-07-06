package controller

import (
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	SetLogger(logrus.StandardLogger())
}
