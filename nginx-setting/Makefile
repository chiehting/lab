.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

nr: ## nginx config reload
	docker exec -it lib_nginx_1 nginx -s reload

up: ## Start Nginx
	$(call docker-compose,up -d)

down: ## Stop Nginx
	$(call docker-compose,down)

ps: ## Show process
	$(call docker-compose,ps)

logs: ## Show logs
	$(call docker-compose,logs -f --tail 100)

.PHONY: du
du:
	@make down
	@make up

.PHONY: dul
dul:
	@make du
	@make logs

define docker-compose
    @docker-compose -f docker-compose.yml -p lib $1;
endef
