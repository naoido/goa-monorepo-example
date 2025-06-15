DIRS := $(wildcard microservices/*/)
TARGETS := $(notdir $(patsubst %/,%,$(DIRS)))

.PHONY: $(TARGETS) $(addprefix run-, $(TARGETS)) $(addprefix gen-, $(TARGETS))
$(addprefix gen-, $(TARGETS)): gen-%:
	@echo "Running goa gen: $*"
	cd microservices/$* && export JWT_SECRET=secret && goa gen goa-example/microservices/$*/design

$(addprefix example-, $(TARGETS)): example-%:
	@echo "Running goa gen: $*"
	cd microservices/$* && export JWT_SECRET=secret && goa example goa-example/microservices/$*/design

$(addprefix run-, $(TARGETS)): run-%:
	@echo "Running server: $*"
	cd microservices/$* && export JWT_SECRET=secret && go build -o server ./cmd/$* && ./server
