# vim: set noet ci pi sts=0 sw=4 ts=4 :
# http://www.gnu.org/software/make/manual/make.html
# http://linuxlib.ru/prog/make_379_manual.html
SHELL := $(shell which bash)
#SHELL := $(shell which sh)
DEBUG ?= 0

########################################################################
# Default variables
########################################################################
-include .env
export
########################################################################
GOBIN := $(or $(GOBIN), $(GOPATH)/bin)
SUDO := $(or $(SUDO),)
GO111MODULE := $(or $(GO111MODULE), on)
GOROOT := $(GOPATH/src)
TAGS       :=
LDFLAGS    := -w -s
GOFLAGS    :=

.PHONY: db-init
db-init:
	set -o allexport;
	source .env;
	set +o allexport;
	./infra/db/$(DB_CONNECTION)_init.sh
