help:
	@echo "commands:"
	@echo "  run-p => docker-compose run -e COMMAND=partitioning --rm app"
	@echo "  down   => docker-comopose down"
	@echo "  login  => login MySQL"

run-p:
	@docker-compose run -e COMMAND=partitioning --rm app

run-t:
	@docker-compose run -e COMMAND=table --rm app

down:
	@docker-compose down

login:
	@docker exec -it $$(docker ps | grep "ysql-on-go-for-multitenancy_db" | awk '{print $$1;}') mysql -u root -D test_db

# benchmark:
# 	@docker exec -it $$(docker ps | grep "ysql-on-go-for-multitenancy_db" | awk '{print $$1;}') mysqlslap --no-defaults --concurrency=500 --iterations=30 --query="SELECT * FROM users WHERE campany_id = 1"

.PHONY: help run-p run-t down login
