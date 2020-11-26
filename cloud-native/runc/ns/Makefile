

.PHONY: all 

all: ns_uts ns_exec

ns_uts:
	gcc ns_uts.c -o bin/ns_uts
ns_exec:
	gcc ns_exec.c -o bin/ns_exec
ns_share:
	gcc ns_unshare.c -o bin/ns_unshare
clean:
	rm -rf bin/*
