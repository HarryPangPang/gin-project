package controller

import (
	"gmt-go/helper"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = helper.Logger()
}
