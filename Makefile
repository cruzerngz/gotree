BINARIES = gotree
BINDIR = bin
SRCDIR = cmd

BINPATHS = $(foreach bin, $(BINARIES), $(BINDIR)/$(bin))

.PHONY: all builddir clean

# this target should be the first one in the Makefile
default:
	@$(MAKE) --no-print-directory all

all: builddir $(BINPATHS)

## matching names get built
$(BINDIR)/%: $(SRCDIR)/%.go
	go build -o $@ $<

builddir:
	@mkdir -p ./${BINDIR}

clean:
	@rm -f ./bin/*
