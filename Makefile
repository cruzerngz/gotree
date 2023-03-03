BINARIES = gotree
BINDIR = bin

# this target should be the first one in the Makefile
default:
	@$(MAKE) --no-print-directory all

all: builddir $(foreach bin, $(BINARIES), $(BINDIR)/$(bin))
	make $(BINARIES)

## magic sauce targets
# build when binary name provided
%: $(BINDIR)/%
	@make $(BINDIR)/$*

# build when binary path provided
$(BINDIR)/%:
	go build -o $(BINDIR)/$* cmd/$*/$*.go

builddir:
	@mkdir -p ./${BINDIR}

clean:
	@rm -f ./bin/*
