run:
	xcaddy run --config examples/Caddyfile

build:
	xcaddy build --with github.com/sebastianbrunnert/caddy-advanced-metrics@latest --output dist/caddy