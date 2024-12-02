DEPS_DIR = node_modules

clean:
	@if [ -d "$(DEPS_DIR)" ]; then \
		echo "Removing $(DEPS_DIR)..."; \
		rm -r "$(DEPS_DIR)"; \
		rm -r "package-lock.json"; \
	fi

# Setup 
setup: clean
	@npm i

# Run server in dev mode
serve:
	@npm run serve

# Run prettier source
pretty:
	@npx prettier . --write
