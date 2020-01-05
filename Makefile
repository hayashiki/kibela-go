

.PHONY: gox
gox:
	gox -ldflags=$(LDFLAGS) -output="bin/$(NAME)_{{.OS}}_{{.Arch}}"