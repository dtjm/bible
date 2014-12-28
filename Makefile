build/bible:
	mkdir -pv build
	GOPATH=$(PWD)/Godeps/_workspace:$(GOPATH) GOBIN=$(PWD)/build go install -v .

/usr/local/include/portaudio.h:
	brew install portaudio

/usr/local/include/mpg123.h:
	brew install mpg123