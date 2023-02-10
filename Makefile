GITTAG=$(shell git describe --abbrev=0)
GITHASH=$(shell git rev-parse --short HEAD)
BUILDTIME=$(shell LC_ALL=C date)

all: cleanup build

build:
	go build -ldflags "-X 'mailcp/cmd.versionNumber=$(GITTAG)' -X 'mailcp/cmd.gitHash=$(GITHASH)' -X 'mailcp/cmd.buildTime=$(BUILDTIME)' -X 'mailcp/cmd.author=Viktor Chukhantsev' -X 'mailcp/cmd.email=mailcp@vch.su'"

cleanup:
	rm -f mailcp
