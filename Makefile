generate:
	docker run -it -v ${PWD}:/task_wizard_schema -w /task_wizard_schema golang:latest go generate ./ent
