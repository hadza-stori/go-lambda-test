.PHONY: build

run_local_dev:
	@sam build "HadzaPokemonFunction"
	@sam local start-api
